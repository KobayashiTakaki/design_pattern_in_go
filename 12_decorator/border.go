package main

import (
	"fmt"
	"strings"
)

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

type UpDownBorder struct {
	Border
	ch string
}

var _ Display = (*UpDownBorder)(nil)

func NewUpDownBorder(display Display, ch string) *UpDownBorder {
	return &UpDownBorder{
		Border: Border{
			display: display,
		},
		ch: ch,
	}
}

func (b *UpDownBorder) GetColumns() int {
	return b.display.GetColumns()
}

func (b *UpDownBorder) GetRows() int {
	// 行数は中身の行数に上下の飾り文字分を加えたもの
	return 1 + b.display.GetRows() + 1
}

func (b *UpDownBorder) GetRowText(row int) (string, error) {
	if row == 0 { // 上端の枠
		return strings.Repeat(b.ch, b.display.GetColumns()), nil
	} else if row == b.display.GetRows()+1 { // 下端の枠
		return strings.Repeat(b.ch, b.display.GetColumns()), nil
	} else { // それ以外
		t, err := b.display.GetRowText(row - 1)
		if err != nil {
			return "", err
		}
		return t, nil
	}
}

func (b *UpDownBorder) Show() {
	for i := 0; i < b.GetRows(); i++ {
		t, err := b.GetRowText(i)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t)
	}
}
