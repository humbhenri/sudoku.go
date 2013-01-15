// Sudoku sover in Go
package main

import (
	"bytes"
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
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


// genericSplit provides a generic version of Split and SplitAfter.
// Set the includeSep bool to true to have it include the separtor.
func genericSplit(re *regexp.Regexp,s string, numFields int, includeSep bool) []string {
	if numFields == 0 {
		return make([]string, 0)
	}

	// Using regexp, including the separator is really easy. Instead of
	// including up to the start of the separator we include to the end.
	// The start of the separator is stored in index 0.
	// The end of the separator is stored in index 1.
	var includeTo int
	if includeSep {
		includeTo = 1
	} else {
		includeTo = 0
	}

	count := re.FindAllStringIndex(s, numFields-1)
	n := len(count) + 1
	stor := make([]string, n)

	if n == 1 {
		stor[0] = s
		return stor
	}

	stor[0] = s[:count[0][includeTo]]

	for i := 1; i < n-1; i++ {
		stor[i] = s[count[i-1][1]:count[i][includeTo]]
	}

	stor[n-1] = s[count[n-2][1]:]

	return stor
}


func processBatch(file string) {
	data, err := ioutil.ReadFile(file)
	if (err != nil) {
		panic(err)
	}
	count := 0
	before := time.Now()
	for _, line := range strings.Split(string(data), "\n") {
		sudoku := FromStr(line)
		solve(sudoku)
		count++
	}
	diff := time.Now().Sub(before)
	fmt.Printf("-- Solved %d sudokus. Elapsed time: %f seconds\n",count, diff.Seconds())
}


func main() {
	file := os.Args[1]
	processBatch(file)
}
