package sorting

import "testing"

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Test Case 1",
			args: []int{64, 34, 25, 12, 22, 11, 90},
			want: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name: "Test Case 2",
			args: []int{5, 2, 8, 12, 1, 3},
			want: []int{1, 2, 3, 5, 8, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertionSort(tt.args)
			for i := range tt.args {
				if tt.args[i] != tt.want[i] {
					t.Errorf("InsertionSort() = %v, want %v", tt.args, tt.want)
					break
				}
			}
		})
	}
}
