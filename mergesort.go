package main

import (
	"fmt"
)

func Merge(u, v, s []int) {

	ui := 0
	vi := 0

	i := 0
	for ui < len(u) && vi < len(v) {
		if u[ui] < v[vi] {
			s[i] = u[ui]
			ui++
		} else {
			s[i] = v[vi]
			vi++
		}
		i++
	}

	if ui >= len(u) {
		copy(s[i:], v[vi:])
	} else {
		copy(s[i:], u[ui:])
	}
}

func MergeSort(arr []int, leafSize int) {

	var recurse func([]int)
	recurse = func(subArr []int) {
		if len(subArr) < leafSize {
			InsertionSort(subArr)
			return
		}

		if len(subArr) == 1 {
			return
		}

		mid := len(subArr) / 2

		subLow := make([]int, mid)
		subHigh := make([]int, len(subArr)-mid)

		copy(subLow, subArr[0:mid])
		copy(subHigh, subArr[mid:])

		recurse(subLow)
		recurse(subHigh)
		Merge(subLow, subHigh, subArr)
	}

	recurse(arr)
}

func InsertionSort(arr []int) {
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
	u := []int{2, 7, 12, 13, 19, 21, 27}
	v := []int{1, 3, 4, 15, 16, 25}

	merged := make([]int, len(u)+len(v))
	Merge(u, v, merged)

	fmt.Printf("%v + %v = %v\n", u, v, merged)

	for leafSize := 0; leafSize < 5; leafSize++ {
		for _, arr := range [][]int{
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
			MergeSort(arr, leafSize)
			fmt.Printf("n=%d, leafSize=%d, arr=%v\n", len(arr), leafSize, arr)
		}
	}
}
