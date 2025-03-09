package validanagram_test

import (
	validanagram "neetcode150/1-validanagram" // Correct import
	"testing"
)

func TestValidAnagram(t *testing.T) {
	t.Parallel()

	tests := []struct {
		word     string
		test     string
		expected bool
	}{
		{"elbow", "below", true},
		{"listen", "silent", true},
		{"rat", "car", false},
		{"abc", "ab", false},
		{"a", "a", true},
	}

	for _, tt := range tests {
		got := validanagram.ValidAnagram(tt.word, tt.test)
		if got != tt.expected {
			t.Errorf("ValidAnagram(%q, %q) = %v; want %v", tt.word, tt.test, got, tt.expected)
		}
	}
}
