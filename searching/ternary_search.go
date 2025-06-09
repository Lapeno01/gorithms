package searching

func TernarySearch(arr []int, target int) int {
	return ternarySearchHelper(arr, target, 0, len(arr)-1)
}

func ternarySearchHelper(arr []int, target, left, right int) int {
	if right < left {
		return -1
	}

	partitionSize := (right - left) / 3
	mid1 := left + partitionSize
	mid2 := right - partitionSize

	if arr[mid1] == target {
		return mid1
	}
	if arr[mid2] == target {
		return mid2
	}

	if target < arr[mid1] {
		return ternarySearchHelper(arr, target, left, mid1-1)
	} else if target > arr[mid2] {
		return ternarySearchHelper(arr, target, mid2+1, right)
	} else {
		return ternarySearchHelper(arr, target, mid1+1, mid2-1)
	}
}
