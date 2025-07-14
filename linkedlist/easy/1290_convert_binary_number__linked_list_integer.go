package easy

import "github.com/kweheliye/leetcode-go/linkedlist"

func getDecimalValue(head *linkedlist.ListNode) int {
	num := 0

	for head != nil {
		num = num*2 + head.Val
		head = head.Next
	}
	return num
}
