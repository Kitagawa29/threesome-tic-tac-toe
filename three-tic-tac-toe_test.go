package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

//oが正しく置けるか
func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("test1")
	}
}

//xが正しく置けるか
func TestPutToken02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(4, 4, "x")
	if b.get(4, 4) != "x" {
		t.Errorf("test2")
	}
}

//+が正しく置けるか
func TestPutToken03(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "+")
	if b.get(0, 0) != "+" {
		t.Errorf("test3")
	}
}

//盤が表示できるか
func TestShow01(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout := os.Stdout
	os.Stdout = w

	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	b.put(0, 0, "o")
	b.put(0, 4, "+")
	b.put(4, 4, "x")
	b.show()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if buf.String() != "o....\n.....\n.....\n.....\n+...x\n" {
		t.Errorf("test4")
	}
}

//横が3つ揃った時
func TestJudge01(t *testing.T) {
	b := &Board{
		tokens: []int{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	if b.judge() != "o" {
		t.Errorf("test4")
	}
}

//横がooxo+のような時(揃っていない)
func TestJudge02(t *testing.T) {
	b := &Board{
		tokens: []int{1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	if b.judge() != "undergo" {
		t.Errorf("test5")
	}
}

//縦が3つ揃った時
func TestJudge03(t *testing.T) {
	b := &Board{
		tokens: []int{2, 0, 0, 0, 0, 2, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	if b.judge() != "x" {
		t.Errorf("test6")
	}
}

//縦がxxox+のような時(揃っていない)
func TestJudge04(t *testing.T) {
	b := &Board{
		tokens: []int{2, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	if b.judge() != "undergo" {
		t.Errorf("test7")
	}
}

//斜め右下パターン１
func TestJudge05(t *testing.T) {
	b := &Board{
		tokens: []int{3, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 3, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	if b.judge() != "+" {
		t.Errorf("test8")
	}
}

//斜め右下パターン２
func TestJudge06(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0},
	}
	if b.judge() != "+" {
		t.Errorf("test9")
	}
}

//斜め右下パターン３
func TestJudge07(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 3},
	}
	if b.judge() != "+" {
		t.Errorf("test10")
	}
}

//斜め右下パターン４(失敗)
func TestJudge08(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 3, 0, 0},
	}
	if b.judge() != "undergo" {
		t.Errorf("test11")
	}
}

//斜め左下パターン３
func TestJudge09(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 3, 0, 0, 0, 3, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	if b.judge() != "+" {
		t.Errorf("test12")
	}
}

//斜め左下パターン４(失敗)
func TestJudge10(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 3, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 3, 0, 0},
	}
	if b.judge() != "undergo" {
		t.Errorf("test13")
	}
}
