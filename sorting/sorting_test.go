package sorting

import (
	"reflect"
	"testing"
)

var sortingTests = []struct {
	input []int
	want  []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
	{[]int{-2, 0, -1}, []int{-2, -1, 0}},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range sortingTests {
		got := make([]int, len(tt.input))
		copy(got, tt.input)

		BubbleSort(got)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("BubbleSort(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range sortingTests {
		got := make([]int, len(tt.input))
		copy(got, tt.input)

		SelectionSort(got)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SelectionSort(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
