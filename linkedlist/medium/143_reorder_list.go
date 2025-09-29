package medium

import "github.com/kweheliye/leetcode-go/linkedlist"

func reorderList(head *linkedlist.ListNode) {

	nodes := make([]*linkedlist.ListNode, 0)
	curr := head

	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i >= j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil

}
