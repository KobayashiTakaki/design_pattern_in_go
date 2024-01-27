package main

import "fmt"

type Print interface {
	PrintWeak()
	PrintStrong()
}

type Banner struct {
	String string
}

func (b *Banner) showWithParen() {
	fmt.Println("(" + b.String + ")")
}

func (b *Banner) showWithAster() {
	fmt.Println("*" + b.String + "*")
}

type PrintBanner struct {
	Banner
}

var _ Print = (*PrintBanner)(nil) // PrintBannerはPrintを満たす

func NewPrintBanner(s string) *PrintBanner {
	return &PrintBanner{
		Banner: Banner{
			String: s,
		},
	}
}

func (b *PrintBanner) PrintWeak() {
	b.showWithParen()
}

func (b *PrintBanner) PrintStrong() {
	b.showWithAster()
}

func main() {
	var p Print = NewPrintBanner("Hello")
	p.PrintWeak()
	p.PrintStrong()
}
