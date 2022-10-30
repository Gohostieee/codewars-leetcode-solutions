package main

import "fmt"

func GetCount(str string) (count int) {
	// Enter solution here
	for _, x := range str {
		switch x {
		case 'a', 'e', 'o', 'i', 'u':
			count++
		}
	}
	return count
}
func main() {
	fmt.Printf("%d", GetCount("aesssdiddddodu"))
}
