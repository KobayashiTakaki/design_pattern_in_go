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

type MultiStringDisplay struct {
	// 表示する複数行の文字列
	Lines []string
}

var _ Display = (*MultiStringDisplay)(nil) // MultiStringDisplayはDisplayを満たす

func NewMultiStringDisplay() *MultiStringDisplay {
	return &MultiStringDisplay{
		Lines: []string{},
	}
}

func (d *MultiStringDisplay) GetColumns() int {
	// すべての行の中で最大の文字数を返す
	var maxLen int
	for _, line := range d.Lines {
		if maxLen < len(line) {
			maxLen = len(line)
		}
	}
	return maxLen
}

func (d *MultiStringDisplay) GetRows() int {
	return len(d.Lines)
}

func (d *MultiStringDisplay) GetRowText(row int) (string, error) {
	if row > len(d.Lines)-1 {
		return "", errors.New("out of bounds")
	}
	// 行の長さが揃うように後ろにスペースを入れる
	pad := strings.Repeat(" ", d.GetColumns()-len(d.Lines[row]))
	return d.Lines[row] + pad, nil
}

func (d *MultiStringDisplay) Show() {
	for i := 0; i < d.GetRows(); i++ {
		t, err := d.GetRowText(i)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t)
	}
}

func (d *MultiStringDisplay) add(s string) {
	d.Lines = append(d.Lines, s)
}
