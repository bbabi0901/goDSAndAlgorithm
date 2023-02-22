package main

import (
	"fmt"

	"github.com/bbabi0901/dataStructureAlgorithm/algorithm"
)

func main() {
	arr := []int{26, 3, 52, 62, 33, 12, 7}

	fmt.Println(algorithm.BubbleSort(arr))
	fmt.Println(algorithm.SelectionSort(arr))
}
