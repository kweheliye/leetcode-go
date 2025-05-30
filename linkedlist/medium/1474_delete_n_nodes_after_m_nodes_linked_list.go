package medium

import "github.com/kweheliye/leetcode-go/linkedlist"

func deleteNodes(head *linkedlist.ListNode, m, n int) *linkedlist.ListNode {
	currentNode := head
	listMNode := head

	for currentNode != nil {
		mCount, nCount := m, n

		for currentNode != nil && mCount != 0 {
			listMNode = currentNode
			currentNode = currentNode.Next
			mCount--
		}

		for currentNode != nil && nCount != 0 {
			currentNode = currentNode.Next
			nCount--
		}
		listMNode.Next = currentNode

	}
	return head
}
