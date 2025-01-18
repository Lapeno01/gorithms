package searching

import (
	"testing"
)

var searchingTests = []struct {
	input         []int
	target        int
	expectedIndex int
}{
	{[]int{1, 2, 3, 4, 5, 6}, 2, 1},
	{[]int{1, 2, 3, 4, 5, 6}, 6, 5},
	{[]int{1, 2, 3, 4, 5, 6}, -9, -1},
}

type searchFunc func([]int, int) int

func TestSearching(t *testing.T) {
	searchAlgorithms := map[string]searchFunc{
		"BinarySearch": BinarySearch,
		"LinearSearch": LinearSearch,
	}

	for name, function := range searchAlgorithms {
		for _, tt := range searchingTests {
			t.Run(name, func(t *testing.T) {
				result := function(tt.input, tt.target)
				if result != tt.expectedIndex {
					t.Errorf("%s failed for input %v and target %d: expected %d, got %d",
						name, tt.input, tt.target, tt.expectedIndex, result)
				}
			})
		}
	}
}
