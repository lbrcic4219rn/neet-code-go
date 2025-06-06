package main

import "fmt"

func getConcatenation(nums []int) []int {
	result := make([]int, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		result[i] = nums[i]
		result[i+len(nums)] = nums[i]
	}
	return result
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 3, 4, 5}))
}
