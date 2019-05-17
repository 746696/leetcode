package question_71_80

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"test#1", []int{1}, [][]int{nil, {1}}},
		{"test#2", []int{1, 2}, [][]int{nil, {1, 2}, {1}, {2}}},
		{"test#3", []int{1, 2, 3}, [][]int{nil, {1, 2, 3}, {1}, {2, 3}, {2}, {1, 3}, {3}, {1, 2}}},
		{"test#4", []int{1, 2, 3, 4}, [][]int{nil, {1, 2, 3, 4}, {1}, {2, 3, 4}, {2}, {1, 3, 4}, {3}, {1, 2, 4}, {4}, {1, 2, 3}, {1, 2}, {3, 4}, {1, 3}, {2, 4}, {1, 4}, {2, 3}}},
		{"test#5", []int{1, 2, 3, 4, 5}, [][]int{nil, {1, 2, 3, 4, 5}, {1}, {2, 3, 4, 5}, {2}, {1, 3, 4, 5}, {3}, {1, 2, 4, 5}, {4}, {1, 2, 3, 5}, {5}, {1, 2, 3, 4}, {1, 2}, {3, 4, 5}, {1, 3}, {2, 4, 5}, {1, 4}, {2, 3, 5}, {1, 5}, {2, 3, 4}, {2, 3}, {1, 4, 5}, {2, 4}, {1, 3, 5}, {2, 5}, {1, 3, 4}, {3, 4}, {1, 2, 5}, {3, 5}, {1, 2, 4}, {4, 5}, {1, 2, 3}}},
		{"test#6", []int{1, 2, 3, 4, 5, 6}, [][]int{nil, {1, 2, 3, 4, 5, 6}, {1}, {2, 3, 4, 5, 6}, {2}, {1, 3, 4, 5, 6}, {3}, {1, 2, 4, 5, 6}, {4}, {1, 2, 3, 5, 6}, {5}, {1, 2, 3, 4, 6}, {6}, {1, 2, 3, 4, 5}, {1, 2}, {3, 4, 5, 6}, {1, 3}, {2, 4, 5, 6}, {1, 4}, {2, 3, 5, 6}, {1, 5}, {2, 3, 4, 6}, {1, 6}, {2, 3, 4, 5}, {2, 3}, {1, 4, 5, 6}, {2, 4}, {1, 3, 5, 6}, {2, 5}, {1, 3, 4, 6}, {2, 6}, {1, 3, 4, 5}, {3, 4}, {1, 2, 5, 6}, {3, 5}, {1, 2, 4, 6}, {3, 6}, {1, 2, 4, 5}, {4, 5}, {1, 2, 3, 6}, {4, 6}, {1, 2, 3, 5}, {5, 6}, {1, 2, 3, 4}, {1, 2, 3}, {4, 5, 6}, {1, 2, 4}, {3, 5, 6}, {1, 2, 5}, {3, 4, 6}, {1, 2, 6}, {3, 4, 5}, {1, 3, 4}, {2, 5, 6}, {1, 3, 5}, {2, 4, 6}, {1, 3, 6}, {2, 4, 5}, {1, 4, 5}, {2, 3, 6}, {1, 4, 6}, {2, 3, 5}, {1, 5, 6}, {2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
