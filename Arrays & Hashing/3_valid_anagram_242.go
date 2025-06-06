package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make([]int, 26)
	for i := 0; i < len(s); i++ {
		letters[s[i]-'a']++
		letters[t[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if letters[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("anagram", "nagaram"))
}
