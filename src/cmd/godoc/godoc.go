// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/printer"
	"go/token"
	"http"
	"io"
	"io/ioutil"
	"log"
	"os"
	pathutil "path"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"template"
	"time"
	"unicode"
	"utf8"
)


// ----------------------------------------------------------------------------
// Support types

// An RWValue wraps a value and permits mutually exclusive
// access to it and records the time the value was last set.
type RWValue struct {
	mutex     sync.RWMutex
	value     interface{}
	timestamp int64 // time of last set(), in seconds since epoch
}


func (v *RWValue) set(value interface{}) {
	v.mutex.Lock()
	v.value = value
	v.timestamp = time.Seconds()
	v.mutex.Unlock()
}


func (v *RWValue) get() (interface{}, int64) {
	v.mutex.RLock()
	defer v.mutex.RUnlock()
	return v.value, v.timestamp
}


// ----------------------------------------------------------------------------
// Globals

type delayTime struct {
	RWValue
}


func (dt *delayTime) backoff(max int) {
	dt.mutex.Lock()
	v := dt.value.(int) * 2
	if v > max {
		v = max
	}
	dt.value = v
	dt.mutex.Unlock()
}


var (
	verbose = flag.Bool("v", false, "verbose mode")

	// file system roots
	goroot = flag.String("goroot", runtime.GOROOT(), "Go root directory")
	path   = flag.String("path", "", "additional package directories (colon-separated)")

	// layout control
	tabwidth = flag.Int("tabwidth", 4, "tab width")

	// file system mapping
	fsMap  Mapping // user-defined mapping
	fsTree RWValue // *Directory tree of packages, updated with each sync

	// http handlers
	fileServer http.Handler // default file server
	cmdHandler httpHandler
	pkgHandler httpHandler
)


func initHandlers() {
	fsMap.Init(*path)
	fileServer = http.FileServer(*goroot, "")
	cmdHandler = httpHandler{"/cmd/", pathutil.Join(*goroot, "src/cmd"), false}
	pkgHandler = httpHandler{"/pkg/", pathutil.Join(*goroot, "src/pkg"), true}
}


func registerPublicHandlers(mux *http.ServeMux) {
	mux.Handle(cmdHandler.pattern, &cmdHandler)
	mux.Handle(pkgHandler.pattern, &pkgHandler)
	mux.HandleFunc("/doc/codewalk/", codewalk)
	mux.HandleFunc("/search", search)
	mux.HandleFunc("/", serveFile)
}


// ----------------------------------------------------------------------------
// Predicates and small utility functions

func isGoFile(f *os.FileInfo) bool {
	return f.IsRegular() &&
		!strings.HasPrefix(f.Name, ".") && // ignore .files
		pathutil.Ext(f.Name) == ".go"
}


func isPkgFile(f *os.FileInfo) bool {
	return isGoFile(f) &&
		!strings.HasSuffix(f.Name, "_test.go") // ignore test files
}


func isPkgDir(f *os.FileInfo) bool {
	return f.IsDirectory() && len(f.Name) > 0 && f.Name[0] != '_'
}


func pkgName(filename string) string {
	file, err := parser.ParseFile(filename, nil, parser.PackageClauseOnly)
	if err != nil || file == nil {
		return ""
	}
	return file.Name.Name
}


func htmlEscape(s string) string {
	var buf bytes.Buffer
	template.HTMLEscape(&buf, []byte(s))
	return buf.String()
}


func firstSentence(s string) string {
	i := -1 // index+1 of first terminator (punctuation ending a sentence)
	j := -1 // index+1 of first terminator followed by white space
	prev := 'A'
	for k, ch := range s {
		k1 := k + 1
		if ch == '.' || ch == '!' || ch == '?' {
			if i < 0 {
				i = k1 // first terminator
			}
			if k1 < len(s) && s[k1] <= ' ' {
				if j < 0 {
					j = k1 // first terminator followed by white space
				}
				if !unicode.IsUpper(prev) {
					j = k1
					break
				}
			}
		}
		prev = ch
	}

	if j < 0 {
		// use the next best terminator
		j = i
		if j < 0 {
			// no terminator at all, use the entire string
			j = len(s)
		}
	}

	return s[0:j]
}


func absolutePath(path, defaultRoot string) string {
	abspath := fsMap.ToAbsolute(path)
	if abspath == "" {
		// no user-defined mapping found; use default mapping
		abspath = pathutil.Join(defaultRoot, path)
	}
	return abspath
}


func relativePath(path string) string {
	relpath := fsMap.ToRelative(path)
	if relpath == "" && strings.HasPrefix(path, *goroot+"/") {
		// no user-defined mapping found; use default mapping
		relpath = path[len(*goroot)+1:]
	}
	// Only if path is an invalid absolute path is relpath == ""
	// at this point. This should never happen since absolute paths
	// are only created via godoc for files that do exist. However,
	// it is ok to return ""; it will simply provide a link to the
	// top of the pkg or src directories.
	return relpath
}


// ----------------------------------------------------------------------------
// Package directories

type Directory struct {
	Depth int
	Path  string // includes Name
	Name  string
	Text  string       // package documentation, if any
	Dirs  []*Directory // subdirectories
}


func newDirTree(path, name string, depth, maxDepth int) *Directory {
	if depth >= maxDepth {
		// return a dummy directory so that the parent directory
		// doesn't get discarded just because we reached the max
		// directory depth
		return &Directory{depth, path, name, "", nil}
	}

	list, _ := ioutil.ReadDir(path) // ignore errors

	// determine number of subdirectories and package files
	ndirs := 0
	nfiles := 0
	var synopses [4]string // prioritized package documentation (0 == highest priority)
	for _, d := range list {
		switch {
		case isPkgDir(d):
			ndirs++
		case isPkgFile(d):
			nfiles++
			if synopses[0] == "" {
				// no "optimal" package synopsis yet; continue to collect synopses
				file, err := parser.ParseFile(pathutil.Join(path, d.Name), nil,
					parser.ParseComments|parser.PackageClauseOnly)
				if err == nil && file.Doc != nil {
					// prioritize documentation
					i := -1
					switch file.Name.Name {
					case name:
						i = 0 // normal case: directory name matches package name
					case fakePkgName:
						i = 1 // synopses for commands
					case "main":
						i = 2 // directory contains a main package
					default:
						i = 3 // none of the above
					}
					if 0 <= i && i < len(synopses) && synopses[i] == "" {
						synopses[i] = firstSentence(doc.CommentText(file.Doc))
					}
				}
			}
		}
	}

	// create subdirectory tree
	var dirs []*Directory
	if ndirs > 0 {
		dirs = make([]*Directory, ndirs)
		i := 0
		for _, d := range list {
			if isPkgDir(d) {
				dd := newDirTree(pathutil.Join(path, d.Name), d.Name, depth+1, maxDepth)
				if dd != nil {
					dirs[i] = dd
					i++
				}
			}
		}
		dirs = dirs[0:i]
	}

	// if there are no package files and no subdirectories
	// (with package files), ignore the directory
	if nfiles == 0 && len(dirs) == 0 {
		return nil
	}

	// select the highest-priority synopsis for the directory entry, if any
	synopsis := ""
	for _, synopsis = range synopses {
		if synopsis != "" {
			break
		}
	}

	return &Directory{depth, path, name, synopsis, dirs}
}


// newDirectory creates a new package directory tree with at most maxDepth
// levels, anchored at root. The result tree is pruned such that it only
// contains directories that contain package files or that contain
// subdirectories containing package files (transitively). If maxDepth is
// too shallow, the leaf nodes are assumed to contain package files even if
// their contents are not known (i.e., in this case the tree may contain
// directories w/o any package files).
//
func newDirectory(root string, maxDepth int) *Directory {
	d, err := os.Lstat(root)
	if err != nil || !isPkgDir(d) {
		return nil
	}
	return newDirTree(root, d.Name, 0, maxDepth)
}


func (dir *Directory) walk(c chan<- *Directory, skipRoot bool) {
	if dir != nil {
		if !skipRoot {
			c <- dir
		}
		for _, d := range dir.Dirs {
			d.walk(c, false)
		}
	}
}


func (dir *Directory) iter(skipRoot bool) <-chan *Directory {
	c := make(chan *Directory)
	go func() {
		dir.walk(c, skipRoot)
		close(c)
	}()
	return c
}


func (dir *Directory) lookupLocal(name string) *Directory {
	for _, d := range dir.Dirs {
		if d.Name == name {
			return d
		}
	}
	return nil
}


// lookup looks for the *Directory for a given path, relative to dir.
func (dir *Directory) lookup(path string) *Directory {
	d := strings.Split(dir.Path, "/", -1)
	p := strings.Split(path, "/", -1)
	i := 0
	for i < len(d) {
		if i >= len(p) || d[i] != p[i] {
			return nil
		}
		i++
	}
	for dir != nil && i < len(p) {
		dir = dir.lookupLocal(p[i])
		i++
	}
	return dir
}


// DirEntry describes a directory entry. The Depth and Height values
// are useful for presenting an entry in an indented fashion.
//
type DirEntry struct {
	Depth    int    // >= 0
	Height   int    // = DirList.MaxHeight - Depth, > 0
	Path     string // includes Name, relative to DirList root
	Name     string
	Synopsis string
}


type DirList struct {
	MaxHeight int // directory tree height, > 0
	List      []DirEntry
}


// listing creates a (linear) directory listing from a directory tree.
// If skipRoot is set, the root directory itself is excluded from the list.
//
func (root *Directory) listing(skipRoot bool) *DirList {
	if root == nil {
		return nil
	}

	// determine number of entries n and maximum height
	n := 0
	minDepth := 1 << 30 // infinity
	maxDepth := 0
	for d := range root.iter(skipRoot) {
		n++
		if minDepth > d.Depth {
			minDepth = d.Depth
		}
		if maxDepth < d.Depth {
			maxDepth = d.Depth
		}
	}
	maxHeight := maxDepth - minDepth + 1

	if n == 0 {
		return nil
	}

	// create list
	list := make([]DirEntry, n)
	i := 0
	for d := range root.iter(skipRoot) {
		p := &list[i]
		p.Depth = d.Depth - minDepth
		p.Height = maxHeight - p.Depth
		// the path is relative to root.Path - remove the root.Path
		// prefix (the prefix should always be present but avoid
		// crashes and check)
		path := d.Path
		if strings.HasPrefix(d.Path, root.Path) {
			path = d.Path[len(root.Path):]
		}
		// remove trailing '/' if any - path must be relative
		if len(path) > 0 && path[0] == '/' {
			path = path[1:]
		}
		p.Path = path
		p.Name = d.Name
		p.Synopsis = d.Text
		i++
	}

	return &DirList{maxHeight, list}
}


// ----------------------------------------------------------------------------
// HTML formatting support

// Styler implements a printer.Styler.
type Styler struct {
	linetags  bool
	highlight string
	objmap    map[*ast.Object]int
	idcount   int
}


func newStyler(highlight string) *Styler {
	return &Styler{true, highlight, make(map[*ast.Object]int), 0}
}


// identId returns a number >= 0 identifying the *ast.Object
// denoted by name. If no object is denoted, the result is < 0.
//
// TODO(gri): Consider making this a mapping from popup info
//            (for that name) to id, instead of *ast.Object
//            to id. If a lot of the popup info is the same
//            (e.g. type information), this will reduce the
//            size of the html generated.
func (s *Styler) identId(name *ast.Ident) int {
	obj := name.Obj
	if obj == nil || s.objmap == nil {
		return -1
	}
	id, found := s.objmap[obj]
	if !found {
		// first occurence
		id = s.idcount
		s.objmap[obj] = id
		s.idcount++
	}
	return id
}


// writeObjInfo writes the popup info corresponding to obj to w.
// The text is HTML-escaped and does not contain single quotes.
func writeObjInfo(w io.Writer, obj *ast.Object) {
	// for now, show object kind and name; eventually
	// do something more interesting (show declaration,
	// for instance)
	if obj.Kind != ast.Bad {
		fmt.Fprintf(w, "%s ", obj.Kind)
	}
	template.HTMLEscape(w, []byte(obj.Name))
}


// idList returns a Javascript array (source) with identifier popup
// information: The i'th array entry is a single-quoted string with
// the popup information for an identifier x with s.identId(x) == i,
// for 0 <= i < s.idcount.
func (s *Styler) idList() []byte {
	var buf bytes.Buffer
	fmt.Fprintln(&buf, "[")

	if s.idcount > 0 {
		// invert objmap: create an array [id]obj from map[obj]id
		a := make([]*ast.Object, s.idcount)
		for obj, id := range s.objmap {
			a[id] = obj
		}

		// for each id, print object info as single-quoted Javascript string
		for id, obj := range a {
			printIndex := false // enable for debugging (but longer html)
			if printIndex {
				fmt.Fprintf(&buf, "/* %4d */ ", id)
			}
			fmt.Fprint(&buf, "'")
			writeObjInfo(&buf, obj)
			fmt.Fprint(&buf, "',\n")
		}
	}

	fmt.Fprintln(&buf, "]")
	return buf.Bytes()
}


// Use the defaultStyler when there is no specific styler.
// The defaultStyler does not emit line tags since they may
// interfere with tags emitted by templates.
// TODO(gri): Should emit line tags at the beginning of a line;
//            never in the middle of code.
var defaultStyler Styler


func (s *Styler) LineTag(line int) (text []byte, tag printer.HTMLTag) {
	if s.linetags {
		tag = printer.HTMLTag{fmt.Sprintf(`<a id="L%d">`, line), "</a>"}
	}
	return
}


func (s *Styler) Comment(c *ast.Comment, line []byte) (text []byte, tag printer.HTMLTag) {
	text = line
	// minimal syntax-coloring of comments for now - people will want more
	// (don't do anything more until there's a button to turn it on/off)
	tag = printer.HTMLTag{`<span class="comment">`, "</span>"}
	return
}


func (s *Styler) BasicLit(x *ast.BasicLit) (text []byte, tag printer.HTMLTag) {
	text = x.Value
	return
}


func (s *Styler) Ident(name *ast.Ident) (text []byte, tag printer.HTMLTag) {
	text = []byte(name.Name)
	var str string
	if id := s.identId(name); id >= 0 {
		str = fmt.Sprintf(` id="%d"`, id)
	}
	if s.highlight == name.Name {
		str += ` class="highlight"`
	}
	if str != "" {
		tag = printer.HTMLTag{"<span" + str + ">", "</span>"}
	}
	return
}


func (s *Styler) Token(tok token.Token) (text []byte, tag printer.HTMLTag) {
	text = []byte(tok.String())
	return
}


// ----------------------------------------------------------------------------
// Tab conversion

var spaces = []byte("                ") // 16 spaces seems like a good number

const (
	indenting = iota
	collecting
)

// A tconv is an io.Writer filter for converting leading tabs into spaces.
type tconv struct {
	output io.Writer
	state  int // indenting or collecting
	indent int // valid if state == indenting
}


func (p *tconv) writeIndent() (err os.Error) {
	i := p.indent
	for i > len(spaces) {
		i -= len(spaces)
		if _, err = p.output.Write(spaces); err != nil {
			return
		}
	}
	_, err = p.output.Write(spaces[0:i])
	return
}


func (p *tconv) Write(data []byte) (n int, err os.Error) {
	pos := 0 // valid if p.state == collecting
	var b byte
	for n, b = range data {
		switch p.state {
		case indenting:
			switch b {
			case '\t', '\v':
				p.indent += *tabwidth
			case '\n':
				p.indent = 0
				if _, err = p.output.Write(data[n : n+1]); err != nil {
					return
				}
			case ' ':
				p.indent++
			default:
				p.state = collecting
				pos = n
				if err = p.writeIndent(); err != nil {
					return
				}
			}
		case collecting:
			if b == '\n' {
				p.state = indenting
				p.indent = 0
				if _, err = p.output.Write(data[pos : n+1]); err != nil {
					return
				}
			}
		}
	}
	n = len(data)
	if p.state == collecting {
		_, err = p.output.Write(data[pos:])
	}
	return
}


// ----------------------------------------------------------------------------
// Templates

// Write an AST-node to w; optionally html-escaped.
func writeNode(w io.Writer, node interface{}, html bool, styler printer.Styler) {
	mode := printer.TabIndent | printer.UseSpaces
	if html {
		mode |= printer.GenHTML
	}
	// convert trailing tabs into spaces using a tconv filter
	// to ensure a good outcome in most browsers (there may still
	// be tabs in comments and strings, but converting those into
	// the right number of spaces is much harder)
	(&printer.Config{mode, *tabwidth, styler}).Fprint(&tconv{output: w}, node)
}


// Write text to w; optionally html-escaped.
func writeText(w io.Writer, text []byte, html bool) {
	if html {
		template.HTMLEscape(w, text)
		return
	}
	w.Write(text)
}


// Write anything to w; optionally html-escaped.
func writeAny(w io.Writer, x interface{}, html bool) {
	switch v := x.(type) {
	case []byte:
		writeText(w, v, html)
	case string:
		writeText(w, []byte(v), html)
	case ast.Decl, ast.Expr, ast.Stmt, *ast.File:
		writeNode(w, x, html, &defaultStyler)
	default:
		if html {
			var buf bytes.Buffer
			fmt.Fprint(&buf, x)
			writeText(w, buf.Bytes(), true)
		} else {
			fmt.Fprint(w, x)
		}
	}
}


// Template formatter for "html" format.
func htmlFmt(w io.Writer, x interface{}, format string) {
	writeAny(w, x, true)
}


// Template formatter for "html-esc" format.
func htmlEscFmt(w io.Writer, x interface{}, format string) {
	var buf bytes.Buffer
	writeAny(&buf, x, false)
	template.HTMLEscape(w, buf.Bytes())
}


// Template formatter for "html-comment" format.
func htmlCommentFmt(w io.Writer, x interface{}, format string) {
	var buf bytes.Buffer
	writeAny(&buf, x, false)
	// TODO(gri) Provide list of words (e.g. function parameters)
	//           to be emphasized by ToHTML.
	doc.ToHTML(w, buf.Bytes(), nil) // does html-escaping
}


// Template formatter for "" (default) format.
func textFmt(w io.Writer, x interface{}, format string) {
	writeAny(w, x, false)
}


// Template formatter for the various "url-xxx" formats.
func urlFmt(w io.Writer, x interface{}, format string) {
	var path string
	var line int

	// determine path and position info, if any
	type positioner interface {
		Pos() token.Position
	}
	switch t := x.(type) {
	case string:
		path = t
	case positioner:
		pos := t.Pos()
		if pos.IsValid() {
			path = pos.Filename
			line = pos.Line
		}
	}

	// map path
	relpath := relativePath(path)

	// convert to relative URLs so that they can also
	// be used as relative file names in .txt templates
	switch format {
	default:
		// we should never reach here, but be resilient
		// and assume the url-pkg format instead
		log.Stderrf("INTERNAL ERROR: urlFmt(%s)", format)
		fallthrough
	case "url-pkg":
		// because of the irregular mapping under goroot
		// we need to correct certain relative paths
		if strings.HasPrefix(relpath, "src/pkg/") {
			relpath = relpath[len("src/pkg/"):]
		}
		template.HTMLEscape(w, []byte(pkgHandler.pattern[1:]+relpath)) // remove trailing '/' for relative URL
	case "url-src":
		template.HTMLEscape(w, []byte(relpath))
	case "url-pos":
		// line id's in html-printed source are of the
		// form "L%d" where %d stands for the line number
		template.HTMLEscape(w, []byte(relpath))
		fmt.Fprintf(w, "#L%d", line)
	}
}


// The strings in infoKinds must be properly html-escaped.
var infoKinds = [nKinds]string{
	PackageClause: "package&nbsp;clause",
	ImportDecl:    "import&nbsp;decl",
	ConstDecl:     "const&nbsp;decl",
	TypeDecl:      "type&nbsp;decl",
	VarDecl:       "var&nbsp;decl",
	FuncDecl:      "func&nbsp;decl",
	MethodDecl:    "method&nbsp;decl",
	Use:           "use",
}


// Template formatter for "infoKind" format.
func infoKindFmt(w io.Writer, x interface{}, format string) {
	fmt.Fprintf(w, infoKinds[x.(SpotKind)]) // infoKind entries are html-escaped
}


// Template formatter for "infoLine" format.
func infoLineFmt(w io.Writer, x interface{}, format string) {
	info := x.(SpotInfo)
	line := info.Lori()
	if info.IsIndex() {
		index, _ := searchIndex.get()
		if index != nil {
			line = index.(*Index).Snippet(line).Line
		} else {
			// no line information available because
			// we don't have an index - this should
			// never happen; be conservative and don't
			// crash
			line = 0
		}
	}
	fmt.Fprintf(w, "%d", line)
}


// Template formatter for "infoSnippet" format.
func infoSnippetFmt(w io.Writer, x interface{}, format string) {
	info := x.(SpotInfo)
	text := `<span class="alert">no snippet text available</span>`
	if info.IsIndex() {
		index, _ := searchIndex.get()
		// no escaping of snippet text needed;
		// snippet text is escaped when generated
		text = index.(*Index).Snippet(info.Lori()).Text
	}
	fmt.Fprint(w, text)
}


// Template formatter for "padding" format.
func paddingFmt(w io.Writer, x interface{}, format string) {
	for i := x.(int); i > 0; i-- {
		fmt.Fprint(w, `<td width="25"></td>`)
	}
}


// Template formatter for "time" format.
func timeFmt(w io.Writer, x interface{}, format string) {
	template.HTMLEscape(w, []byte(time.SecondsToLocalTime(x.(int64)/1e9).String()))
}


// Template formatter for "dir/" format.
func dirslashFmt(w io.Writer, x interface{}, format string) {
	if x.(*os.FileInfo).IsDirectory() {
		w.Write([]byte{'/'})
	}
}


// Template formatter for "localname" format.
func localnameFmt(w io.Writer, x interface{}, format string) {
	_, localname := pathutil.Split(x.(string))
	template.HTMLEscape(w, []byte(localname))
}


var fmap = template.FormatterMap{
	"":             textFmt,
	"html":         htmlFmt,
	"html-esc":     htmlEscFmt,
	"html-comment": htmlCommentFmt,
	"url-pkg":      urlFmt,
	"url-src":      urlFmt,
	"url-pos":      urlFmt,
	"infoKind":     infoKindFmt,
	"infoLine":     infoLineFmt,
	"infoSnippet":  infoSnippetFmt,
	"padding":      paddingFmt,
	"time":         timeFmt,
	"dir/":         dirslashFmt,
	"localname":    localnameFmt,
}


func readTemplate(name string) *template.Template {
	path := pathutil.Join(*goroot, "lib/godoc/"+name)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Exitf("ReadFile %s: %v", path, err)
	}
	t, err := template.Parse(string(data), fmap)
	if err != nil {
		log.Exitf("%s: %v", name, err)
	}
	return t
}


var (
	codewalkHTML,
	codewalkdirHTML,
	dirlistHTML,
	errorHTML,
	godocHTML,
	packageHTML,
	packageText,
	searchHTML,
	searchText,
	sourceHTML *template.Template
)

func readTemplates() {
	// have to delay until after flags processing since paths depend on goroot
	codewalkHTML = readTemplate("codewalk.html")
	codewalkdirHTML = readTemplate("codewalkdir.html")
	dirlistHTML = readTemplate("dirlist.html")
	errorHTML = readTemplate("error.html")
	godocHTML = readTemplate("godoc.html")
	packageHTML = readTemplate("package.html")
	packageText = readTemplate("package.txt")
	searchHTML = readTemplate("search.html")
	searchText = readTemplate("search.txt")
	sourceHTML = readTemplate("source.html")
}


// ----------------------------------------------------------------------------
// Generic HTML wrapper

func servePage(c *http.Conn, title, subtitle, query string, content []byte) {
	type Data struct {
		Title     string
		Subtitle  string
		PkgRoots  []string
		Timestamp int64
		Query     string
		Version   string
		Menu      []byte
		Content   []byte
	}

	_, ts := fsTree.get()
	d := Data{
		Title:     title,
		Subtitle:  subtitle,
		PkgRoots:  fsMap.PrefixList(),
		Timestamp: ts * 1e9, // timestamp in ns
		Query:     query,
		Version:   runtime.Version(),
		Menu:      nil,
		Content:   content,
	}

	if err := godocHTML.Execute(&d, c); err != nil {
		log.Stderrf("godocHTML.Execute: %s", err)
	}
}


func serveText(c *http.Conn, text []byte) {
	c.SetHeader("Content-Type", "text/plain; charset=utf-8")
	c.Write(text)
}


// ----------------------------------------------------------------------------
// Files

var (
	titleRx        = regexp.MustCompile(`<!-- title ([^\-]*)-->`)
	subtitleRx     = regexp.MustCompile(`<!-- subtitle ([^\-]*)-->`)
	firstCommentRx = regexp.MustCompile(`<!--([^\-]*)-->`)
)


func extractString(src []byte, rx *regexp.Regexp) (s string) {
	m := rx.FindSubmatch(src)
	if m != nil {
		s = strings.TrimSpace(string(m[1]))
	}
	return
}


func serveHTMLDoc(c *http.Conn, r *http.Request, abspath, relpath string) {
	// get HTML body contents
	src, err := ioutil.ReadFile(abspath)
	if err != nil {
		log.Stderrf("ioutil.ReadFile: %s", err)
		serveError(c, r, relpath, err)
		return
	}

	// if it begins with "<!DOCTYPE " assume it is standalone
	// html that doesn't need the template wrapping.
	if bytes.HasPrefix(src, []byte("<!DOCTYPE ")) {
		c.Write(src)
		return
	}

	// if it's the language spec, add tags to EBNF productions
	if strings.HasSuffix(abspath, "go_spec.html") {
		var buf bytes.Buffer
		linkify(&buf, src)
		src = buf.Bytes()
	}

	// get title and subtitle, if any
	title := extractString(src, titleRx)
	if title == "" {
		// no title found; try first comment for backward-compatibility
		title = extractString(src, firstCommentRx)
	}
	subtitle := extractString(src, subtitleRx)

	servePage(c, title, subtitle, "", src)
}


func applyTemplate(t *template.Template, name string, data interface{}) []byte {
	var buf bytes.Buffer
	if err := t.Execute(data, &buf); err != nil {
		log.Stderrf("%s.Execute: %s", name, err)
	}
	return buf.Bytes()
}


func serveGoSource(c *http.Conn, r *http.Request, abspath, relpath string) {
	file, err := parser.ParseFile(abspath, nil, parser.ParseComments)
	if err != nil {
		log.Stderrf("parser.ParseFile: %s", err)
		serveError(c, r, relpath, err)
		return
	}

	// augment AST with types; ignore errors (partial type information ok)
	// TODO(gri): invoke typechecker

	var buf bytes.Buffer
	styler := newStyler(r.FormValue("h"))
	writeNode(&buf, file, true, styler)

	type SourceInfo struct {
		IdList []byte
		Source []byte
	}
	info := &SourceInfo{styler.idList(), buf.Bytes()}

	contents := applyTemplate(sourceHTML, "sourceHTML", info)
	servePage(c, "Source file "+relpath, "", "", contents)
}


func redirect(c *http.Conn, r *http.Request) (redirected bool) {
	if canonical := pathutil.Clean(r.URL.Path) + "/"; r.URL.Path != canonical {
		http.Redirect(c, canonical, http.StatusMovedPermanently)
		redirected = true
	}
	return
}


// TODO(gri): Should have a mapping from extension to handler, eventually.

// textExt[x] is true if the extension x indicates a text file, and false otherwise.
var textExt = map[string]bool{
	".css": false, // must be served raw
	".js":  false, // must be served raw
}


func isTextFile(path string) bool {
	// if the extension is known, use it for decision making
	if isText, found := textExt[pathutil.Ext(path)]; found {
		return isText
	}

	// the extension is not known; read an initial chunk of
	// file and check if it looks like correct UTF-8; if it
	// does, it's probably a text file
	f, err := os.Open(path, os.O_RDONLY, 0)
	if err != nil {
		return false
	}
	defer f.Close()

	var buf [1024]byte
	n, err := f.Read(buf[0:])
	if err != nil {
		return false
	}

	s := string(buf[0:n])
	n -= utf8.UTFMax // make sure there's enough bytes for a complete unicode char
	for i, c := range s {
		if i > n {
			break
		}
		if c == 0xFFFD || c < ' ' && c != '\n' && c != '\t' {
			// decoding error or control character - not a text file
			return false
		}
	}

	// likely a text file
	return true
}


func serveTextFile(c *http.Conn, r *http.Request, abspath, relpath string) {
	src, err := ioutil.ReadFile(abspath)
	if err != nil {
		log.Stderrf("ioutil.ReadFile: %s", err)
		serveError(c, r, relpath, err)
		return
	}

	var buf bytes.Buffer
	fmt.Fprintln(&buf, "<pre>")
	template.HTMLEscape(&buf, src)
	fmt.Fprintln(&buf, "</pre>")

	servePage(c, "Text file "+relpath, "", "", buf.Bytes())
}


func serveDirectory(c *http.Conn, r *http.Request, abspath, relpath string) {
	if redirect(c, r) {
		return
	}

	list, err := ioutil.ReadDir(abspath)
	if err != nil {
		log.Stderrf("ioutil.ReadDir: %s", err)
		serveError(c, r, relpath, err)
		return
	}

	for _, d := range list {
		if d.IsDirectory() {
			d.Size = 0
		}
	}

	contents := applyTemplate(dirlistHTML, "dirlistHTML", list)
	servePage(c, "Directory "+relpath, "", "", contents)
}


func serveFile(c *http.Conn, r *http.Request) {
	relpath := r.URL.Path[1:] // serveFile URL paths start with '/'
	abspath := absolutePath(relpath, *goroot)

	// pick off special cases and hand the rest to the standard file server
	switch r.URL.Path {
	case "/":
		serveHTMLDoc(c, r, pathutil.Join(*goroot, "doc/root.html"), "doc/root.html")
		return

	case "/doc/root.html":
		// hide landing page from its real name
		http.Redirect(c, "/", http.StatusMovedPermanently)
		return
	}

	switch pathutil.Ext(abspath) {
	case ".html":
		if strings.HasSuffix(abspath, "/index.html") {
			// We'll show index.html for the directory.
			// Use the dir/ version as canonical instead of dir/index.html.
			http.Redirect(c, r.URL.Path[0:len(r.URL.Path)-len("index.html")], http.StatusMovedPermanently)
			return
		}
		serveHTMLDoc(c, r, abspath, relpath)
		return

	case ".go":
		serveGoSource(c, r, abspath, relpath)
		return
	}

	dir, err := os.Lstat(abspath)
	if err != nil {
		log.Stderr(err)
		serveError(c, r, relpath, err)
		return
	}

	if dir != nil && dir.IsDirectory() {
		if redirect(c, r) {
			return
		}
		if index := abspath + "/index.html"; isTextFile(index) {
			serveHTMLDoc(c, r, index, relativePath(index))
			return
		}
		serveDirectory(c, r, abspath, relpath)
		return
	}

	if isTextFile(abspath) {
		serveTextFile(c, r, abspath, relpath)
		return
	}

	fileServer.ServeHTTP(c, r)
}


// ----------------------------------------------------------------------------
// Packages

// Fake package file and name for commands. Contains the command documentation.
const fakePkgFile = "doc.go"
const fakePkgName = "documentation"

type PageInfoMode uint

const (
	exportsOnly PageInfoMode = 1 << iota // only keep exported stuff
	genDoc                               // generate documentation
)


type PageInfo struct {
	Dirname string          // directory containing the package
	PList   []string        // list of package names found
	PAst    *ast.File       // nil if no single AST with package exports
	PDoc    *doc.PackageDoc // nil if no single package documentation
	Dirs    *DirList        // nil if no directory information
	IsPkg   bool            // false if this is not documenting a real package
	Err     os.Error        // directory read error or nil
}


type httpHandler struct {
	pattern string // url pattern; e.g. "/pkg/"
	fsRoot  string // file system root to which the pattern is mapped
	isPkg   bool   // true if this handler serves real package documentation (as opposed to command documentation)
}


// getPageInfo returns the PageInfo for a package directory abspath. If the
// parameter genAST is set, an AST containing only the package exports is
// computed (PageInfo.PAst), otherwise package documentation (PageInfo.Doc)
// is extracted from the AST. If there is no corresponding package in the
// directory, PageInfo.PAst and PageInfo.PDoc are nil. If there are no sub-
// directories, PageInfo.Dirs is nil. If a directory read error occured,
// PageInfo.Err is set to the respective error but the error is not logged.
//
func (h *httpHandler) getPageInfo(abspath, relpath, pkgname string, mode PageInfoMode) PageInfo {
	// filter function to select the desired .go files
	filter := func(d *os.FileInfo) bool {
		// If we are looking at cmd documentation, only accept
		// the special fakePkgFile containing the documentation.
		return isPkgFile(d) && (h.isPkg || d.Name == fakePkgFile)
	}

	// get package ASTs
	pkgs, err := parser.ParseDir(abspath, filter, parser.ParseComments)
	if err != nil && pkgs == nil {
		// only report directory read errors, ignore parse errors
		// (may be able to extract partial package information)
		return PageInfo{Dirname: abspath, Err: err}
	}

	// select package
	var pkg *ast.Package // selected package
	var plist []string   // list of other package (names), if any
	if len(pkgs) == 1 {
		// Exactly one package - select it.
		for _, p := range pkgs {
			pkg = p
		}

	} else if len(pkgs) > 1 {
		// Multiple packages - select the best matching package: The
		// 1st choice is the package with pkgname, the 2nd choice is
		// the package with dirname, and the 3rd choice is a package
		// that is not called "main" if there is exactly one such
		// package. Otherwise, don't select a package.
		dirpath, dirname := pathutil.Split(abspath)

		// If the dirname is "go" we might be in a sub-directory for
		// .go files - use the outer directory name instead for better
		// results.
		if dirname == "go" {
			_, dirname = pathutil.Split(pathutil.Clean(dirpath))
		}

		var choice3 *ast.Package
	loop:
		for _, p := range pkgs {
			switch {
			case p.Name == pkgname:
				pkg = p
				break loop // 1st choice; we are done
			case p.Name == dirname:
				pkg = p // 2nd choice
			case p.Name != "main":
				choice3 = p
			}
		}
		if pkg == nil && len(pkgs) == 2 {
			pkg = choice3
		}

		// Compute the list of other packages
		// (excluding the selected package, if any).
		plist = make([]string, len(pkgs))
		i := 0
		for name, _ := range pkgs {
			if pkg == nil || name != pkg.Name {
				plist[i] = name
				i++
			}
		}
		plist = plist[0:i]
	}

	// compute package documentation
	var past *ast.File
	var pdoc *doc.PackageDoc
	if pkg != nil {
		if mode&exportsOnly != 0 {
			ast.PackageExports(pkg)
		}
		if mode&genDoc != 0 {
			pdoc = doc.NewPackageDoc(pkg, pathutil.Clean(relpath)) // no trailing '/' in importpath
		} else {
			past = ast.MergePackageFiles(pkg, ast.FilterUnassociatedComments)
		}
	}

	// get directory information
	var dir *Directory
	if tree, _ := fsTree.get(); tree != nil && tree.(*Directory) != nil {
		// directory tree is present; lookup respective directory
		// (may still fail if the file system was updated and the
		// new directory tree has not yet been computed)
		// TODO(gri) Need to build directory tree for fsMap entries
		dir = tree.(*Directory).lookup(abspath)
	}
	if dir == nil {
		// no directory tree present (either early after startup
		// or command-line mode, or we don't build a tree for the
		// directory; e.g. google3); compute one level for this page
		dir = newDirectory(abspath, 1)
	}

	return PageInfo{abspath, plist, past, pdoc, dir.listing(true), h.isPkg, nil}
}


func (h *httpHandler) ServeHTTP(c *http.Conn, r *http.Request) {
	if redirect(c, r) {
		return
	}

	relpath := r.URL.Path[len(h.pattern):]
	abspath := absolutePath(relpath, h.fsRoot)
	mode := exportsOnly
	if r.FormValue("m") != "src" {
		mode |= genDoc
	}
	info := h.getPageInfo(abspath, relpath, r.FormValue("p"), mode)
	if info.Err != nil {
		log.Stderr(info.Err)
		serveError(c, r, relpath, info.Err)
		return
	}

	if r.FormValue("f") == "text" {
		contents := applyTemplate(packageText, "packageText", info)
		serveText(c, contents)
		return
	}

	var title string
	switch {
	case info.PAst != nil:
		title = "Package " + info.PAst.Name.Name
	case info.PDoc != nil:
		switch {
		case h.isPkg:
			title = "Package " + info.PDoc.PackageName
		case info.PDoc.PackageName == fakePkgName:
			// assume that the directory name is the command name
			_, pkgname := pathutil.Split(pathutil.Clean(relpath))
			title = "Command " + pkgname
		default:
			title = "Command " + info.PDoc.PackageName
		}
	default:
		title = "Directory " + relativePath(info.Dirname)
	}

	contents := applyTemplate(packageHTML, "packageHTML", info)
	servePage(c, title, "", "", contents)
}


// ----------------------------------------------------------------------------
// Search

var searchIndex RWValue

type SearchResult struct {
	Query    string
	Hit      *LookupResult
	Alt      *AltWords
	Illegal  bool
	Accurate bool
}


func lookup(query string) (result SearchResult) {
	result.Query = query
	if index, timestamp := searchIndex.get(); index != nil {
		result.Hit, result.Alt, result.Illegal = index.(*Index).Lookup(query)
		_, ts := fsTree.get()
		result.Accurate = timestamp >= ts
	}
	return
}


func search(c *http.Conn, r *http.Request) {
	query := strings.TrimSpace(r.FormValue("q"))
	result := lookup(query)

	if r.FormValue("f") == "text" {
		contents := applyTemplate(searchText, "searchText", result)
		serveText(c, contents)
		return
	}

	var title string
	if result.Hit != nil {
		title = fmt.Sprintf(`Results for query %q`, query)
	} else {
		title = fmt.Sprintf(`No results found for query %q`, query)
	}

	contents := applyTemplate(searchHTML, "searchHTML", result)
	servePage(c, title, "", query, contents)
}


// ----------------------------------------------------------------------------
// Indexer

func indexer() {
	for {
		_, ts := fsTree.get()
		if _, timestamp := searchIndex.get(); timestamp < ts {
			// index possibly out of date - make a new one
			// (could use a channel to send an explicit signal
			// from the sync goroutine, but this solution is
			// more decoupled, trivial, and works well enough)
			start := time.Nanoseconds()
			index := NewIndex(*goroot)
			stop := time.Nanoseconds()
			searchIndex.set(index)
			if *verbose {
				secs := float64((stop-start)/1e6) / 1e3
				nwords, nspots := index.Size()
				log.Stderrf("index updated (%gs, %d unique words, %d spots)", secs, nwords, nspots)
			}
			log.Stderrf("bytes=%d footprint=%d\n", runtime.MemStats.HeapAlloc, runtime.MemStats.Sys)
			runtime.GC()
			log.Stderrf("bytes=%d footprint=%d\n", runtime.MemStats.HeapAlloc, runtime.MemStats.Sys)
		}
		time.Sleep(1 * 60e9) // try once a minute
	}
}
