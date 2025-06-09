package searching

func InterpolationSearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right && target >= arr[left] && target <= arr[right] {
		if left == right {
			if arr[left] == target {
				return left
			}
			return -1
		}

		pos := left + ((target-arr[left])*(right-left))/(arr[right]-arr[left])
		if pos < 0 || pos >= len(arr) {
			return -1
		}

		if arr[pos] == target {
			return pos
		} else if arr[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}

	return -1
}
