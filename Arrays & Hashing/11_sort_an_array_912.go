package main

import "fmt"

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, low int, high int) {
	if low < high {
		pivot := partition(nums, low, high)
		quickSort(nums, low, pivot-1)
		quickSort(nums, pivot+1, high)
	}
}

func partition(nums []int, low int, high int) int {
	pivot := nums[high]
	i := low - 1

	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i] // swap
		}
	}

	// place pivot at correct position
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}
