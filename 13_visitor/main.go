package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Visitor interface {
	visitFile(*File)
	visitDirectory(*Directory)
}

type Element interface {
	accept(Visitor)
}

type IEntry interface {
	Element
	getName() string
	getSize() int
}

type Entry struct {
	IEntry
}

func (e *Entry) toString() string {
	return e.getName() + " (" + strconv.Itoa(e.getSize()) + ")"
}

type File struct {
	*Entry
	name string
	size int
}

func NewFile(name string, size int) *File {
	e := &Entry{}
	f := &File{
		Entry: e,
		name:  name,
		size:  size,
	}
	e.IEntry = f
	return f
}

func (f *File) getName() string {
	return f.name
}

func (f *File) getSize() int {
	return f.size
}

func (f *File) accept(v Visitor) {
	v.visitFile(f)
}

type Directory struct {
	*Entry
	name      string
	directory []IEntry
}

func NewDirectory(name string) *Directory {
	e := &Entry{}
	d := &Directory{
		Entry: e,
		name:  name,
	}
	e.IEntry = d
	return d
}

func (d *Directory) getName() string {
	return d.name
}

func (d *Directory) getSize() int {
	var size int
	for _, e := range d.directory {
		size += e.getSize()
	}
	return size
}

func (d *Directory) add(entry IEntry) {
	d.directory = append(d.directory, entry)
}

func (d *Directory) accept(v Visitor) {
	v.visitDirectory(d)
}

type ListVisitor struct {
	currentDir string
}

var _ Visitor = (*ListVisitor)(nil)

func NewListVisitor() *ListVisitor {
	return &ListVisitor{}
}

func (v *ListVisitor) visitFile(f *File) {
	fmt.Println(v.currentDir + "/" + f.toString())
}

func (v *ListVisitor) visitDirectory(d *Directory) {
	fmt.Println(v.currentDir + "/" + d.toString())
	saveDir := v.currentDir
	v.currentDir = v.currentDir + "/" + d.getName()
	for _, entry := range d.directory {
		entry.accept(v)
	}
	v.currentDir = saveDir
}

type FileFindVisitor struct {
	pattern    *regexp.Regexp
	foundFiles []*File
}

var _ Visitor = (*FileFindVisitor)(nil)

func NewFileFindVisitor(pattern string) *FileFindVisitor {
	re := regexp.MustCompile(fmt.Sprintf(`%s$`, pattern))
	return &FileFindVisitor{
		pattern: re,
	}
}

func (v *FileFindVisitor) visitFile(f *File) {
	if v.pattern.MatchString(f.getName()) {
		v.foundFiles = append(v.foundFiles, f)
	}
}

func (v *FileFindVisitor) visitDirectory(d *Directory) {
	for _, entry := range d.directory {
		entry.accept(v)
	}
}

func (v *FileFindVisitor) getFoundFiles() []*File {
	return v.foundFiles
}

func main() {
	fmt.Println("Making root entries...")
	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	tmpDir := NewDirectory("tmp")
	usrDir := NewDirectory("usr")
	rootDir.add(binDir)
	rootDir.add(tmpDir)
	rootDir.add(usrDir)
	binDir.add(NewFile("vi", 10000))
	binDir.add(NewFile("latex", 20000))
	rootDir.accept(NewListVisitor())
	fmt.Println("")

	fmt.Println("Making user entries...")
	yuki := NewDirectory("yuki")
	hanako := NewDirectory("hanako")
	tomura := NewDirectory("tomura")
	usrDir.add(yuki)
	usrDir.add(hanako)
	usrDir.add(tomura)
	yuki.add(NewFile("diary.html", 100))
	yuki.add(NewFile("Composite.java", 200))
	hanako.add(NewFile("memo.tex", 300))
	tomura.add(NewFile("game.doc", 400))
	tomura.add(NewFile("junk.mail", 500))
	rootDir.accept(NewListVisitor())

	ffv := NewFileFindVisitor(".html")
	rootDir.accept(ffv)
	fmt.Println("HTML files are:")
	for _, f := range ffv.getFoundFiles() {
		fmt.Println(f.toString())
	}
}
