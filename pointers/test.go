package main

import "fmt"

func main() {

	var a int = 20
	var p *int = &a + 10

	fmt.Println(*p)
}
