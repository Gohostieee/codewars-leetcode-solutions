package main

import (
	"fmt"
	"strings"
)

/*
Build a pyramid-shaped tower, as an array/list of strings, given a positive integer number of floors. A tower block is represented with "*" character.
*/
func TowerBuilder(nFloors int) []string {
	str := []string{}
	stars := 1
	spaces := 0
	index := 1
	floor := ""
	for index < nFloors {
		stars += 2
		index++
	}
	for index > 0 {
		floor = ""
		floor = floor + strings.Repeat(" ", spaces)
		floor = floor + strings.Repeat("*", stars)
		floor = floor + strings.Repeat(" ", spaces)
		stars -= 2
		spaces += 2
		index--
		str = append([]string{floor}, str...)
	}
	return str
}

func main() {
	fmt.Printf("|%s|", TowerBuilder(6))
}
