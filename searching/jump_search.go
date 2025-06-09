package searching

import (
	"gorithms/utils"
	"math"
)

func JumpSearch(arr []int, target int) int {
	n := len(arr)
	step := int(math.Sqrt(float64(n)))
	prev := 0

	for arr[utils.MinValue(step, n)-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	for i := prev; i < utils.MinValue(step, n); i++ {
		if arr[i] == target {
			return i
		}
	}

	return -1
}
