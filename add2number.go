package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	carry := 0
	for currentNode := node; l1 != nil || l2 != nil || carry != 0; currentNode = currentNode.Next {
		if l1 == nil {
			l1 = &ListNode{0, nil}
		}
		if l2 == nil {
			l2 = &ListNode{0, nil}
		}
		sum := l1.Val + l2.Val + carry

		l1 = l1.Next
		l2 = l2.Next

		currentNode.Val = sum % 10
		carry = sum / 10
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		currentNode.Next = &ListNode{0, nil}
	}
	return node
}
