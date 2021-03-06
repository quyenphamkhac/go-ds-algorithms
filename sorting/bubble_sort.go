package sorting

// BubbleSort implement
func BubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

// BubbleSort2 enhanced version implement
func BubbleSort2(arr []int, comp func(int, int) bool) {
	size := len(arr)
	swapped := 1
	for i := 0; i < size-1 && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = 1
			}
		}
	}
}
