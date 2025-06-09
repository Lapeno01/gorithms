package searching

import "gorithms/utils"

func ExponentialSearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	if arr[0] == target {
		return 0
	}

	i := 1
	for i < len(arr) && arr[i] <= target {
		i = i * 2
	}

	return BinarySearch(arr[:utils.MinValue(i, len(arr))], target)
}
