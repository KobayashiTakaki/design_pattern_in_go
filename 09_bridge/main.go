package main

import (
	"fmt"
	"strings"
)

type DisplayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}

type Display struct {
	impl DisplayImpl
}

func NewDisplay(impl DisplayImpl) *Display {
	return &Display{
		impl: impl,
	}
}

func (d *Display) open() {
	d.impl.rawOpen()
}

func (d *Display) print() {
	d.impl.rawPrint()
}

func (d *Display) close() {
	d.impl.rawClose()
}

func (d *Display) display() {
	d.open()
	d.print()
	d.close()
}

type CountDisplay struct {
	Display
}

func NewCountDisplay(impl DisplayImpl) *CountDisplay {
	return &CountDisplay{
		Display: Display{
			impl: impl,
		},
	}
}

func (d *CountDisplay) multiDisplay(times int) {
	d.open()
	for i := 0; i < times; i++ {
		d.print()
	}
	d.close()
}

type StringDisplayImpl struct {
	str   string
	width int
}

func NewStringDisplayImpl(str string) *StringDisplayImpl {
	return &StringDisplayImpl{
		str:   str,
		width: len(str),
	}
}

var _ DisplayImpl = (*StringDisplayImpl)(nil)

func (d *StringDisplayImpl) rawOpen() {
	fmt.Println("+" + strings.Repeat("-", d.width) + "+")
}

func (d *StringDisplayImpl) rawPrint() {
	fmt.Println("|" + d.str + "|")
}

func (d *StringDisplayImpl) rawClose() {
	fmt.Println("+" + strings.Repeat("-", d.width) + "+")
}

func main() {
	d1 := NewDisplay(NewStringDisplayImpl("Hello, Japan."))
	d2 := NewCountDisplay(NewStringDisplayImpl("Hello, World."))
	d3 := NewCountDisplay(NewStringDisplayImpl("Hello, Universe."))
	d1.display()
	d2.display()
	d3.display()
	d3.multiDisplay(5)
}
