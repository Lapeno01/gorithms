package sorting

func MergeSort(array []int) {
	left := 0
	right := len(array) - 1
	mergeSort(array, left, right)
}

func mergeSort(array []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(array, left, mid)
		mergeSort(array, mid+1, right)
		merge(array, left, mid, right)
	}
}

func merge(array []int, left int, mid int, right int) {
	len1 := mid - left + 1
	len2 := right - mid
	arr1 := make([]int, len1)
	arr2 := make([]int, len2)

	for i := 0; i < len1; i++ {
		arr1[i] = array[left+i]
	}
	for i := 0; i < len2; i++ {
		arr2[i] = array[mid+i+1]
	}

	i, j, k := 0, 0, left
	for i < len1 && j < len2 {
		if arr1[i] <= arr2[j] {
			array[k] = arr1[i]
			i++
		} else {
			array[k] = arr2[j]
			j++
		}
		k++
	}

	// Copy remaining elements of arr1, if any
	for i < len1 {
		array[k] = arr1[i]
		i++
		k++
	}

	// Copy remaining elements of arr2, if any
	for j < len2 {
		array[k] = arr2[j]
		j++
		k++
	}
}
