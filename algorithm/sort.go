package algorithm

import "golang.org/x/exp/constraints"

// constraints.Ordered는 대소 비교가 가능한 타입의 집합(int, float, string)
func swap[T constraints.Ordered](idx1, idx2 int, arr []T) {
	arr[idx1], arr[idx2] = arr[idx2], arr[idx1]
}

func BubbleSort[T constraints.Ordered](arr []T) []T {
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

func SelectionSort[T constraints.Ordered](arr []T) []T {
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

func InsertionSort[T constraints.Ordered](arr []T) []T {
	var j int = 0
	for i := 1; i < len(arr); i++ {
		var curr T = arr[i]
		for j = i - 1; j >= 0; j-- {
			if curr < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = curr
	}
	return arr
}

func QuickSort[T constraints.Ordered](arr []T) []T {
	// pivot 정해서 보다 작은 수, 큰 수로 분류. 분류를 반복
	if len(arr) <= 1 {
		return arr
	}

	var pivot T = arr[0]
	var left, right []T

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	var leftSorted []T = QuickSort(left)
	var rightSorted []T = QuickSort(right)

	return append(append(leftSorted, pivot), rightSorted...)
}

func merge[T constraints.Ordered](left, right []T) []T {
	var leftIdx, rightIdx int = 0, 0
	var merged []T
	totalSize := len(left) + len(right)

	for i := 0; i < totalSize; i++ {
		if leftIdx >= len(left) {
			merged = append(merged, right[rightIdx])
			rightIdx++
		} else if rightIdx >= len(right) || left[leftIdx] <= right[rightIdx] {
			merged = append(merged, left[leftIdx])
			leftIdx++
		} else {
			merged = append(merged, right[rightIdx])
			rightIdx++
		}
	}
	return merged
}

func MergeSort[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	var left []T = MergeSort(arr[:mid])
	var right []T = MergeSort(arr[mid:])
	var merged = merge(left, right)
	return merged
}
