package dedupe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func BenchmarkDedupe(b *testing.B) {
// 	words := []string{"abc", "bcd", "abcde", "cdefg", "bg"}
// 	for range b.N {

// 	}
// }

func TestDedupeStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "no duplicates",
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "with duplicates",
			input:    []string{"abc", "bcd", "abcde", "cdefg", "bg"},
			expected: []string{"abcde", "cdefg", "bg"},
		},
		{
			name:     "all duplicates",
			input:    []string{"a", "a", "a"},
			expected: []string{"a"},
		},
		{
			name:     "with empty strings",
			input:    []string{"", "a", "", "b", ""},
			expected: []string{"", "a", "b"},
		},
		{
			name:     "with spaces",
			input:    []string{" ", "a", " ", "b", "  "},
			expected: []string{" ", "  ", "a", "b"},
		},
	}
	s := NewSet()
	for _, tt := range tests {
		s.AddWords(tt.input)
		s.DedupeWords()
		t.Run(tt.name, func(t *testing.T) {
			result := s.Deduped()
			t.Log("result", result)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

// func TestDedupeInts(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    []int
// 		expected []int
// 	}{
// 		{
// 			name:     "empty slice",
// 			input:    []int{},
// 			expected: []int{},
// 		},
// 		{
// 			name:     "no duplicates",
// 			input:    []int{1, 2, 3},
// 			expected: []int{1, 2, 3},
// 		},
// 		{
// 			name:     "with duplicates",
// 			input:    []int{1, 2, 1, 3, 2},
// 			expected: []int{1, 2, 3},
// 		},
// 		{
// 			name:     "all duplicates",
// 			input:    []int{1, 1, 1},
// 			expected: []int{1},
// 		},
// 		{
// 			name:     "with zero values",
// 			input:    []int{0, 1, 0, 2, 0},
// 			expected: []int{0, 1, 2},
// 		},
// 		{
// 			name:     "negative numbers",
// 			input:    []int{-1, 1, -1, 0, 1},
// 			expected: []int{-1, 0, 1},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := Ints(tt.input)
// 			assert.ElementsMatch(t, tt.expected, result)
// 		})
// 	}
// }
