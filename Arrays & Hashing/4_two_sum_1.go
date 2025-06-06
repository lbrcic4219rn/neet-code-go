package main

import "fmt"

func twoSum(nums []int, target int) []int {
	twoSumMap := make(map[int]int)
	for i, num := range nums {
		index, exists := twoSumMap[target-num]
		if exists {
			return []int{index, i}
		}
		twoSumMap[num] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
