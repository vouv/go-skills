// +build OMIT

package main

import "fmt"

func exist(board [][]byte, word string) bool {
	queue := []byte(word)
	var result = false
	if len(queue) == 0 {
		return true
	}
	for i, line := range board {
		for j := range line {
			if board[i][j] == queue[0] {
				next(board, queue[1:], i, j, &result)
				if result {
					return true
				}
			}
		}
	}
	return result
}

func next(board [][]byte, queue []byte, i, j int, result *bool) {
	if len(queue) == 0 {
		*result = true
		return
	}
	tmp := board[i][j]
	board[i][j] = '-'
	c := queue[0]

	if i > 0 && board[i-1][j] == c {
		next(board, queue[1:], i-1, j, result)
		if *result {
			return
		}
	}
	if i+1 < len(board) && board[i+1][j] == c {
		next(board, queue[1:], i+1, j, result)
		if *result {
			return
		}
	}
	if j > 0 && board[i][j-1] == c {
		next(board, queue[1:], i, j-1, result)
		if *result {
			return
		}
	}
	if j+1 < len(board[0]) && board[i][j+1] == c {
		next(board, queue[1:], i, j+1, result)
		if *result {
			return
		}
	}
	board[i][j] = tmp
	return
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println("expect true and get", exist(board, "ABCCED"))
}
