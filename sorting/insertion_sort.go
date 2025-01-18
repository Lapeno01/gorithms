package sorting

// InsertionSort Time Complexity: O(n^2), Space Complexity: O(1) (in-place)
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i
		for j > 0 && arr[j-1] > tmp {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = tmp
	}
}
