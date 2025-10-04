package easy

import (
	"reflect"
	"testing"

	"github.com/kweheliye/leetcode-go/linkedlist"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *linkedlist.ListNode
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.ListNode
	}{
		{
			name: "Test Case 1",
			args: args{
				head: &linkedlist.ListNode{
					Val: 1,
					Next: &linkedlist.ListNode{
						Val: 1,
						Next: &linkedlist.ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
			want: &linkedlist.ListNode{
				Val: 1,
				Next: &linkedlist.ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
