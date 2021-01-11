package question_1311_1320

import (
	"reflect"
	"testing"
)

func Test_matrixBlockSum(t *testing.T) {
	tests := []struct {
		mat  [][]int
		K    int
		want [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1, [][]int{{}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2, [][]int{{}}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, 1, [][]int{{}}},
		{
			[][]int{
				{76, 4, 73},
				{21, 8, 56},
				{4, 56, 61},
				{70, 32, 38},
				{31, 94, 67}},
			1,
			[][]int{
				{109, 238, 141},
				{169, 359, 258},
				{191, 346, 251},
				{287, 453, 348},
				{227, 332, 231},
			},
		},
		{
			[][]int{
				{80, 82, 40, 51, 53, 69, 8, 84},
				{99, 51, 27, 72, 59, 91, 99, 9},
				{27, 81, 95, 50, 15, 33, 67, 80},
				{62, 9, 90, 89, 89, 69, 36, 12},
				{6, 59, 76, 33, 21, 83, 65, 5},
				{94, 86, 35, 1, 98, 25, 93, 41},
				{43, 9, 39, 5, 54, 32, 10, 30},
				{99, 57, 75, 75, 23, 14, 83, 75},
				{16, 56, 97, 82, 19, 22, 18, 34},
				{89, 78, 13, 9, 51, 20, 7, 36}},
			2,
			[][]int{
				{582, 755, 882, 869, 829, 840, 667, 540},
				{743, 1005, 1221, 1215, 1202, 1135, 873, 657},
				{884, 1179, 1416, 1487, 1480, 1342, 1047, 810},
				{897, 1142, 1424, 1437, 1511, 1335, 1090, 808},
				{811, 989, 1266, 1276, 1303, 1136, 958, 681},
				{839, 1042, 1327, 1246, 1313, 1161, 958, 673},
				{847, 1043, 1258, 1176, 1178, 1041, 845, 630},
				{886, 1058, 1303, 1075, 1000, 957, 785, 540},
				{671, 842, 989, 830, 748, 699, 528, 381},
				{580, 746, 839, 691, 608, 568, 402, 309},
			},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := matrixBlockSum(tt.mat, tt.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixBlockSum() = %v, want %v", got, tt.want)
			}
		})
	}
}