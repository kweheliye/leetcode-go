package easy

import (
	"reflect"
	"testing"
)

func TestFindWordsContaining(t *testing.T) {
	type args struct {
		words []string
		x     byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				words: []string{"cat", "bt", "hat", "rider"},
				x:     't',
			},
			want: []int{0, 1, 2},
		},
		{
			name: "Test Case 2",
			args: args{
				words: []string{"hello", "world", "leetcode"},
				x:     'z',
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWordsContaining(tt.args.words, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWordsContaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
