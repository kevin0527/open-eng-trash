package main

import "fmt"

// QuickSort sorts a slice of ints in ascending order in place using the
// quicksort algorithm. The original slice is mutated and also returned for
// convenience.
func QuickSort(a []int) []int {
	quicksort(a, 0, len(a)-1)
	return a
}

// quicksort recursively sorts the sub-slice a[lo..hi] (inclusive).
func quicksort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quicksort(a, lo, p-1)
	quicksort(a, p+1, hi)
}

// partition uses the Lomuto scheme with the last element as the pivot. It
// places the pivot at its final sorted position and returns that index.
func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

func main() {
	data := []int{5, 2, 9, 1, 5, 6, 3, 8, 7, 0, 4}
	fmt.Println("before:", data)
	QuickSort(data)
	fmt.Println("after: ", data)
}
