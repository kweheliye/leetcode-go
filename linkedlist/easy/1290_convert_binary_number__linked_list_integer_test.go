package easy

import (
	"github.com/kweheliye/leetcode-go/linkedlist"
	"testing"
)

func TestGetDecimalValue(t *testing.T) {
	type args struct {
		head *linkedlist.ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				head: &linkedlist.ListNode{
					Val: 1,
					Next: &linkedlist.ListNode{
						Val: 0,
						Next: &linkedlist.ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
