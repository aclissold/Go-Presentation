// +build OMIT

package main

import "fmt"

func main() {
    // START OMIT
    if name := "Andrew"; name == "Andrew" {
        fmt.Println("Last time I checked, my name was still Andrew!")
    } else if name == "Bob" {
        fmt.Println("I'm Bob now?!?! That can't be right.")
    } else {
        fmt.Println("I don't even know who I am.")
    }
    // fmt.Println(name, "was local to the if/else block and is out of scope here.")
    // END OMIT
}
