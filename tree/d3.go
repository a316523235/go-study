package tree

/*
https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
*/
func connect(root *Node) *Node {
	connectXX(root)
	return root
}

func connectXX(root *Node) {
	nodeList := []*Node{root}
	i := 0
	cnt := 1
	for i < cnt {
		var pre *Node
		for ; i < cnt; i++ {
			currentNode := nodeList[i]
			if currentNode == nil {
				continue
			}
			if pre != nil {
				pre.Next = currentNode
			}
			nodeList = append(nodeList, currentNode.Left, currentNode.Right)
			pre = currentNode
		}
		cnt = len(nodeList)
	}
	return
}
