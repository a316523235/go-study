package tree

/*
https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
*/
func postorderTraversal(root *TreeNode) []int {
	res = []int{}
	postorderTraversalMethod1(root)
	return res
}

var res []int

func postorderTraversalMethod1(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversalMethod1(root.Left)
	postorderTraversalMethod1(root.Right)
	res = append(res, root.Val)
}
