package question_151_160

// 153. 寻找旋转排序数组中的最小值
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
// Topics: 数组 二分查找

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for r > l {
		c := (l + r) / 2
		if nums[c] > nums[r] {
			l = c + 1
		} else if nums[c] < nums[r] {
			r = c
		}
	}
	return nums[l]
}
