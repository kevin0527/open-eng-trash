package main

import (
	"math/rand"
	"sort"
	"testing"
)

// TestQuickSort 覆盖空切片、单元素、已排序、逆序、含重复元素及负数等情况。
func TestQuickSort(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{3, 1, 2, 3, 1}, []int{1, 1, 2, 3, 3}},
		{"negatives", []int{0, -2, 5, -1, 3}, []int{-2, -1, 0, 3, 5}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			QuickSort(c.in)
			if !equal(c.in, c.want) {
				t.Errorf("QuickSort() = %v, want %v", c.in, c.want)
			}
		})
	}
}

// TestQuickSortRandom 用随机数据与标准库 sort.Ints 对照，验证正确性。
func TestQuickSortRandom(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	for i := 0; i < 100; i++ {
		n := r.Intn(200)
		nums := make([]int, n)
		for j := range nums {
			nums[j] = r.Intn(1000) - 500
		}
		want := make([]int, n)
		copy(want, nums)
		sort.Ints(want)

		QuickSort(nums)
		if !equal(nums, want) {
			t.Fatalf("随机用例排序错误\n got = %v\nwant = %v", nums, want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// BenchmarkQuickSort 对 1000 个随机整数排序进行基准测试。
func BenchmarkQuickSort(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	base := make([]int, 1000)
	for i := range base {
		base[i] = r.Intn(10000)
	}

	buf := make([]int, len(base))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(buf, base)
		QuickSort(buf)
	}
}
