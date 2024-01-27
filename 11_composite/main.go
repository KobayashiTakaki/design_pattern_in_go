package main

import (
	"fmt"
	"strconv"
)

type Entry interface {
	getName() string
	getSize() int
	printList(prefix string)
}

type AbstractEntry struct {
	Entry
}

func (e *AbstractEntry) String() string {
	return e.getName() + " (" + strconv.Itoa(e.getSize()) + ")"
}

type File struct {
	*AbstractEntry
	name string
	size int
}

func NewFile(name string, size int) *File {
	a := &AbstractEntry{}
	f := &File{
		AbstractEntry: a,
		name:          name,
		size:          size,
	}
	a.Entry = f
	return f
}

var _ Entry = (*File)(nil)

func (f *File) getName() string {
	return f.name
}

func (f *File) getSize() int {
	return f.size
}

func (f *File) printList(prefix string) {
	fmt.Println(prefix + "/" + f.String())
}

type Directory struct {
	*AbstractEntry
	name      string
	directory []Entry
}

func NewDirectory(name string) *Directory {
	a := &AbstractEntry{}
	d := &Directory{
		AbstractEntry: a,
		name:          name,
	}
	a.Entry = d
	return d
}

var _ Entry = (*Directory)(nil)

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

func (d *Directory) printList(prefix string) {
	fmt.Println(prefix + "/" + d.String())
	for _, e := range d.directory {
		e.printList(prefix + "/" + d.name)
	}
}

func (d *Directory) add(entry Entry) {
	d.directory = append(d.directory, entry)
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
	rootDir.printList("")
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
	rootDir.printList("")
}
