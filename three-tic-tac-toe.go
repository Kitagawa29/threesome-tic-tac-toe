package main

import "fmt"

import "strconv"

//import "regexp"

var p [3]string
var n [3]string

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

// func (b *Board) put(x, y int, u string) // interface
// func (b *Board) get(x, y int) string    // interface

func (b *Board) put(x, y int, u string) {
	if u == p[0] {
		b.tokens[x+5*y] = 1
	} else if u == p[1] {
		b.tokens[x+5*y] = 2
	} else if u == p[2] {
		b.tokens[x+5*y] = 3
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+5*y] == 1 {
		return p[0]
	} else if b.tokens[x+5*y] == 2 {
		return p[1]
	} else if b.tokens[x+5*y] == 3 {
		return p[2]
	} else {
		return "."
	}
}

func (b *Board) show() {
	for i := 0; i < 25; i++ {
		x := i % 5
		y := i / 5
		fmt.Print(b.get(x, y))
		if (i+1)%5 == 0 {
			fmt.Println()
		}
	}
}

//勝敗チェック
func (b *Board) judge() string {
	//横の勝敗
	for row := 0; row < 5; row++ {
		for column := 0; column < (5 - 2); column++ {
			if b.tokens[column+5*row] == b.tokens[(column+1)+5*row] && b.tokens[column+5*row] == b.tokens[(column+2)+5*row] && b.tokens[column+5*row] != 0 {
				return b.get(column, row)
			}
		}
	}
	//縦の勝敗
	for column := 0; column < 5; column++ {
		for row := 0; row < (5 - 2); row++ {
			if b.tokens[column+5*row] == b.tokens[column+5*(row+1)] && b.tokens[column+5*row] == b.tokens[column+5*(row+2)] && b.tokens[column+5*row] != 0 {
				return b.get(column, row)
			}
		}
	}
	//斜め（右下）
	for column := 0; column < (5 - 2); column++ {
		for row := 0; row < (5 - 2); row++ {
			if b.tokens[column+5*row] == b.tokens[(column+1)+5*(row+1)] && b.tokens[column+5*row] == b.tokens[(column+2)+5*(row+2)] && b.tokens[column+5*row] != 0 {
				return b.get(column, row)
			}
		}
	}
	//斜め（左下）
	for column := (0 + 2); column < 5; column++ {
		for row := 0; row < (5 - 2); row++ {
			if b.tokens[column+5*row] == b.tokens[(column-1)+5*(row+1)] && b.tokens[column+5*row] == b.tokens[(column-2)+5*(row+2)] && b.tokens[column+5*row] != 0 {
				return b.get(column, row)
			}
		}
	}
	return "undergo"
}

func (b *Board) play() {

	var x, y int
	var player, s string

	fmt.Println("Please Input Your Names")
	for i := 0; i < 3; i++ {
		fmt.Print("Player ", i+1, " : ")
		fmt.Scanf("%s", &n[i])
	}

	fmt.Println("Please Input Your Pieces")
	for i := 0; i < 3; i++ {
		fmt.Print(n[i], ": ")
		fmt.Scanf("%s", &p[i])
	}
	m := map[string]string{p[0]: n[0], p[1]: n[1], p[2]: n[2]}

	for i := 0; i < 25; i++ {
		if (i+1)%3 == 1 {
			fmt.Print(n[0], ": Input x,y ")
			player = p[0]
		} else if (i+1)%3 == 2 {
			fmt.Print(n[1], ": Input x,y ")
			player = p[1]
		} else {
			fmt.Print(n[2], ": Input x,y ")
			player = p[2]
		}

		// fmt.Scanf("%d,%d", &y, &x)
		fmt.Scanf("%s", &s)

		//形式が正しい時にのみ処理を行う
		if (len(s) == 3) && (string(s[1]) == ",") {
			y, _ = strconv.Atoi(string(s[0]))
			x, _ = strconv.Atoi(string(s[2]))
			x = x - 1
			y = y - 1
		} else {
			i -= 1
			fmt.Println("Error! Input form is wrong!")
			continue
		}
		//盤の外に置こうとした時・重複した場所に置こうとした時のエラー処理
		if ((0 <= x) && (x <= 4) && (0 <= y) && (y <= 4)) && (b.get(x, y) == ".") {
			b.put(x, y, player)
		} else {
			i -= 1
			fmt.Println("Error! Input again!")
			continue
		}

		b.show()

		if b.judge() == player {
			fmt.Println("Player " + m[player] + " won")
			return
		}
	}
	fmt.Println("Draw")
	return
}

func main() {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.play()
}
