package sorting

func HeapSort(array []int) {
	n := len(array)

	// Step 1: Build a max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(array, n, i)
	}

	// Step 2: Extract elements from the heap
	for i := n - 1; i > 0; i-- {
		// Move the root (largest element) to the end
		array[0], array[i] = array[i], array[0]
		// Call heapify on the reduced heap
		heapify(array, i, 0)
	}
}

func heapify(array []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && array[left] > array[largest] {
		largest = left
	}

	if right < n && array[right] > array[largest] {
		largest = right
	}

	if largest != i {
		array[i], array[largest] = array[largest], array[i]
		heapify(array, n, largest)
	}
}
