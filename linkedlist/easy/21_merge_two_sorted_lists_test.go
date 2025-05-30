package easy

import (
	"github.com/kweheliye/leetcode-go/linkedlist"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *linkedlist.ListNode
		l2 *linkedlist.ListNode
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.ListNode
	}{
		{
			name: "Test Case 1",
			args: args{
				l1: &linkedlist.ListNode{
					Val: 1,
					Next: &linkedlist.ListNode{
						Val: 2,
						Next: &linkedlist.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				l2: &linkedlist.ListNode{
					Val: 1,
					Next: &linkedlist.ListNode{
						Val: 3,
						Next: &linkedlist.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &linkedlist.ListNode{
				Val: 1,
				Next: &linkedlist.ListNode{
					Val: 1,
					Next: &linkedlist.ListNode{
						Val: 2,
						Next: &linkedlist.ListNode{
							Val: 3,
							Next: &linkedlist.ListNode{
								Val: 4,
								Next: &linkedlist.ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+" (Recursive)", func(t *testing.T) {
			if got := mergeTwoListsV1(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoListsV1() = %v, expected %v", got, tt.want)
			}
		})

		t.Run(tt.name+" (Iterative)", func(t *testing.T) {
			if got := mergeTwoListsV2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoListsV2() = %v, expected %v", got, tt.want)
			}
		})
	}
}
