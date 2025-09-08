package easy

import (
	"reflect"
	"testing"
)

func TestGetNoZeroIntegers(t *testing.T) {
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
				n: 201,
			},
			want: []int{2, 199},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNoZeroIntegers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNoZeroIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
