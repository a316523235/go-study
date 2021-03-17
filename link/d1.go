package link

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
var head = BuildList([]int{4,5,1, 9})

func BuildList(arr []int) *ListNode {
	next := &ListNode{Val:0}
	next = nil
	for i := len(arr) - 1; i > -1; i--  {
		node := &ListNode{
			Val:arr[i],
			Next:next,
		}
		next = node
	}
	return next
}


func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cnt := 0
	for temp := head; temp != nil; temp = temp.Next {
		cnt++
	}
	if cnt == n {
		return head.Next
	}
	temp := head
	for i := cnt; i > n + 1; i-- {
		temp = temp.Next
	}

	if temp.Next.Next == nil {
		temp.Next = nil
	} else {
		temp.Next = temp.Next.Next
	}
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil  {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseListByRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListByRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{-1, nil}
	cur := res
	for l1 != nil || l2 != nil {
		var next *ListNode
		switch  {
		case l1 == nil:
			cur.Next = l2
			goto over
		case l2 == nil:
			cur.Next = l1
			goto over
		case l1.Val < l2.Val:
			next = l1
			l1 = l1.Next
		default:
			next = l2
			l2 = l2.Next
		}
		cur.Next = next
		cur = cur.Next
	}
	over:
	return res.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	var lst []*ListNode
	for temp := head; temp != nil; temp = temp.Next {
		lst = append(lst, temp)
	}

	for i, j := 0, len(lst); i < j; i, j = i+1, j-1 {
		if lst[i].Val != lst[j].Val {
			return false
		}
	}
	return true
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return false
}





