package easy

import (
	"reflect"
	"testing"
)

func TestSumZero(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				n: 5,
			},
			want: []int{1, -1, 2, -2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumZero(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumZero() = %v, want %v", tt.want, got)
			}
		})
	}
}
