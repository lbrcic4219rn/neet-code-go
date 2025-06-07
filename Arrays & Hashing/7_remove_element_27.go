package main

import "fmt"

func removeElement(nums []int, val int) int {
	index := 0
	for _, num := range nums {
		if num == val {
			continue
		}
		nums[index] = num
		index++
	}
	return index
}

func main() {
	arr := []int{3, 2, 2, 3}
	fmt.Println(removeElement(arr, 3))
	fmt.Println(arr)
	arr = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(arr, 2))
	fmt.Println(arr)
}
