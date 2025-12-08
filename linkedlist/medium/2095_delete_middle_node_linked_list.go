package medium

import "github.com/kweheliye/leetcode-go/linkedlist"

func deleteMiddle(head *linkedlist.ListNode) *linkedlist.ListNode {
	// If the list is empty or has only one node, deleting the middle leaves an empty list
	if head == nil || head.Next == nil {
		return nil
	}

	if head.Next.Next == nil {
		return head.Next
	}

	slow, fast := head, head

	prev := (*linkedlist.ListNode)(nil)

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next

	return head

}
