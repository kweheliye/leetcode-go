package easy

import "github.com/kweheliye/leetcode-go/linkedlist"

func getDecimalValue(head *linkedlist.ListNode) int {
	num := head.Val

	for head.Next != nil {
		num = num*2 + head.Next.Val
		head = head.Next
	}
	return num
}
