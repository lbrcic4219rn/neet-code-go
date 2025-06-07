package main

import "fmt"

func tribonacci(n int) int {
	dpArr := make([]int, n+1)
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}
	dpArr[0] = 0
	dpArr[1] = 1
	dpArr[2] = 1

	for i := 3; i <= n; i++ {
		dpArr[i] = dpArr[i-1] + dpArr[i-2] + dpArr[i-3]
	}

	return dpArr[n]
}

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
}
