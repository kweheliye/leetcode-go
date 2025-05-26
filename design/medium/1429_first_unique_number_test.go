package medium

import (
	"testing"
)

func TestFirstUnique(t *testing.T) {
	// Test case 1: Basic functionality
	t.Run("Basic Functionality", func(t *testing.T) {
		// Initialize with [2,3,5]
		fu := Constructor([]int{2, 3, 5})
		
		// First unique should be 2
		if got := fu.ShowFirstUnique(); got != 2 {
			t.Errorf("ShowFirstUnique() = %v, want %v", got, 2)
		}
		
		// Add 5, which is already in the queue
		fu.Add(5)
		
		// First unique should still be 2
		if got := fu.ShowFirstUnique(); got != 2 {
			t.Errorf("ShowFirstUnique() = %v, want %v", got, 2)
		}
		
		// Add 2, which was unique but now becomes non-unique
		fu.Add(2)
		
		// First unique should now be 3
		if got := fu.ShowFirstUnique(); got != 3 {
			t.Errorf("ShowFirstUnique() = %v, want %v", got, 3)
		}
		
		// Add 3, which was unique but now becomes non-unique
		fu.Add(3)
		
		// First unique should now be 5, but 5 is also non-unique
		if got := fu.ShowFirstUnique(); got != -1 {
			t.Errorf("ShowFirstUnique() = %v, want %v", got, -1)
		}
	})
	
	// Test case 2: Empty initialization
	t.Run("Empty Initialization", func(t *testing.T) {
		// Initialize with empty array
		fu := Constructor([]int{})
		
		// No unique elements, should return -1
		if got := fu.ShowFirstUnique(); got != -1 {
			t.Errorf("ShowFirstUnique() = %v, want %v", got, -1)
		}
		
		// Add a unique element
		fu.Add(1)
		
		// First unique should be 1
		if got := fu.ShowFirstUnique(); got != 1 {
			t.Errorf("ShowFirstUnique() = %v, want %v", got, 1)
		}
	})
	
	// Test case 3: All duplicates
	t.Run("All Duplicates", func(t *testing.T) {
		// Initialize with all duplicates
		fu := Constructor([]int{1, 1, 2, 2, 3, 3})
		
		// No unique elements, should return -1
		if got := fu.ShowFirstUnique(); got != -1 {
			t.Errorf("ShowFirstUnique() = %v, want %v", got, -1)
		}
		
		// Add a unique element
		fu.Add(4)
		
		// First unique should be 4
		if got := fu.ShowFirstUnique(); got != 4 {
			t.Errorf("ShowFirstUnique() = %v, want %v", got, 4)
		}
	})
	
	// Test case 4: LeetCode example
	t.Run("LeetCode Example", func(t *testing.T) {
		// Initialize with [7,7,7,7,7,7]
		fu := Constructor([]int{7, 7, 7, 7, 7, 7})
		
		// No unique elements, should return -1
		if got := fu.ShowFirstUnique(); got != -1 {
			t.Errorf("ShowFirstUnique() = %v, want %v", got, -1)
		}
		
		// Add 7, which is already non-unique
		fu.Add(7)
		
		// Add 3, 3, 7, 17
		fu.Add(3)
		fu.Add(3)
		fu.Add(7)
		fu.Add(17)
		
		// First unique should be 17
		if got := fu.ShowFirstUnique(); got != 17 {
			t.Errorf("ShowFirstUnique() = %v, want %v", got, 17)
		}
	})
}