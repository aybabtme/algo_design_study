package main

import (
	"fmt"
)

type swapFunc func([]int, int, int)
type cmpFunc func(int, int) int

func SelectionSort(arr []int, swap swapFunc, cmp cmpFunc) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr); j++ {
			if cmp(arr[i], arr[j]) < 0 {
				swap(arr, i, j)
			}
		}
	}
}

func InsertionSort(arr []int, swap swapFunc, cmp cmpFunc) {
	for i := 0; i < len(arr); i++ {
		val := arr[i]
		hole := i
		for hole > 0 && val < arr[hole-1] {
			arr[hole] = arr[hole-1]
			hole--
		}
		arr[hole] = val
	}
}

func main() {

	for _, tt := range [][]int{
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{8, 7, 6, 5, 4, 3, 2, 1},
		{7, 6, 5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1},
		{5, 4, 3, 2, 1},
		{4, 3, 2, 1},
		{3, 2, 1},
		{2, 1},
		{1},
	} {
		// cmp, swap := CountSelectionSort(tt)
		cmp, swap := CountInsertionSort(tt)
		fmt.Printf("n=%d, cmp=%d, swap=%d, arr=%v\n", len(tt), cmp, swap, tt)
	}

}

func CountSelectionSort(arr []int) (compares int, swaps int) {
	compares = 0
	cmp := func(i, j int) int {
		compares++
		return i - j
	}

	swaps = 0
	swap := func(arr []int, i, j int) {
		swaps++
		arr[i], arr[j] = arr[j], arr[i]
	}
	SelectionSort(arr, swap, cmp)
	return
}

func CountInsertionSort(arr []int) (compares int, swaps int) {
	compares = 0
	cmp := func(i, j int) int {
		compares++
		return i - j
	}

	swaps = 0
	swap := func(arr []int, i, j int) {
		swaps++
		arr[i], arr[j] = arr[j], arr[i]
	}
	InsertionSort(arr, swap, cmp)
	return
}
