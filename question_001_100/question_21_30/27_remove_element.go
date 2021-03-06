package question_0011_0020

// 27. 移除元素
// https://leetcode-cn.com/problems/remove-element
// Topics: 数组 双指针

func removeElement(nums []int, val int) int {
	slow := 0
	for _, n := range nums {
		if n != val {
			nums[slow], slow = n, slow+1
		}
	}
	return slow
}
