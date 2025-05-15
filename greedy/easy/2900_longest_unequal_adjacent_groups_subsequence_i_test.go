package easy

import (
	"reflect"
	"testing"
)

func TestGetLongestSubsequence(t *testing.T) {
	type args struct {
		words  []string
		groups []int
	}

	testCases := []struct {
		name     string
		args     args
		expected []string
	}{
		{
			name: "Test Case 1",
			args: args{
				words:  []string{"e", "a", "b"},
				groups: []int{0, 0, 1},
			},
			expected: []string{"e", "b"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestSubsequence(tt.args.words, tt.args.groups); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("getLongestSubsequence() = %v, want %v", got, tt.expected)

			}

		})
	}

}
