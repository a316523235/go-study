package tree

/*
https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if q.Val < p.Val {
		p, q = q, p
	}
	return lowestCommonAncestorXx(root, p, q)
}

func lowestCommonAncestorXx(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	if root.Val > p.Val && root.Val < q.Val {
		return root
	}
	if q.Val < root.Val {
		return lowestCommonAncestorXx(root.Left, p, q)
	} else {
		return lowestCommonAncestorXx(root.Right, p, q)
	}
}
