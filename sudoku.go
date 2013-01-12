// Sudoku sover in Go
package main

import (
	"bytes"
	"fmt"
)

const BOARD_SIZE = 9 // rows and column size

// Build a slice of slices to represent the board from a string
func FromStr(boardRepr string) [][]int {
	board := make([][]int, BOARD_SIZE)
	for i := range board {
		board[i] = make([]int, BOARD_SIZE)
	}

	x, y := 0, 0
	for i := range boardRepr {
		if boardRepr[i] >= '0' && boardRepr[i] <= '9' {
			board[x][y] = int(boardRepr[i]) - 48
			if y == BOARD_SIZE-1 {
				x += 1 % BOARD_SIZE
				y = 0
			} else {
				y += 1 % BOARD_SIZE
			}
		}
	}

	return board
}

// Board to string
func ToStr(board [][]int) string {
	var b bytes.Buffer
	for i := range board {
		for j := range board[i] {
			b.WriteString(fmt.Sprintf("%d ", board[i][j]))
		}
		b.WriteString("\n")
	}
	return b.String()
}

// Return the next spot where is possible to put a number
func nextEmpty(board [][]int) []int {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Return true if the number n can be put in the board at specific spot
func canPut(board [][]int, spot []int, n int) bool {
	x, y := spot[0], spot[1]
	for i := range board {
		// test line
		if board[i][y] == n {
			return false
		}
		// test column
		if board[x][i] == n {
			return false
		}
	}
	// test square
	a, b := x-(x%3), y-(y%3)
	for i := a; i < a+3; i++ {
		for j := b; j < b+3; j++ {
			if board[i][j] == n {
				return false
			}
		}
	}

	return true
}

// Solve using backtracking
func solve(board [][]int) [][]int {
	spot := nextEmpty(board)
	if spot == nil {
		return board
	}
	x, y := spot[0], spot[1]
	for i := 1; i < 10; i++ {
		if canPut(board, spot, i) {
			board[x][y] = i
			newBoard := solve(board)
			if nextEmpty(newBoard) == nil {
				// solution found
				return newBoard
			}
		}
	}
	// solution not found, backtrack
	board[x][y] = 0
	return board
}

func main() {
	board := FromStr(
		`0 0 0 1 0 0 7 0 2
		0 3 0 9 5 0 0 0 0
		0 0 1 0 0 2 0 0 3
		5 9 0 0 0 0 3 0 1
		0 2 0 0 0 0 0 7 0
		7 0 3 0 0 0 0 9 8
		8 0 0 2 0 0 1 0 0
		0 0 0 0 8 5 0 6 0
		6 0 5 0 0 9 0 0 0`)
	board2 := FromStr(`5 3 4 6 7 8 9 1 2
        6 7 2 1 9 5 3 4 8
        1 9 8 3 4 2 5 6 7
        8 5 9 7 6 1 4 2 3
        4 2 6 8 5 3 7 9 1
        7 1 3 9 2 4 8 5 6
        9 6 1 5 3 7 2 8 4
        2 8 7 4 1 9 6 3 5
        3 4 5 2 8 6 1 7 9`)
	board3 := FromStr(`0 0 0 1 0 0 7 0 2
0 3 0 9 5 0 0 0 0
0 0 1 0 0 2 0 0 3
5 9 0 0 0 0 3 0 1
0 2 0 0 3 0 0 7 0
7 0 3 0 0 0 0 9 8
8 0 0 2 0 0 1 0 0
0 0 0 0 8 5 0 6 0
6 0 5 0 0 9 0 0 0`)
	fmt.Println(ToStr(solve(board)))
	fmt.Println(ToStr(solve(board2)))
	fmt.Println(ToStr(solve(board3)))
}
