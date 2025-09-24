package easy

import "github.com/kweheliye/leetcode-go/linkedlist"

func addTwoNumbers(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {

	dummyHead := &linkedlist.ListNode{}
	curr := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		x := 0
		if l1 != nil {
			x = l1.Val
		}

		y := 0
		if l2 != nil {
			y = l2.Val
		}

		sum := carry + x + y
		carry = sum / 10
		curr.Next = &linkedlist.ListNode{Val: sum % 10}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return dummyHead.Next

}
