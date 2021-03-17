package tree

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	low, fast := head, head.Next
	for low != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		low = low.Next
		fast = fast.Next.Next
	}
	return true
}
