package question_141_150

// 142. 环形链表 II
// https://leetcode-cn.com/problems/linked-list-cycle-ii/
// Topics: 链表 双指针

import (
	. "github.com/lupes/leetcode/common"
)

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for next := head; next != nil; next = next.Next {
		if _, ok := m[next]; ok {
			return next
		} else {
			m[next] = struct{}{}
		}
	}
	return nil
}
