package easy

import "github.com/kweheliye/leetcode-go/linkedlist"

func deleteDuplicates(head *linkedlist.ListNode) *linkedlist.ListNode {

	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
