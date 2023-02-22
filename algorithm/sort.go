package algorithm

func swap(idx1, idx2 int, arr []int) {
	arr[idx1], arr[idx2] = arr[idx2], arr[idx1]
}

func BubbleSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap(j, j+1, arr)
			}
		}
	}
	return arr
}

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(i, minIndex, arr)
	}
	return arr
}
