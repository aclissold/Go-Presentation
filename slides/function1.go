// +build OMIT

package main

import "fmt"

// START OMIT
func main() {
	a, b := 2, 3
	fmt.Println(pow(a, b))
}

func pow(x int, y int) int {
	ans := x
	for i := 1; i < y; i++ {
		ans *= x
	}
	return ans
}

// END OMIT
