package main

import (
	"errors"
	"fmt"
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
