// +build OMIT

package main

import "fmt"
import "log"

func main() {
    a, b := 2, 3
    // Multiple return values!
    n, err := fmt.Println(pow(a, b))
    fmt.Printf("fmt.Println wrote %d bytes\n", n)
    fmt.Printf("and had no error (its value: %v)\n", err)
}

// Grouped parameters! Named return values!
func pow(x, y int) (ans int) {
    ans = x
    for i := 1; i < y; i++ {
        ans *= x
    }
    return
}
