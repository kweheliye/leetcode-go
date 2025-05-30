package easy

import "github.com/kweheliye/leetcode-go/linkedlist"

func mergeTwoListsV1(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoListsV1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsV1(l1, l2.Next)
		return l2
	}
}

func mergeTwoListsV2(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	prehead := &linkedlist.ListNode{-1, nil}
	prev := prehead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}
	return prehead.Next
}
