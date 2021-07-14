package main

import "fmt"

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

// 	if b.tokens[0]+b.tokens[3]+b.tokens[6] == 3 {
// 		return "o"
// 	} else if b.tokens[1]+b.tokens[4]+b.tokens[7] == 3 {
// 		return "o"
// 	} else if b.tokens[2]+b.tokens[5]+b.tokens[8] == 3 {
// 		return "o"
// 	} else if b.tokens[0]+b.tokens[1]+b.tokens[2] == 3 {
// 		return "o"
// 	} else if b.tokens[3]+b.tokens[4]+b.tokens[5] == 3 {
// 		return "o"
// 	} else if b.tokens[6]+b.tokens[7]+b.tokens[8] == 3 {
// 		return "o"
// 	} else if b.tokens[0]+b.tokens[4]+b.tokens[8] == 3 {
// 		return "o"
// 	} else if b.tokens[2]+b.tokens[4]+b.tokens[6] == 3 {
// 		return "o"
// 	} else if b.tokens[0]+b.tokens[3]+b.tokens[6] == -3 {
// 		return "x"
// 	} else if b.tokens[1]+b.tokens[4]+b.tokens[7] == -3 {
// 		return "x"
// 	} else if b.tokens[2]+b.tokens[5]+b.tokens[8] == -3 {
// 		return "x"
// 	} else if b.tokens[0]+b.tokens[1]+b.tokens[2] == -3 {
// 		return "x"
// 	} else if b.tokens[3]+b.tokens[4]+b.tokens[5] == -3 {
// 		return "x"
// 	} else if b.tokens[6]+b.tokens[7]+b.tokens[8] == -3 {
// 		return "x"
// 	} else if b.tokens[0]+b.tokens[4]+b.tokens[8] == -3 {
// 		return "x"
// 	} else if b.tokens[2]+b.tokens[4]+b.tokens[6] == -3 {
// 		return "x"
// 	} else {
// 		return "undergo"
// 	}
// }

// func (b *Board) play() {

// 	var x, y int
// 	var player string
// 	m := map[string]string{"o": "1", "x": "2"}

// 	for i := 0; i < 9; i++ {
// 		if (i+1)%2 == 1 {
// 			fmt.Print("Player 1: Input (x,y) ")
// 			player = "o"
// 		} else {
// 			fmt.Print("Player 2: Input (x,y) ")
// 			player = "x"
// 		}

// 		fmt.Scanf("%d,%d", &y, &x)

// 		if ((0 <= x) && (x <= 2) && (0 <= y) && (y <= 2)) && (b.get(x, y) == ".") {
// 			b.put(x, y, player)
// 		} else {
// 			i -= 1
// 			fmt.Println("Error! Input again!")
// 			continue
// 		}
// 		b.show()

// 		if b.judge() == player {
// 			fmt.Println("Player " + m[player] + " won")
// 			return
// 		}
// 	}
// 	fmt.Println("Draw")
// 	return

func main() {
	// b := &Board{
	// 	tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// }
	// b.play()
}
