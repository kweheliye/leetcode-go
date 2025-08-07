package medium

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	numMatrix := ConstructorBruteForce(matrix)

	if len(numMatrix.data) != len(matrix) {
		t.Errorf("ConstructorBruteForce() failed: expected %v rows, got %v", len(matrix), len(numMatrix.data))
	}
}

func TestNumMatrix_SumRegionBruteForce(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	tests := []struct {
		name   string
		matrix [][]int
		args   struct {
			row1, col1, row2, col2 int
		}
		want int
	}{
		{
			name:   "Example 1",
			matrix: matrix,
			args:   struct{ row1, col1, row2, col2 int }{2, 1, 4, 3},
			want:   8,
		},
		{
			name:   "Example 2",
			matrix: matrix,
			args:   struct{ row1, col1, row2, col2 int }{1, 1, 2, 2},
			want:   11,
		},
		{
			name:   "Example 3",
			matrix: matrix,
			args:   struct{ row1, col1, row2, col2 int }{1, 2, 2, 4},
			want:   12,
		},
		{
			name:   "Single Cell",
			matrix: matrix,
			args:   struct{ row1, col1, row2, col2 int }{0, 0, 0, 0},
			want:   3,
		},
		{
			name:   "Entire Matrix",
			matrix: matrix,
			args:   struct{ row1, col1, row2, col2 int }{0, 0, 4, 4},
			want:   58,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := ConstructorBruteForce(tt.matrix)
			got := nm.SumRegionBruteForce(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2)
			if got != tt.want {
				t.Errorf("SumRegionBruteForce() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			nm := ConstructorOneDimensionalPrefixSum(tt.matrix)
			got := nm.SumRegionOneDimensionalPrefixSum(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2)
			if got != tt.want {
				t.Errorf("SumRegionOneDimensionalPrefixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
