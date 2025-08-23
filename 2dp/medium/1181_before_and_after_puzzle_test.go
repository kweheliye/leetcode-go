package medium

import (
	"reflect"
	"testing"
)

func TestBeforeAndAfterPuzzles(t *testing.T) {
	type args struct {
		phrases []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test Case 1",
			args: args{
				phrases: []string{"sunny", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			},
			want: []string{"sunny day", "day is", "is the", "the sunny", "sunny the"},
		},
		{
			name: "Test Case 2",
			args: args{
				phrases: []string{"writing code", "code rocks"},
			},
			want: []string{"writing code rocks"},
		},
		{
			name: "Test Case 3",
			args: args{
				phrases: []string{"mission statement",
					"a quick bite to eat",
					"a chip off the old block",
					"chocolate bar",
					"mission impossible",
					"a man on a mission",
					"block party",
					"eat my words",
					"bar of soap"},
			},
			want: []string{"a chip off the old block party",
				"a man on a mission impossible",
				"a man on a mission statement",
				"a quick bite to eat my words",
				"chocolate bar of soap"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beforeAndAfterPuzzles(tt.args.phrases); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("beforeAndAfterPuzzles() = %v, want %v", got, tt.want)
			}
		})
	}
}
