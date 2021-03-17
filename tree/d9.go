package tree

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a ,b int) int {
	if a > b {
		return a
	}
	return b
}

var pre *TreeNode

func isValidBST(root *TreeNode) bool {


	if root == nil {
		return true
	}
	//left
	if !isValidBST(root.Left) {
		return false
	}

	//middle
	if root.Val <= pre.Val {
		return false
	}
	pre = root


	//right
	if !isValidBST(root.Right) {
		return false
	}

	return true
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return isSymmetric2(root.Left, root.Right)
}

func isSymmetric2(rootA, rootB *TreeNode) bool {
	if rootA == nil && rootB == nil {
		return true
	}
	if rootA == nil || rootB == nil || rootA.Val != rootB.Val {
		return false
	}
	return isSymmetric2(rootA.Left, rootB.Right) && isSymmetric2(rootA.Right, rootB.Left)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	q := []*TreeNode{root}
	res := [][]int{}
	for len(q) > 0  {
		nextQ := []*TreeNode{}
		arr := []int{}
		for _, node := range q {
			if node == nil {
				continue
			}
			arr = append(arr, node.Val)
			if node.Left != nil {
				nextQ = append(nextQ, node.Left)
			}
			if node.Right != nil {
				nextQ = append(nextQ, node.Right)
			}
		}
		res = append(res, arr)
		q = nextQ
	}
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return buildMid(nums, 0, len(nums) - 1)
}

func buildMid(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	tree := &TreeNode{ Val:nums[mid] }
	tree.Left = buildMid(nums, l, mid - 1)
	tree.Right = buildMid(nums, mid + 1, r)
	return tree
}
