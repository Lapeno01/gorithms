package sorting

// BubbleSort Time Complexity: O(n^2), Space Complexity: O(1) (in-place)
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		swapped := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				swapped = true
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
		if !swapped {
			return
		}
	}
}
