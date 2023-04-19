package main

import "fmt"

func main() {
	board := [3][3]int{
		{2, 1, 2},
		{2, 1, 1},
		{1, 1, 2},
	}

	fmt.Println(IsSolved(board))
}

func IsSolved(board [3][3]int) int {
	empty := 0

	if board[0][0] == board[1][1] &&
		board[1][1] == board[2][2] &&
		board[2][2] != 0 {
		return board[1][1]
	}

	if board[0][2] == board[1][1] &&
		board[1][1] == board[2][0] &&
		board[2][0] != 0 {
		return board[1][1]
	}

	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] &&
			board[i][1] == board[i][2] &&
			board[i][2] != 0 {
			return board[i][0]
		}

		if board[0][i] == board[1][i] &&
			board[1][i] == board[2][i] &&
			board[2][i] != 0 {
			return board[0][i]
		}

		if board[0][i] == 0 ||
			board[1][i] == 0 ||
			board[2][i] == 0 {
			empty = 1
		}
	}

	if empty == 1 {
		return -1
	}

	return 0
}
