package main

import (
	"fmt"
)

/*
Your task is to construct a building which will be a pile of n cubes. The cube at the bottom will have a volume of n3 n^3n
3
 , the cube above will have volume of (n−1)3 (n-1)^3(n−1)
3
  and so on until the top which will have a volume of 13 1^31
3
 .

You are given the total volume m of the building. Being given m can you find the number n of cubes you will have to build?

The parameter of the function findNb (find_nb, find-nb, findNb, ...) will be an integer m and you have to return the integer n such as n3+(n−1)3+...+13=m n^3 + (n-1)^3 + ... + 1^3 = mn
3
 +(n−1)
3
 +...+1
3
 =m if such a n exists or -1 if there is no such n.


*/

func FindNb(m int) int {
	// your code
	volume := 1
	cubeCount := 0
	for m != volume {
		fmt.Printf("%d %d\n", volume, m)
		if volume > m {
			return -1
		}
		cubeCount++
		volume += (cubeCount + 1) * (cubeCount + 1) * (cubeCount + 1)

	}

	return cubeCount + 1
}
func main() {
	fmt.Printf("%d", FindNb(24723578342962))
}
