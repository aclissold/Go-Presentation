// +build OMIT

package main

import "fmt"

// START OMIT
type parent struct {
	father   bool
	age      uint8
	children []string
}

func main() {
	children := make([]string, 3)
	children[0], children[1], children[2] = "Junior", "Jack", "Jill"
	john := &parent{true, 42, children}
	fmt.Printf("%+v\n", john)
	// john.bobbify()
	// fmt.Println(john)
}

func (p *parent) bobbify() {
	if p.father {
		for i, child := range p.children {
			fmt.Println("Replacing", child, "with Bob.")
			p.children[i] = "Bob"
		}
	}
}
// END OMIT
