package main

import (
	"fmt"
	"time"

	"github.com/bbabi0901/dataStructureAlgorithm/algorithm"
	"github.com/bbabi0901/dataStructureAlgorithm/toy"
)

func main() {
	arr := []int{26, 3, 52, 62, 33, 12, 7}

	fmt.Println(algorithm.BubbleSort(arr))
	fmt.Println(algorithm.SelectionSort(arr))
	fmt.Println(algorithm.InsertionSort(arr))
	fmt.Println(algorithm.QuickSort(arr))
	fmt.Println(algorithm.MergeSort(arr))

	start := time.Now()
	fmt.Println(toy.Fibo(40), time.Since(start))
	fmt.Println(toy.Tiling(4))
}
