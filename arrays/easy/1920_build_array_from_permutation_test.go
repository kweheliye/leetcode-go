package easy

import (
	"reflect"
	"testing"
)

func TestBuildArray(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "TestCase1",
			input:    []int{0, 2, 1, 5, 3, 4},
			expected: []int{0, 1, 2, 4, 5, 3},
		},
		{
			name:     "TestCase2",
			input:    []int{5, 0, 1, 2, 3, 4},
			expected: []int{4, 5, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("buildArray() = %v, want %v", got, tt.expected)
			}
		})
	}
}
