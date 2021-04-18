package question_1631_1640

// 1637. 两点之间不包含任何点的最宽垂直面积
// https://leetcode-cn.com/problems/widest-vertical-area-between-two-points-containing-no-points/
// Topics: 排序

import (
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	var max = 0
	for i := 1; i < len(points); i++ {
		if points[i][0]-points[i-1][0] > max {
			max = points[i][0] - points[i-1][0]
		}
	}
	return max
}
