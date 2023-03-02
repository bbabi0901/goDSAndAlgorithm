package main

import (
	"fmt"
	"math/rand"

	"github.com/bbabi0901/dataStructureAlgorithm/dataStructure"
)

func main() {
	/*
		arr := []int{26, 3, 52, 62, 33, 12, 7}
		fmt.Println(algorithm.BubbleSort(arr))
		fmt.Println(algorithm.SelectionSort(arr))
		fmt.Println(algorithm.InsertionSort(arr))
		fmt.Println(algorithm.QuickSort(arr))
		fmt.Println(algorithm.MergeSort(arr))
	*/

	/*
		start := time.Now()
		fmt.Println(toy.Fibo(40), time.Since(start))
		fmt.Println(toy.Tiling(4))
	*/

	treeRoot := dataStructure.MakeTreeNode(5)
	// 자식 4개 추가
	for i := 0; i < 4; i++ {
		// rand.Intn returns a random int n, 0 <= n < 100.
		treeRoot.AddNode(rand.Intn(100))
	}
	for i := 0; i < len(treeRoot.Children); i++ {
		//	손자 3개씩
		for j := 0; j < 3; j++ {
			treeRoot.Children[i].AddNode(rand.Intn(100))
		}
	}

	fmt.Println(dataStructure.DFSWithRecursive(treeRoot)...)
	fmt.Println(dataStructure.DFSWithStack(treeRoot)...)
	// Recursive랑 Stack이랑 탐색 방향이 다르다. -> vs <-
	fmt.Println(dataStructure.BFS(treeRoot)...)
}
