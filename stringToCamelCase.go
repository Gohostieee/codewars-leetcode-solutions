package main

import (
	"fmt"
	"strings"
)

/*
Complete the method/function so that it converts dash/underscore delimited words into camel casing. The first word within the output should be capitalized only if the original word was capitalized (known as Upper Camel Case, also often referred to as Pascal case).
*/
func ToCamelCase(s string) string {
	// your code
	length := len(s)
	index := 0

	for index < length {

		if s[index] == '-' || s[index] == '_' {
			fmt.Printf("%s\n", s)
			s = s[:index] + s[index+1:]
			s = s[:index] + strings.ToUpper(string(s[index])) + s[index+1:]
			index--
			length--
		} else {
			index++

		}
	}
	return s
}
func main() {
	fmt.Printf("%s", ToCamelCase("Samurai_down_Black_Yellow_side_mountain	"))
}
