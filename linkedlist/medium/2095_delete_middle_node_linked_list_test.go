package medium

import (
	"reflect"
	"testing"

	"github.com/kweheliye/leetcode-go/linkedlist"
)

func Test_deleteMiddle(t *testing.T) {
	type args struct {
		head *linkedlist.ListNode
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.ListNode
	}{
		{
			name: "Single node - should become empty",
			args: args{
				head: &linkedlist.ListNode{Val: 1},
			},
			want: nil,
		},
		{
			name: "Two nodes - delete first (middle)",
			args: args{
				head: &linkedlist.ListNode{
					Val:  1,
					Next: &linkedlist.ListNode{Val: 2},
				},
			},
			want: &linkedlist.ListNode{Val: 2}, // Only node 2 remains
		},
		{
			name: "Three nodes - delete middle (2)",
			args: args{
				head: &linkedlist.ListNode{
					Val: 1,
					Next: &linkedlist.ListNode{
						Val:  2,
						Next: &linkedlist.ListNode{Val: 3},
					},
				},
			},
			want: &linkedlist.ListNode{
				Val:  1,
				Next: &linkedlist.ListNode{Val: 3},
			},
		},
		{
			name: "Four nodes - delete second middle (3)",
			args: args{
				head: &linkedlist.ListNode{
					Val: 1,
					Next: &linkedlist.ListNode{
						Val: 2,
						Next: &linkedlist.ListNode{
							Val:  3,
							Next: &linkedlist.ListNode{Val: 4},
						},
					},
				},
			},
			want: &linkedlist.ListNode{
				Val: 1,
				Next: &linkedlist.ListNode{
					Val:  2,
					Next: &linkedlist.ListNode{Val: 4},
				},
			},
		},
		{
			name: "Five nodes - delete middle (3)",
			args: args{
				head: &linkedlist.ListNode{
					Val: 1,
					Next: &linkedlist.ListNode{
						Val: 2,
						Next: &linkedlist.ListNode{
							Val: 3,
							Next: &linkedlist.ListNode{
								Val:  4,
								Next: &linkedlist.ListNode{Val: 5},
							},
						},
					},
				},
			},
			want: &linkedlist.ListNode{
				Val: 1,
				Next: &linkedlist.ListNode{
					Val: 2,
					Next: &linkedlist.ListNode{
						Val:  4,
						Next: &linkedlist.ListNode{Val: 5},
					},
				},
			},
		},
		{
			name: "Empty list",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
