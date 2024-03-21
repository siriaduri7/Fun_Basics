package main

import "fmt"

func main() {

	a := needA
	a()

	b := needB
	b()

	accessAB(needA, needB)
	a()
	b()

	// 	add(4, 5)
	// }

	//	func add(a, b int) {
	//		fmt.Println(a + b)
}

func needA() {
	fmt.Println("needA")
}

func needB() {
	fmt.Println("needB")
}

func accessAB(a func(), b func()) {

}
