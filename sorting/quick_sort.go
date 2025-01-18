package sorting

func QuickSort(array []int) {
	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, low, high int) {
	if low < high {
		pivotIndex := partition(array, low, high)
		quickSort(array, low, pivotIndex-1)
		quickSort(array, pivotIndex+1, high)
	}
}

func partition(array []int, low, high int) int {
	pivot := array[high]
	i := low - 1
	for j := low; j < high; j++ {
		if array[j] <= pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[high] = array[high], array[i+1]
	return i + 1
}
