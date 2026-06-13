// Package main 提供快速排序（quicksort）的实现。
package main

import "fmt"

// QuickSort 对 nums 原地进行升序排序，使用经典的递归快速排序算法。
// 采用 Lomuto 分区方案，平均时间复杂度 O(n log n)，空间复杂度 O(log n)（递归栈）。
func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

// quickSort 对 nums[low..high] 这一闭区间进行排序。
func quickSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(nums, low, high)
	quickSort(nums, low, p-1)
	quickSort(nums, p+1, high)
}

// partition 以 nums[high] 作为基准，将区间划分为「小于基准」和「大于等于基准」两部分，
// 返回基准元素最终所在的下标。
func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[high] = nums[high], nums[i]
	return i
}

func main() {
	nums := []int{5, 2, 9, 1, 5, 6, 3, 8, 7, 4}
	fmt.Println("排序前:", nums)
	QuickSort(nums)
	fmt.Println("排序后:", nums)
}
