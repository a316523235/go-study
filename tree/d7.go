package tree

import "math"

func getMinimumDifference(root *TreeNode) int {
	pre, minDiff = -1, math.MaxInt32
	getMinimumDifferenceD1(root)
	return minDiff
}

var pre, minDiff int

func getMinimumDifferenceD1(root *TreeNode) {
	if root == nil {
		return
	}
	getMinimumDifferenceD1(root.Left)
	if root.Val-pre < minDiff {
		minDiff = root.Val - pre
	}
	pre = root.Val
	getMinimumDifferenceD1(root.Right)
}
