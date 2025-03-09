package containsduplicate_test

import (
	containsduplicate "neetcode150/2-containsduplicate"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{1, 2, 3, 4, 1}, true},
		{[]int{}, false},
		{[]int{7, 7, 7, 7}, true},
		{[]int{10, 20, 30, 40, 50}, false},
		{[]int{-1, -2, 93, 2, -2}, true},
	}

	for _, tt := range tests {
		got := containsduplicate.ContainsDuplicate(tt.nums)
		if got != tt.expected {
			t.Errorf("ContainsDuplicate(%v) = %v; want %v", tt.nums, got, tt.expected)
		}
	}
}
