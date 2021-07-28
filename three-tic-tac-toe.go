package main

import "fmt"

//import "regexp"

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

// func (b *Board) put(x, y int, u string) // interface
// func (b *Board) get(x, y int) string    // interface

func (b *Board) put(x, y int, u string) {
	if u == "o" {
		b.tokens[x+5*y] = 1
	} else if u == "x" {
		b.tokens[x+5*y] = 2
	} else if u == "+" {
		b.tokens[x+5*y] = 3
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+5*y] == 1 {
		return "o"
	} else if b.tokens[x+5*y] == 2 {
		return "x"
	} else if b.tokens[x+5*y] == 3 {
		return "+"
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
	var player string
	m := map[string]string{"o": "1", "x": "2", "+": "3"}

	for i := 0; i < 25; i++ {
		if (i+1)%3 == 1 {
			fmt.Print("Player 1: Input (x,y) ")
			player = "o"
		} else if (i+1)%3 == 2 {
			fmt.Print("Player 2: Input (x,y) ")
			player = "x"
		} else {
			fmt.Print("Player 3: Input (x,y) ")
			player = "+"
		}

		fmt.Scanf("%d,%d", &y, &x)
		// fmt.Scanf("%s", &s)

		x = x - 1
		y = y - 1

		//盤の外に置こうとした時・重複した場所に置いた時のエラー処理
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
