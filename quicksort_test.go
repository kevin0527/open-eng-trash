package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"single", []int{42}, []int{42}},
		{"sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{3, 1, 2, 3, 1, 2}, []int{1, 1, 2, 2, 3, 3}},
		{"negatives", []int{0, -3, 5, -1, 2}, []int{-3, -1, 0, 2, 5}},
		{"mixed", []int{5, 2, 9, 1, 5, 6, 3, 8, 7, 0, 4}, []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := QuickSort(tc.in)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("QuickSort(%v) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}

// TestQuickSortMatchesStdlib fuzzes QuickSort against the standard library
// sort over many random slices of varying lengths.
func TestQuickSortMatchesStdlib(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for iter := 0; iter < 1000; iter++ {
		n := rng.Intn(200)
		in := make([]int, n)
		for i := range in {
			in[i] = rng.Intn(100) - 50
		}
		want := make([]int, n)
		copy(want, in)
		sort.Ints(want)

		got := QuickSort(in)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("mismatch on iter %d: got %v, want %v", iter, got, want)
		}
	}
}

func benchmarkQuickSort(b *testing.B, n int) {
	rng := rand.New(rand.NewSource(1))
	base := make([]int, n)
	for i := range base {
		base[i] = rng.Int()
	}
	buf := make([]int, n)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(buf, base)
		QuickSort(buf)
	}
}

func BenchmarkQuickSort100(b *testing.B)   { benchmarkQuickSort(b, 100) }
func BenchmarkQuickSort1000(b *testing.B)  { benchmarkQuickSort(b, 1000) }
func BenchmarkQuickSort10000(b *testing.B) { benchmarkQuickSort(b, 10000) }
