package main

import (
	"errors"
	"fmt"
	"strings"
)

type Display interface {
	// 横の文字数を得る
	GetColumns() int
	// 縦の行数を得る
	GetRows() int
	// row行目の文字列を得る
	GetRowText(row int) (string, error)
	// すべての行を表示する
	Show()
}

type StringDisplay struct {
	// 表示文字列
	String string
}

var _ Display = (*StringDisplay)(nil) // StringDisplayはDisplayを満たす

func NewStringDisplay(s string) *StringDisplay {
	return &StringDisplay{
		String: s,
	}
}

func (d *StringDisplay) GetColumns() int {
	return len(d.String)
}

func (d *StringDisplay) GetRows() int {
	return 1 // 行数は1
}

func (d *StringDisplay) GetRowText(row int) (string, error) {
	if row != 0 {
		return "", errors.New("out of bounds")
	}
	return d.String, nil
}

func (d *StringDisplay) Show() {
	for i := 0; i < d.GetRows(); i++ {
		t, err := d.GetRowText(i)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t)
	}
}

type Border struct {
	// この飾り枠がくるんでいる「中身」
	display Display
}

type SideBorder struct {
	Border
	// 飾り文字
	ch string
}

var _ Display = (*SideBorder)(nil) // SideBorderはDisplayを満たす

func NewSideBorder(display Display, ch string) *SideBorder {
	return &SideBorder{
		Border: Border{
			display: display,
		},
		ch: ch,
	}
}

func (b *SideBorder) GetColumns() int {
	// 文字数は中身の両側に左右の飾り文字分を加えたもの
	return 1 + b.display.GetColumns() + 1
}

func (b *SideBorder) GetRows() int {
	// 行数は中身の行数に同じ
	return b.display.GetRows()
}

func (b *SideBorder) GetRowText(row int) (string, error) {
	t, err := b.display.GetRowText(row)
	if err != nil {
		return "", err
	}
	// 指定業の内容は、中身の指定業の両側に飾りをつけたもの
	return b.ch + t + b.ch, nil
}

func (b *SideBorder) Show() {
	for i := 0; i < b.GetRows(); i++ {
		t, err := b.GetRowText(i)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t)
	}
}

type FullBorder struct {
	Border
}

var _ Display = (*FullBorder)(nil) // SideBorderはDisplayを満たす

func NewFullBorder(display Display) *FullBorder {
	return &FullBorder{
		Border: Border{
			display: display,
		},
	}
}

func (b *FullBorder) GetColumns() int {
	// 文字数は中身の両側に左右の飾り文字分を加えたもの
	return 1 + b.display.GetColumns() + 1
}

func (b *FullBorder) GetRows() int {
	// 行数は中身の行数に上下の飾り文字分を加えたもの
	return 1 + b.display.GetRows() + 1
}

func (b *FullBorder) GetRowText(row int) (string, error) {
	if row == 0 { // 上端の枠
		return "+" + strings.Repeat("-", b.display.GetColumns()) + "+", nil
	} else if row == b.display.GetRows()+1 { // 下端の枠
		return "+" + strings.Repeat("-", b.display.GetColumns()) + "+", nil
	} else { // それ以外
		t, err := b.display.GetRowText(row - 1)
		if err != nil {
			return "", err
		}
		return "|" + t + "|", nil
	}
}

func (b *FullBorder) Show() {
	for i := 0; i < b.GetRows(); i++ {
		t, err := b.GetRowText(i)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t)
	}
}

func main() {
	b1 := NewStringDisplay("Hello, world.")
	b2 := NewSideBorder(b1, "#")
	b3 := NewFullBorder(b2)
	b1.Show()
	b2.Show()
	b3.Show()
	b4 := NewSideBorder(
		NewFullBorder(
			NewFullBorder(
				NewSideBorder(
					NewStringDisplay("Hello, world"),
					"*",
				),
			),
		),
		"/",
	)
	b4.Show()
}
