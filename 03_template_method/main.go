package main

import (
	"fmt"
	"strings"
)

type Display interface {
	open()
	print()
	close()
}

type AbstractDisplay struct {
	Display
}

func (d *AbstractDisplay) display() { // これがTemplate methodに位置する
	d.open()
	for i := 0; i < 5; i++ {
		d.print()
	}
	d.close()
}

type CharDisplay struct {
	*AbstractDisplay
	ch string
}

var _ Display = (*CharDisplay)(nil) // CharDisplayはDisplayを満たす

func NewCharDisplay(ch string) *CharDisplay {
	a := &AbstractDisplay{}
	d := &CharDisplay{
		AbstractDisplay: a,
		ch:              ch,
	}
	// AbstractDisplayにはDisplayが実装されていない
	// Displayを実装した型を渡す
	a.Display = d
	return d
}

func (d *CharDisplay) open() {
	fmt.Print("<<")
}

func (d *CharDisplay) print() {
	fmt.Print(d.ch)
}

func (d *CharDisplay) close() {
	fmt.Println(">>")
}

type StringDisplay struct {
	*AbstractDisplay
	str   string
	width int
}

var _ Display = (*StringDisplay)(nil) // StringDisplayはDisplayを満たす

func NewStringDisplay(s string) *StringDisplay {
	a := &AbstractDisplay{}
	d := &StringDisplay{
		AbstractDisplay: a,
		str:             s,
		width:           len(s),
	}
	a.Display = d
	return d
}

func (d *StringDisplay) open() {
	d.printLine()
}

func (d *StringDisplay) print() {
	fmt.Println("|" + d.str + "|")
}

func (d *StringDisplay) close() {
	d.printLine()
}

func (d *StringDisplay) printLine() {
	fmt.Println("+" + strings.Repeat("-", d.width) + "+")
}

func main() {
	d1 := NewCharDisplay("H")
	d2 := NewStringDisplay("Hello, world.")
	d1.display()
	d2.display()
}
