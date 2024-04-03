package main

import (
	"fmt"
)

func main() {

	// var a int = 10
	// var b *int = &a
	// var c *int= &b

	// fmt.Println(*c)


	var a int = 10
	b := &a
	c := &b

	fmt.Println(*c)


	// var i int = 10
	// p1 := &i
	// p2 := &p1

	// fmt.Println(*p2, *p1)

	// var a int = 20
	// var p *int = &a + 10

	// fmt.Println(*p)
}
