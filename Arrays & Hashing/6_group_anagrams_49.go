package main

import (
	"fmt"
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		chars := make([]int, 26)
		for i := 0; i < len(str); i++ {
			chars[str[i]-'a']++
		}

		var builder strings.Builder
		for i := 0; i < 26; i++ {
			builder.WriteString(strconv.Itoa(chars[i]))
			builder.WriteByte(',')
		}
		stringArr := builder.String()
		anagramMap[stringArr] = append(anagramMap[stringArr], str)
	}

	var result [][]string
	for _, value := range anagramMap {
		result = append(result, value)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"bdddddddddd", "bbbbbbbbbbc"}))
}
