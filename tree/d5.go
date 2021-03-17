package tree

/*
https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	insertIntoBSTXx(root, val)
	return root
}

func insertIntoBSTXx(root *TreeNode, val int) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
			return
		}
		insertIntoBSTXx(root.Left, val)
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
			return
		}
		insertIntoBSTXx(root.Right, val)
	}
}
