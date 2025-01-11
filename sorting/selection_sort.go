package sorting

// SelectionSort Time Complexity: O(n^2), Space Complexity: O(1) (in-place)
func SelectionSort(arr []int) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		minIndex := findMinIndex(&arr, i)
		tmp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = tmp
	}
}

func findMinIndex(arr *[]int, index int) int {
	minIndex := index
	minValue := (*arr)[index]
	for i := index + 1; i < len(*arr); i++ {
		if (*arr)[i] < minValue {
			minValue = (*arr)[i]
			minIndex = i
		}
	}
	return minIndex
}
