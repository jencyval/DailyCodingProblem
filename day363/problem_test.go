package day363

import "testing"

// nolint
var testcases = []struct {
	input    []int
	expected int
}{
	{[]int{7}, 7},
	{[]int{1, 2, 3}, 0},
	{[]int{-5, 10, 3, 9}, 11},
}

func TestAddSubtract(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		as := AddSubtract(tc.input[0])
		for i := 1; i < len(tc.input); i++ {
			as = as.__(tc.input[i])
		}
		if result := as.Execute(); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkAddSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			as := AddSubtract(tc.input[0])
			for i := 1; i < len(tc.input); i++ {
				as = as.__(tc.input[i])
			}
			as.Execute()
		}
	}
}
