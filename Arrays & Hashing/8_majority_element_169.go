package main

import "fmt"

func majorityElement(nums []int) int {
	current := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == current {
			count++
			continue
		}
		count--
		if count == 0 {
			count = 1
			current = nums[i]
		}
	}
	return current
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
