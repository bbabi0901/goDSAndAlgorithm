package algorithm

func BinarySesarch(list []int, target int) int {
	result := -1
	lower := 0
	upper := len(list)
	for lower <= upper {
		middle := (lower + upper) / 2
		if target > list[middle] {
			lower = list[middle] + 1
		} else if target < list[middle] {
			upper = list[middle] - 1
		} else {
			return middle
		}
	}
	return result
}

// Go는 재귀함수를 지원한다.
func RecursiveBinarySesarch(list []int, target, lower, upper int) int {
	if lower < upper {
		return -1
	}
	middle := (lower + upper) / 2
	if target == list[middle] {
		return middle
	} else if target > list[middle] {
		return RecursiveBinarySesarch(list, target, middle+1, upper)
	} else {
		return RecursiveBinarySesarch(list, target, lower, middle-1)
	}
}
