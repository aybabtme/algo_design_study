package main

import (
	"fmt"
)

func IsSorted(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}

	var n = arr[0]
	var m int
	for i := 1; i < len(arr); i++ {
		m = arr[i]
		if n > m {
			return false
		}
		n = m
	}
	return true
}

func TestIsSorted() {
	sorted := []int{0, 1, 2, 3, 4}
	notSorted := []int{0, 1, 2, 4, 3}

	fmt.Printf("IsSorted=%v, arr=%v\n", IsSorted(sorted), sorted)
	fmt.Printf("IsSorted=%v, arr=%v\n", IsSorted(notSorted), notSorted)
}

func LinearSearch(arr []int, val int) (int, error) {
	for i, curVal := range arr {
		if curVal == val {
			return i, nil
		}
	}
	return -1, fmt.Errorf("%d not found", val)
}

func BinaryIterSearch(arr []int, val int) (int, error) {

	if !IsSorted(arr) {
		return -1, fmt.Errorf("array is not sorted")
	}

	low := 0
	high := len(arr) - 1

	if arr[high] < val {
		return -1, fmt.Errorf("%d is not in the array (too high)", val)
	}

	var mid int

	for low <= high {
		mid = (low + high) / 2
		if val == arr[mid] {
			return mid, nil
		}
		if val < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, fmt.Errorf("%d not found", val)
}

func BinaryRecurSearch(arr []int, val int) (int, error) {

	if !IsSorted(arr) {
		return -1, fmt.Errorf("array is not sorted")
	}

	low := 0
	high := len(arr) - 1

	if arr[high] < val {
		return -1, fmt.Errorf("%d is not in the array (too high)", val)
	}

	var recurse func(int, int) (int, error)
	recurse = func(low, high int) (int, error) {

		if low > high {
			return -1, fmt.Errorf("%d not found", val)
		}

		mid := (high + low) / 2.0

		if arr[mid] == val {
			return mid, nil
		}

		if arr[mid] < val {

			return recurse(mid+1, high)
		}

		return recurse(low, mid-1)
	}

	return recurse(low, high)
}

func TestBinaryRecurSearch() {
	sorted := []int{0, 1, 2, 3, 5}
	notSorted := []int{0, 1, 2, 4, 3}

	for _, sort := range []struct {
		arr  []int
		want int
		find int
	}{
		{sorted, 0, 0},
		{sorted, 1, 1},
		{sorted, 2, 2},
		{sorted, 3, 3},
		{sorted, -1, 4},
		{sorted, 4, 5},
		{sorted, -1, 6},
		{notSorted, -1, 2},
	} {
		got, err := BinaryRecurSearch(sort.arr, sort.find)
		fmt.Printf("BinaryIterSearch=%d (want %d), err=%v, arr=%v\n", got, sort.want, err, sort.arr)
	}
}

func TestBinaryIterSearch() {
	sorted := []int{0, 1, 2, 3, 5}
	notSorted := []int{0, 1, 2, 4, 3}

	for _, sort := range []struct {
		arr  []int
		want int
		find int
	}{
		{sorted, 0, 0},
		{sorted, 1, 1},
		{sorted, 2, 2},
		{sorted, 3, 3},
		{sorted, -1, 4},
		{sorted, 4, 5},
		{sorted, -1, 6},
		{notSorted, -1, 2},
	} {
		got, err := BinaryIterSearch(sort.arr, sort.find)
		fmt.Printf("BinaryIterSearch=%d (want %d), err=%v, arr=%v\n", got, sort.want, err, sort.arr)
	}
}

func main() {
	// TestIsSorted()
	// TestBinaryIterSearch()
	TestBinaryRecurSearch()
}
