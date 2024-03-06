package main

import "fmt"

func main() {

	var a int = 30
	var b *int = &a
	var c **int = &b
	var d *int = b

	fmt.Println(**c, *b, a, *d)
}

//***d is the indirect pointer to pointer to an integer that is pointing to the &c from  to pointer int has b value, in b it has address &a in that we have value of 30 so ***d is indirectly pointing to the value at location.
//*c is pointer int  variable declaration

// in the above program we can also write like if it is indirect d==c==b==a==30 you need to use starts ***d , **c, *b so if you want to use d==b than d *int =b will work based on how indirectly u pointing then that many stars
