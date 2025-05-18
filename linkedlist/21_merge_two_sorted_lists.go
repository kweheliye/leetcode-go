package linkedlist

func mergeTwoListsV1(l1 *ListNode, l2 *ListNode) *ListNode {

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

func mergeTwoListsV2(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{-1, nil}
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
