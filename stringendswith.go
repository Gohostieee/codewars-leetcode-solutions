package main

import "fmt"

func solution(str, ending string) bool {
	if str[len(str)-len(ending):] == ending {
		return true

	}

	return false

}

func main() {
	if solution("fuap", "k") {
		fmt.Println("fuap")
	}
}
