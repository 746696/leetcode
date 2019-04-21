package question_31_40

import (
	"reflect"
	"testing"
)

func Test_solveSudoku(t *testing.T) {
	board1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '.', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	want1 := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '2', '6', '7', '5', '1', '4', '9', '3'},
		{'4', '5', '9', '8', '6', '3', '7', '2', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}

	board2 := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}

	want2 := [][]byte{
		{'5', '1', '9', '7', '4', '8', '6', '3', '2'},
		{'7', '8', '3', '6', '5', '2', '4', '1', '9'},
		{'4', '2', '6', '1', '3', '9', '8', '7', '5'},
		{'3', '5', '7', '9', '8', '6', '2', '4', '1'},
		{'2', '6', '4', '3', '1', '7', '5', '9', '8'},
		{'1', '9', '8', '5', '2', '4', '3', '6', '7'},
		{'9', '7', '5', '8', '6', '3', '1', '2', '4'},
		{'8', '3', '2', '4', '9', '1', '7', '5', '6'},
		{'6', '4', '1', '2', '7', '5', '9', '8', '3'},
	}

	tests := []struct {
		name  string
		board [][]byte
		want  [][]byte
	}{
		{"test#1", board1, want1},
		{"test#2", board2, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.board)
			if !reflect.DeepEqual(tt.board, tt.want) {
				t.Errorf("error")
			}
		})
	}
}