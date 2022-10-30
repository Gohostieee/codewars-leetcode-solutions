package main

// SplitString /*
/*

Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').
*/
func SplitString(str string) []string {
	length := len(str)
	str2 := []string{}
	index := 0
	for index < len(str) {
		if (index) != length-1 {
			str2 = append(str2, str[index:index+2])
			index += 2
		} else {
			str2 = append(str2, str[index:index+1]+string('_'))
			break
		}
	}

	return str2
}

func main() {
	SplitString("fuapz")
}
