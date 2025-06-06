package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		for _, str := range strs {
			if i == len(str) || strs[0][i] != str[i] {
				return str[:i]
			}
		}
	}
	return strs[0]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
