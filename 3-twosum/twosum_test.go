package twosum_test

import (
	twosum "neetcode150/3-twosum"
	"reflect"
	"testing"
)

func TestTwosum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{15, 7, 2, 11}, 9, []int{1, 2}},
		{[]int{-1, 32, 4, 24, 43, 1}, 0, []int{0, 5}},
	}

	for _, tt := range tests {
		got := twosum.Twosum_bf(tt.nums, tt.target)

		// Check if got and expected slices are equal
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("Twosum_bf(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.expected)
		}
	}
}

func TestTwosum_HashMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{15, 7, 2, 11}, 9, []int{1, 2}},
		{[]int{-1, 32, 4, 24, 43, 1}, 0, []int{0, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{0, 1}},
		{[]int{1, 0, 0, 0, 0, 1}, 2, []int{0, 5}},
		{[]int{1, 0, 0, 0, 0, 1}, 1, []int{0, 1}},
	}

	for _, tt := range tests {
		got := twosum.Twosum_bf(tt.nums, tt.target)

		// Check if got and expected slices are equal
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("Twosum_bf(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.expected)
		}
	}
}
