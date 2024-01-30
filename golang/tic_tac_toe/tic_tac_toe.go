/*
Tic-tac-toe is played by two players A and B on a 3 x 3 grid. The rules of Tic-Tac-Toe are:

Players take turns placing characters into empty squares ' '.
The first player A always places 'X' characters, while the second player B always places 'O' characters.
'X' and 'O' characters are always placed into empty squares, never on filled ones.
The game ends when there are three of the same (non-empty) character filling any row, column, or diagonal.
The game also ends if all squares are non-empty.
No more moves can be played if the game is over.
Given a 2D integer array moves where moves[i] = [rowi, coli] indicates that the ith move will be played on grid[rowi][coli]. return the winner of the game if it exists (A or B). In case the game ends in a draw return "Draw". If there are still movements to play return "Pending".

You can assume that moves is valid (i.e., it follows the rules of Tic-Tac-Toe), the grid is initially empty, and A will play first.
*/

package main

import (
	"fmt"
	"reflect"
	"strings"
)

func tictactoe(moves [][]int) string {
	// Problem
	// moves is 2d array where first element is A and then B
	// Need to populate the board
	// Need to determine the winner
	// There can be a draw

	// Solution
	// We'll first want to populate the game board (fillBoard)
	board := fillBoard(moves)

	winner := checkWinner(board)
	// We'll then want to search the game board for a winner (checkWinner)
	return winner
}

func checkWinner(board [][]string) string {
	// - We can manually loop over each row, join, and check if it equals "XXX"
	//		- Also check for "OOO"
	// - On each loop we can also add the column to a variable for column 1, 2, 3
	// - On each loop we can also add the diagonal to a variable diagTopLeft and diagTopRight
	var (
		col1 string 
		col2 string 
		col3 string
	) 

	winner := "Draw"
	var diagTopLeft string
	var diagTopRight string

	for i, row := range board {
		col1 += row[0]
		col2 += row[1]
		col3 += row[2]

		switch i {
		case 0:
			diagTopLeft += row[0]
			diagTopRight += row[2]
		case 1:
			diagTopLeft += row[1]
			diagTopRight += row[1]
		case 2:
			diagTopLeft += row[2]
			diagTopRight += row[0]
		}

		joinedRow := strings.Join(row, "")
		if joinedRow == "XXX" {
			winner = "A"
		} else if joinedRow == "OOO" {
			winner = "B"
		}
	}

	if col1 == "XXX" || col2 == "XXX" || col3 == "XXX" {
		winner = "A"
	} else if col1 == "OOO" || col2 == "OOO" || col3 == "OOO" {
		winner = "B"
	}

	if diagTopLeft == "XXX" || diagTopRight == "XXX" {
		winner = "A"
	} else if diagTopLeft == "OOO" || diagTopRight == "OOO" {
		winner = "B"
	}

	return winner
}

func fillBoard(moves [][]int) [][]string {
	// - initialize an empty 3x3 board
	// - Loop over each move where first element is A (even) and then next is B (odd)
	// - A moves are i % 2 == 0
	// - Each move is row, col
	// - On each loop we'll board[row][col]
	// - return a board
	board := [][]string{
		{"","",""},
		{"","",""},
		{"","",""},
	}

	for i, move := range moves {
		row, col := move[0], move[1]
		// A 
		if i % 2 == 0 {
			board[row][col] = "X"
		} else {
			board[row][col] = "O"
		}
	}

	return board
}

func main() {
	tests := []struct{
		input [][]int
		want string
	}{
		{ input: [][]int{{0,0},{2,0},{1,1},{2,1},{2,2}}, want: "A" },
		{ input: [][]int{{0,0},{1,1},{0,1},{0,2},{1,0},{2,0}}, want: "B" },
		{ input: [][]int{{0,0},{1,1},{2,0},{1,0},{1,2},{2,1},{0,1},{0,2},{2,2}}, want: "Draw" },
	}

	for _, tc := range tests {
		got := tictactoe(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("FAILED! got: %s, want: %s\n", got, tc.want)
		} else {
			fmt.Printf("PASSED! got: %s, want: %s\n", got, tc.want)
		}
	}
}

