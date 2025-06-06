package main

import "fmt"

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		_, exists := set[nums[i]]
		if exists {
			return true
		}
		set[nums[i]] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 3}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
}
