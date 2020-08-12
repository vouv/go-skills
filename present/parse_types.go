package present

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"time"
)

// Doc represents an entire document.
type Doc struct {
	Title      string
	Subtitle   string
	Time       time.Time
	Authors    []Author
	TitleNotes []string
	Sections   []Section
	Tags       []string
}

// Author represents the person who wrote and/or is presenting the document.
type Author struct {
	Elem []Elem
}

// TextElem returns the first text elements of the author details.
// This is used to display the author' name, job title, and company
// without the contact details.
func (p *Author) TextElem() (elems []Elem) {
	for _, el := range p.Elem {
		if _, ok := el.(Text); !ok {
			break
		}
		elems = append(elems, el)
	}
	return
}

// Section represents a section of a document (such as a presentation slide)
// comprising a title and a list of elements.
type Section struct {
	Number  []int
	Title   string
	Elem    []Elem
	Notes   []string
	Classes []string
	Styles  []string
}

// HTMLAttributes for the section
func (s Section) HTMLAttributes() template.HTMLAttr {
	if len(s.Classes) == 0 && len(s.Styles) == 0 {
		return ""
	}

	var class string
	if len(s.Classes) > 0 {
		class = fmt.Sprintf(`class=%q`, strings.Join(s.Classes, " "))
	}
	var style string
	if len(s.Styles) > 0 {
		style = fmt.Sprintf(`style=%q`, strings.Join(s.Styles, " "))
	}
	return template.HTMLAttr(strings.Join([]string{class, style}, " "))
}

// Sections contained within the section.
func (s Section) Sections() (sections []Section) {
	for _, e := range s.Elem {
		if section, ok := e.(Section); ok {
			sections = append(sections, section)
		}
	}
	return
}

// Level returns the level of the given section.
// The document title is level 1, main section 2, etc.
func (s Section) Level() int {
	return len(s.Number) + 1
}

// FormattedNumber returns a string containing the concatenation of the
// numbers identifying a Section.
func (s Section) FormattedNumber() string {
	b := &bytes.Buffer{}
	for _, n := range s.Number {
		fmt.Fprintf(b, "%v.", n)
	}
	return b.String()
}

func (s Section) TemplateName() string { return "section" }

// Elem defines the interface for a present element. That is, something that
// can provide the name of the template used to render the element.
type Elem interface {
	TemplateName() string
}

// Text represents an optionally preformatted paragraph.
type Text struct {
	Lines []string
	Pre   bool
}

func (t Text) TemplateName() string { return "text" }

// Text represents an optionally preformatted paragraph.
type CodeBlock struct {
	Language string
	Lines    []string
}

func (c CodeBlock) TemplateName() string { return "code-block" }

// List represents a bulleted list.
type List struct {
	Bullet []string
}

func (l List) TemplateName() string { return "list" }

// List represents a bulleted list.
type ZList struct {
	Bullet []string
}

func (z ZList) TemplateName() string { return "zlist" }

// Lines is a helper for parsing line-based input.
type Lines struct {
	line int // 0 indexed, so has 1-indexed number of last line returned
	text []string
}

func (l *Lines) next() (text string, ok bool) {
	for {
		current := l.line
		l.line++
		if current >= len(l.text) {
			return "", false
		}
		text = l.text[current]
		// Lines starting with # are comments.
		if len(text) == 0 || !strings.HasPrefix(text, ": ") {
			ok = true
			break
		}
	}
	return
}

func (l *Lines) back() {
	l.line--
}

func (l *Lines) nextNonEmpty() (text string, ok bool) {
	for {
		text, ok = l.next()
		if !ok {
			return
		}
		if len(text) > 0 {
			break
		}
	}
	return
}

type H2 struct {
	Title string
}

func (h H2) TemplateName() string { return "h2" }

type H3 struct {
	Title string
}

func (h H3) TemplateName() string { return "h3" }

type H4 struct {
	Title string
}

func (h H4) TemplateName() string { return "h4" }
