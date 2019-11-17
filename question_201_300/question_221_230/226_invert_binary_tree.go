package question_221_230

// 226. 翻转二叉树
// https://leetcode-cn.com/problems/invert-binary-tree/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}