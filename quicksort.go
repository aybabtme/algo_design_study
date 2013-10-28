package main

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/plotutil"
	"code.google.com/p/plotinum/vg"
	"code.google.com/p/plotinum/vg/vgsvg"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func SelectPivot(arr []int, low, high int) int {
	width := high - low
	if width < 3 {
		return low
	}

	mid := width / 2

	if arr[low] < arr[mid] {
		if arr[low] >= arr[high] {
			return low
		} else if arr[mid] < arr[high] {
			return mid
		}
	}
	if arr[low] < arr[high] {
		return low
	}
	if arr[mid] >= arr[high] {
		return mid
	}
	return high
}

func Partition(arr []int, low, high int) int {
	pivotIdx := SelectPivot(arr, low, high)
	pivotVal := arr[pivotIdx]

	arr[low], arr[pivotIdx] = arr[pivotIdx], arr[low]

	l := low
	r := high

	for {
		// Find item on left to swap, not past `r`
		for l < r && arr[l] <= pivotVal {
			l++
		}

		// Find item on right to swap, not more than 1 past `l`
		for r >= l && arr[r] > pivotVal {
			r--
		}

		// If we crossed, nothing to swap
		if l >= r {
			// Didn't find anything
			break
		}
		// Found something, so swap them
		arr[l], arr[r] = arr[r], arr[l]
	}

	// Put pivotVal where `r` stopped
	arr[r], arr[low] = arr[low], arr[r]
	return r
}

func QuickSort(arr []int, leafSize int) {

	var recurse func(int, int)
	recurse = func(low, high int) {
		if high <= low {
			return
		}

		if high-low < leafSize {
			InsertionSort(arr[low : high+1])
			return
		}
		pivot := Partition(arr, low, high)

		recurse(low, pivot-1)
		recurse(pivot+1, high)
	}

	recurse(0, len(arr)-1)
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
	rand.Seed(time.Now().Unix())

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	n := 1 << 20

	xy := make(plotter.XYs, 0)

	for trial := 0; trial < 20; trial++ {
		for leafSize := 0; leafSize < 128; leafSize++ {

			arr := GenerateArray(n)

			start := time.Now()
			QuickSort(arr, leafSize)
			dT := time.Since(start)

			val := struct {
				X float64
				Y float64
			}{
				X: float64(leafSize),
				Y: dT.Seconds(),
			}

			fmt.Print(".")

			xy = append(xy, val)

			if !IsSorted(arr) {
				fmt.Printf("NOT SORTED, n=%d, arr=%v\n", n, arr)
			}
		}
	}

	p.Title.Text = fmt.Sprintf("Quicksort performance n = %d", n)
	p.X.Label.Text = "Leaf Size"
	p.Y.Label.Text = "Time (s)"

	err = plotutil.AddScatters(p, xy)
	if err != nil {
		panic(err)
	}

	w, h := vg.Centimeters(20), vg.Centimeters(10)
	c := vgsvg.New(w, h)
	da := plot.MakeDrawArea(c)
	p.Draw(da)
	file, err := os.Create("quicksort_leafsize.svg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	c.WriteTo(file)
	// p.Save(8, 4, "quicksort_leafsize.png")
}

func GenerateArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

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
