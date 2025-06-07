package main

import "fmt"

func climbStairs(n int) int {
	stepWayCount := make([]int, n)
	stepWayCount[0] = 1
	stepWayCount[1] = 2
	for i := 2; i < n; i++ {
		stepWayCount[i] = stepWayCount[i-1] + stepWayCount[i-2]
	}
	return stepWayCount[n-1]
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(5))
}
