// +build OMIT

package main

import "fmt"

func main() {
	// START OMIT
	words := []string{"apple", "banana", "cantaloupe"}
	for i, word := range words {
		fmt.Println(word, "is located at index", i)
	}
	// END OMIT
}
