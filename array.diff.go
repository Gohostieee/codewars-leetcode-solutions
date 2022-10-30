//go:build !windows
// +build !windows

package main

import "fmt"

func ArrayDiff(a, b []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				i--
				a = append(a[:i], a[i+1:]...)
			}
		}
	}

	// your code here
	return a
}
func main() {
	a := []int{1, 2, 2}
	b := []int{2}

	fmt.Printf("%d", len(ArrayDiff(a, b)))
}
