package main

import (
	"bytes"
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/imports"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	ServerModeDebug = "debug"
	ServerModeProd  = "production"

	TransportSocket = "SocketTransport"
	TransportHttp   = "HTTPTransport"
)

// socketAddr returns the WebSocket handler address.
var socketAddr = func() string { return "ws://" + host + "/socket" }

// analyticsHTML is optional analytics HTML to insert at the beginning of <head>.
var analyticsHTML = template.HTML(os.Getenv("ANALYTICS_HTML"))

var host = "127.0.0.1"
var port = "3999"

var serverMode = ServerModeDebug

const root = "."

const localhostWarning = `
WARNING!  WARNING!  WARNING!

The tour server appears to be listening on an address that is
not localhost and is configured to run code snippets locally.
Anyone with access to this address and port will have access
to this machine as the user running gotour.

If you don't understand this message, hit Control-C to terminate this process.

WARNING!  WARNING!  WARNING!
`

func main() {

	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	}

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if err := initTour(TransportSocket); err != nil {
		log.Fatal(err)
	}

	registerHandler()

	log.Println("Serving at http://" + host + ":" + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// 注册路由
func registerHandler() {
	// Keep these static file handlers in sync with app.yaml.
	http.Handle("/favicon.ico", http.FileServer(http.Dir(filepath.Join(root, "static", "img"))))
	static := http.FileServer(http.Dir(root))
	http.Handle("/content/img/", static)
	http.Handle("/static/", static)
	http.HandleFunc("/script.js", jsHandler())

	// content
	http.HandleFunc("/", rootHandler())
	http.HandleFunc("/lesson/", lessonHandler())
	http.HandleFunc("/fmt", fmtHandler())

	// socket
	origin := &url.URL{Scheme: "http", Host: host + ":" + port}
	http.Handle("/socket", NewRunHandler(origin))

}

// rootHandler returns a handler for all the requests except the ones for lessons.
func rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := renderUI(w); err != nil {
			log.Println(err)
		}
	}
}

func jsHandler() http.HandlerFunc {
	var modTime = time.Now()
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/javascript")
		// Set expiration time in one week. todo 开发阶段关缓存
		//w.Header().Set("Cache-control", "max-age=604800")
		http.ServeContent(w, r, "", modTime, bytes.NewReader(scriptContent.Bytes()))
	}
}

// lessonHandler handler the HTTP requests for lessons.
func lessonHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lesson := strings.TrimPrefix(r.URL.Path, "/lesson/")
		if err := writeLesson(lesson, w); err != nil {
			if err == lessonNotFound {
				http.NotFound(w, r)
			} else {
				log.Println(err)
			}
		}
	}
}

type fmtResponse struct {
	Body  string
	Error string
}

func fmtHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := new(fmtResponse)
		var body string
		var err error
		if r.FormValue("imports") == "true" {
			var b []byte
			b, err = imports.Process("prog.go", []byte(r.FormValue("body")), nil)
			body = string(b)
		} else {
			body, err = gofmt(r.FormValue("body"))
		}
		if err != nil {
			resp.Error = err.Error()
		} else {
			resp.Body = body
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func gofmt(body string) (string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "prog.go", body, parser.ParseComments)
	if err != nil {
		return "", err
	}
	ast.SortImports(fset, f)
	var buf bytes.Buffer
	config := &printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
	err = config.Fprint(&buf, fset, f)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
