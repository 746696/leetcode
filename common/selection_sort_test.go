package common

import (
	"reflect"
	"testing"
)

func Test_selectionSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{2, 1, 7, 8, 2, 4, 3, 0, 9, 5, 6}, []int{0, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := selectionSort(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
