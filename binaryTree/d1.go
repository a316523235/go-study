package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var idxMap map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {
	idxMap = map[int]int{}
	for i, v := range inorder {
		idxMap[v] = i
	}
	var build func(int, int) *TreeNode

	build = func (inorderLeft, inorderRight int) *TreeNode {
		// 无剩余节点
		if inorderLeft > inorderRight {
			return nil
		}

		val := postorder[len(postorder) - 1]
		postorder = postorder[:len(postorder) - 1]
		node := &TreeNode{ Val:val }

		index := idxMap[val]
		node.Right = build(index + 1, inorderRight)
		node.Left = build(inorderLeft, index - 1)
		return node
	}

	xx := build(0, len(inorder) - 1)
	return xx
}