// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package present

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	parsers = make(map[string]ParseFunc)
	funcs   = template.FuncMap{}
)

// Template returns an empty template with the action functions in its FuncMap.
func Template() *template.Template {
	return template.New("").Funcs(funcs)
}

// Render renders the doc to the given writer using the provided template.
func (d *Doc) Render(w io.Writer, t *template.Template) error {
	data := struct {
		*Doc
		Template     *template.Template
		PlayEnabled  bool
		NotesEnabled bool
	}{d, t, PlayEnabled, NotesEnabled}
	return t.ExecuteTemplate(w, "root", data)
}

// Render renders the section to the given writer using the provided template.
func (s *Section) Render(w io.Writer, t *template.Template) error {
	data := struct {
		*Section
		Template    *template.Template
		PlayEnabled bool
	}{s, t, PlayEnabled}
	return t.ExecuteTemplate(w, "section", data)
}

type ParseFunc func(ctx *Context, fileName string, lineNumber int, inputLine string) (Elem, error)

// Register binds the named action, which does not begin with a period, to the
// specified parser to be invoked when the name, with a period, appears in the
// present input text.
func Register(name string, parser ParseFunc) {
	if len(name) == 0 || name[0] == ';' {
		panic("bad name in Register: " + name)
	}
	parsers["."+name] = parser
}

// renderElem implements the elem template function, used to render
// sub-templates.
func renderElem(t *template.Template, e Elem) (template.HTML, error) {
	var data interface{} = e
	if s, ok := e.(Section); ok {
		data = struct {
			Section
			Template *template.Template
		}{s, t}
	}
	return execTemplate(t, e.TemplateName(), data)
}

// pageNum derives a page number from a section.
func pageNum(s Section, offset int) int {
	if len(s.Number) == 0 {
		return offset
	}
	return s.Number[0] + offset
}

func init() {
	funcs["elem"] = renderElem
	funcs["pagenum"] = pageNum
}

// execTemplate is a helper to execute a template and return the output as a
// template.HTML value.
func execTemplate(t *template.Template, name string, data interface{}) (template.HTML, error) {
	b := new(bytes.Buffer)
	err := t.ExecuteTemplate(b, name, data)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}

func readLines(r io.Reader) (*Lines, error) {
	var lines []string
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return &Lines{0, lines}, nil
}

// A Context specifies the supporting context for parsing a presentation.
type Context struct {
	// ReadFile reads the file named by filename and returns the contents.
	ReadFile func(filename string) ([]byte, error)
}

// ParseMode represents flags for the Parse function.
type ParseMode int

const (
	// If set, parse only the title and subtitle.
	TitlesOnly      ParseMode = 1
	PrefixPageTitle           = "# "
)

// Parse parses a document from r.
func (ctx *Context) Parse(r io.Reader, name string, mode ParseMode) (*Doc, error) {
	doc := new(Doc)
	lines, err := readLines(r)
	if err != nil {
		return nil, err
	}

	for i := lines.line; i < len(lines.text); i++ {
		if strings.HasPrefix(lines.text[i], PrefixPageTitle) {
			break
		}

		// 作者注释忽略
		if isSpeakerNote(lines.text[i]) {
			doc.TitleNotes = append(doc.TitleNotes, lines.text[i][2:])
		}
	}

	err = parseHeader(doc, lines)
	if err != nil {
		return nil, err
	}
	if mode&TitlesOnly != 0 {
		return doc, nil
	}

	// Authors
	if doc.Authors, err = parseAuthors(lines); err != nil {
		return nil, err
	}
	// Sections
	if doc.Sections, err = parseSections(ctx, name, lines, []int{}); err != nil {
		return nil, err
	}
	return doc, nil
}

// Parse parses a document from r. Parse reads assets used by the presentation
// from the file system using ioutil.ReadFile.
func Parse(r io.Reader, name string, mode ParseMode) (*Doc, error) {
	ctx := Context{ReadFile: ioutil.ReadFile}
	return ctx.Parse(r, name, mode)
}

// isHeading matches any section heading.
var isHeading = regexp.MustCompile(`^#+ `)

var isOl = regexp.MustCompile(`^\d\. .*`)

// lesserHeading returns true if text is a heading of a lesser or equal level
// than that denoted by prefix.
func lesserHeading(text, prefix string) bool {
	return isHeading.MatchString(text) && !strings.HasPrefix(text, prefix+"#")
}

// parseSections parses Sections from lines for the section level indicated by
// number (a nil number indicates the top level).
func parseSections(ctx *Context, name string, lines *Lines, number []int) ([]Section, error) {
	var sections []Section
	for i := 1; ; i++ {
		// Next non-empty line is title.
		text, ok := lines.nextNonEmpty()
		for ok && text == "" {
			text, ok = lines.next()
		}
		if !ok {
			break
		}
		if !strings.HasPrefix(text, PrefixPageTitle) {
			lines.back()
			break
		}
		section := Section{
			Number: append(number, i),
			Title:  text[2:],
		}
		text, ok = lines.nextNonEmpty()

		for ok && !strings.HasPrefix(text, PrefixPageTitle) {
			var e Elem
			switch {
			case strings.HasPrefix(text, "## "):
				e = H2{text[3:]}
			case strings.HasPrefix(text, "### "):
				e = H3{text[4:]}
			case strings.HasPrefix(text, "#### "):
				e = H4{text[5:]}
			//	解析代码块，修改为markdown格式
			case strings.HasPrefix(text, "```"):
				lang := strings.TrimPrefix(text, "```")
				text, ok = lines.next()
				var s []string
				for ok && !strings.HasPrefix(text, "```") {
					s = append(s, text+"\n")
					text, ok = lines.next()
				}
				e = CodeBlock{Language: lang, Lines: s}
			//	解析列表
			case strings.HasPrefix(text, "- "):
				var b []string
				for ok && strings.HasPrefix(text, "- ") {
					b = append(b, text[2:])
					text, ok = lines.next()
				}
				lines.back()
				e = List{Bullet: b}
			case isOl.MatchString(text):
				var b []string
				for ok && isOl.MatchString(text) {
					b = append(b, text[2:])
					text, ok = lines.next()
				}
				lines.back()
				e = ZList{Bullet: b}
			//	文章中的注释跳过
			case isSpeakerNote(text):
				section.Notes = append(section.Notes, text[2:])
			case strings.HasPrefix(text, "."):
				args := strings.Fields(text)
				if args[0] == ".background" {
					section.Classes = append(section.Classes, "background")
					section.Styles = append(section.Styles, "background-image: url('"+args[1]+"')")
					break
				}
				parser := parsers[args[0]]
				if parser == nil {
					return nil, fmt.Errorf("%s:%d: unknown command %q\n", name, lines.line, text)
				}
				t, err := parser(ctx, name, lines.line, text)
				if err != nil {
					return nil, err
				}
				e = t
			default:
				var l []string
				for ok && strings.TrimSpace(text) != "" {
					if text[0] == '.' { // Command breaks text block.
						lines.back()
						break
					}
					if strings.HasPrefix(text, `\.`) { // Backslash escapes initial period.
						text = text[1:]
					}
					l = append(l, text)
					text, ok = lines.next()
				}
				if len(l) > 0 {
					e = Text{Lines: l}
				}
			}
			if e != nil {
				section.Elem = append(section.Elem, e)
			}
			text, ok = lines.nextNonEmpty()
		}
		if isHeading.MatchString(text) {
			lines.back()
		}
		sections = append(sections, section)
	}
	return sections, nil
}

func parseHeader(doc *Doc, lines *Lines) error {
	var ok bool
	// First non-empty line starts header.
	doc.Title, ok = lines.nextNonEmpty()
	if !ok {
		return errors.New("unexpected EOF; expected title")
	}
	for {
		text, ok := lines.next()
		if !ok {
			return errors.New("unexpected EOF")
		}
		if text == "" {
			break
		}
		if isSpeakerNote(text) {
			continue
		}
		const tagPrefix = "Tags:"
		if strings.HasPrefix(text, tagPrefix) {
			tags := strings.Split(text[len(tagPrefix):], ",")
			for i := range tags {
				tags[i] = strings.TrimSpace(tags[i])
			}
			doc.Tags = append(doc.Tags, tags...)
		} else if t, ok := parseTime(text); ok {
			doc.Time = t
		} else if doc.Subtitle == "" {
			doc.Subtitle = text
		} else {
			return fmt.Errorf("unexpected header line: %q", text)
		}
	}
	return nil
}

func parseAuthors(lines *Lines) (authors []Author, err error) {
	// This grammar demarcates authors with blanks.

	// Skip blank lines.
	if _, ok := lines.nextNonEmpty(); !ok {
		return nil, errors.New("unexpected EOF")
	}
	lines.back()

	var a *Author
	for {
		text, ok := lines.next()
		if !ok {
			return nil, errors.New("unexpected EOF")
		}

		// 到达章节开始部分，退出
		if strings.HasPrefix(text, "# ") {
			lines.back()
			break
		}

		if isSpeakerNote(text) {
			continue
		}

		// If we encounter a blank we're done with this author.
		if a != nil && len(text) == 0 {
			authors = append(authors, *a)
			a = nil
			continue
		}
		if a == nil {
			a = new(Author)
		}

		// Parse the line. Those that
		// - begin with @ are twitter names,
		// - contain slashes are links, or
		// - contain an @ symbol are an email address.
		// The rest is just text.
		var el Elem
		switch {
		case strings.HasPrefix(text, "@"):
			el = parseURL("http://twitter.com/" + text[1:])
		case strings.Contains(text, ":"):
			el = parseURL(text)
		case strings.Contains(text, "@"):
			el = parseURL("mailto:" + text)
		}
		if l, ok := el.(Link); ok {
			l.Label = text
			el = l
		}
		if el == nil {
			el = Text{Lines: []string{text}}
		}
		a.Elem = append(a.Elem, el)
	}
	if a != nil {
		authors = append(authors, *a)
	}
	return authors, nil
}

func parseURL(text string) Elem {
	u, err := url.Parse(text)
	if err != nil {
		log.Printf("Parse(%q): %v", text, err)
		return nil
	}
	return Link{URL: u}
}

func parseTime(text string) (t time.Time, ok bool) {
	t, err := time.Parse("15:04 2 Jan 2006", text)
	if err == nil {
		return t, true
	}
	t, err = time.Parse("2 Jan 2006", text)
	if err == nil {
		// at 11am UTC it is the same date everywhere
		t = t.Add(time.Hour * 11)
		return t, true
	}
	return time.Time{}, false
}

func isSpeakerNote(s string) bool {
	return strings.HasPrefix(s, ": ")
}
