package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	dpArr := make([]int, len(cost))
	dpArr[0] = cost[0]
	if len(cost) == 1 {
		return cost[0]
	}

	dpArr[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		if dpArr[i-1]+cost[i] < dpArr[i-2]+cost[i] {
			dpArr[i] = dpArr[i-1] + cost[i]
		} else {
			dpArr[i] = dpArr[i-2] + cost[i]
		}
	}

	fmt.Println("dpArr: ", dpArr)

	if dpArr[len(cost)-1] < dpArr[len(cost)-2] {
		return dpArr[len(cost)-1]
	}
	return dpArr[len(cost)-2]
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))                         // 15
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) //6
}
