package main

import "fmt"

func main() {

	var a int = 30
	var b *int = &a
	var c *int = b
	var d **int = &b

	fmt.Println(*c, *b, a, **d)
}
