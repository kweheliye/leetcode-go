package easy

import (
	"reflect"
	"testing"
)

func Test_resultArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{2, 1, 3},
			},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resultArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resultArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
