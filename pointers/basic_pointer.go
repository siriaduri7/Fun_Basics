// Write a Go program to swap two integers using pointers.
package main

import "fmt"

func main() {
	var a, b int = 20, 30
	fmt.Println("enter value  of  a", a, "enter value of b", b)

	swap(&a, &b)

	fmt.Println("after swapping")
	fmt.Println("enter a after swap", a)
	fmt.Println("enter b after swap", b)

	var num int = 20

	fmt.Println("before increment ", num)

	incrementbyone(&num)

	fmt.Println("after increment", num)

	var doub int = 30

	doubleval(&doub)
	fmt.Println("++++++++++++++++++++++++++++++")

	fmt.Println(doub)
	fmt.Println("++++++++++++++++++++++++++++++")
	//Develop a simple Go program that allocates memory for an integer using the new keyword, sets its value, and then prints the value and the memory address.

	var o *int = new(int)
	*o = 40

	fmt.Println(*o)
	fmt.Println(o)

}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

//Create a function in Go that takes an integer pointer as a parameter and increments the value it points to by 1.

func incrementbyone(ptr *int) {
	*ptr = *ptr + 1
}

// Write a program that demonstrates passing a pointer to a function as an argument.
func doubleval(ptr1 *int) {
	*ptr1 = *ptr1 * 2
}
