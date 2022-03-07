package go0203

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	resultNode := removeElements(head.Next, val)
	if head.Val == val {
		return resultNode
	} else {
		head.Next = resultNode
		return head
	}
}
