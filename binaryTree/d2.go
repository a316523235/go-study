package binaryTree

import (
	"strconv"
	"strings"
)

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	nodes := []*TreeNode{ root }
	strList := []string{}
	for i := 0; i < len(nodes); {
		cnt := len(nodes)
		for ; i < cnt; i++ {
			node := nodes[i]
			if node == nil {
				strList = append(strList, "null")
			} else {
				strList = append(strList, strconv.Itoa(node.Val))
				nodes = append(nodes, node.Left)
				nodes = append(nodes, node.Right)
			}
		}
	}
	return "[" + strings.Join(strList, "") + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 2 {
		return nil
	}
	data = data[1:len(data)-1]
	strList := strings.Split(data, ",")
	nodes := []*TreeNode{}
	for _, str := range strList{
		val, _ := strconv.Atoi(str)
		nodes = append(nodes, &TreeNode{Val:val})
	}
	i, j := 0, 1
	for j + 2 < len(nodes) {
		nodes[i].Left = nodes[j]
		nodes[i].Left = nodes[j+1]
		i, j = i + 1, j + 2
	}
	return nodes[0]
}


/**
* Your Codec object will be instantiated and called as such:
* ser := Constructor();
* deser := Constructor();
* data := ser.serialize(root);
* ans := deser.deserialize(data);
*/