package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// TODO 13: try interfence type of variable (i,j:=10,11) also replace 10, 11 with some operations of variables

	i, j := 10, 11
	// Print the values and types of variables
	fmt.Printf("i : %d, type of variable %T \n", i, i)
	fmt.Printf("j : %d, type of variable %T \n", j, j)

	x, y := 5, 6
	fmt.Println(x, y)

	// TODO 14: define a pointer of pointer of pointer int variable and try to print the value and address of the variable

	var a int = 20

	var b *int = &a

	var c *int = &b

	var l ***int = &c

	fmt.Println(***l)
	fmt.Println(&a, &b, &c)

	// TODO 15: Define a new pointer q and assign the address of the pointer p to it (p is pointer to int i variable).
	//			Print the type of the new pointer q.
	//			Change the value of the variable i indirectly by changing the value at the address stored in the pointer q.
	//			Print the value of the variable i by dereferencing the pointer p

	var o int = 100 //i place ising o

	var p *int = &o

	var q **int = &p

	fmt.Printf("type of pointer: %T\n", **q)

	**q = 60

	fmt.Println(o)

	// TODO 16: //Type conversion of pointer to int to pointer to float - is this possible , explain

	var r int = 20
	f := &r

	fmt.Println(*f)

	//converting pointer to int

	intval := uintptr(unsafe.Pointer(f))

	//convert float to pointer

	// Attempt to convert int to pointer to float32 (unsafe)
	// This is highly unsafe and can lead to undefined behavior
	floatPtr := (*float32)(unsafe.Pointer(f))

	//uintptr is often used in conjunction with the unsafe package to perform low-level operations, such as pointer arithmetic and memory manipulation, that are not safe within the bounds of Go's type system.

	fmt.Println("Pointer representation of the float32:", floatPtr)
	fmt.Println("Pointer representation of the float32:", intval)

}
