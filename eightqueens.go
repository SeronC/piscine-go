package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	var solutions [][]int

	var isSafe func(int, int) bool
	isSafe = func(row, col int) bool {
		for i := 0; i < col; i++ {
			if board[i] == row || board[i]-row == i-col || board[i]-row == col-i {
				return false
			}
		}
		return true
	}

	var solve func(int)
	solve = func(col int) {
		if col == 8 {
			copyOfBoard := make([]int, 8)
			copy(copyOfBoard, board[:])
			solutions = append(solutions, copyOfBoard)
			return
		}

		for row := 0; row < 8; row++ {
			if isSafe(row, col) {
				board[col] = row
				solve(col + 1)
			}
		}
	}

	solve(0)

	for _, solution := range solutions {
		for _, row := range solution {
			for col := 0; col < 8; col++ {
				if col == row {
					z01.PrintRune('Q')
				} else {
					z01.PrintRune('.')
				}
			}
			z01.PrintRune(' ')
		}
		z01.PrintRune(' ')
	}
}
