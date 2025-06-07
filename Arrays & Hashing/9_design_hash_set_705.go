package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	fmt.Println(obj.Contains(1)) // true
	fmt.Println(obj.Contains(3)) // false
	obj.Add(2)
	fmt.Println(obj.Contains(2)) // true
	obj.Remove(2)
	fmt.Println(obj.Contains(2))
}
