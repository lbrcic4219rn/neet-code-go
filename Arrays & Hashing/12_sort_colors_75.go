package main

import "fmt"

func sortColors(nums []int) {
	redCount := 0
	whiteCount := 0
	blueCount := 0

	for _, num := range nums {
		switch num {
		case 0:
			redCount++
		case 1:
			whiteCount++
		case 2:
			blueCount++
		}
	}

	for i := 0; i < len(nums); i++ {
		if redCount > 0 {
			nums[i] = 0
			redCount--
			continue
		}
		if whiteCount > 0 {
			nums[i] = 1
			whiteCount--
			continue
		}
		if blueCount > 0 {
			nums[i] = 2
			blueCount--
			continue
		}
	}
}

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)
	arr = []int{2, 0, 1}
	sortColors(arr)
	fmt.Println(arr)
}
