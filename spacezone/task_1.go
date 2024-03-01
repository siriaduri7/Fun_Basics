package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	/// TODO 1: Save int in float and try

	var i int = 10
	f := float32(i)
	fmt.Println(f)

	// TODO 2: Save float in int and try

	x := 20.5
	y := int(x)
	fmt.Println(y)

	// TODO 3: Save int in string and try

	var a int = 4
	var b string = fmt.Sprintf("%d", a)
	fmt.Println(b)

	//2nd way to convert int to string

	b = strconv.Itoa(a)
	fmt.Println(b)

	// TODO 4: Save string in int and try

	var j string = "42"
	m, err := strconv.Atoi(j)

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(m)

	// TODO 5: Save float in string and try

	var l float64 = 20.4
	var k string = fmt.Sprintf("%f", l)
	fmt.Println(k)

	//TODO 6: Save string in float and try

	var z string = "40.5"
	o, err := strconv.ParseFloat(z, 64)

	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(o)

	// TODO 7: Save int in in64 and try,find out sizes of int and int64, other int types like int8, int16, int32
	//for memory optimization we use different kind of ints
	//Without these conversions, assigning x directly to y, z, w, or v would result in a compilation error because Go does not allow implicit conversions to the appropriate size and range required by each integer type between different numeric types.
	//maintaining type safety and preventing unintended data loss or overflow.

	var q int = 70
	var p int64 = int64(q)
	fmt.Println(p)

	//find out sizes of int and int64, other int types like int8, int16, int32

	fmt.Printf("size of int8: %d byte\n", unsafe.Sizeof(int8(2)))
	fmt.Printf("size of int64: %d bytes \n", unsafe.Sizeof(int64(0)))
	fmt.Printf("size of int32: %d bytes \n", unsafe.Sizeof(int32(0)))

	// TODO 8: Save int64 in int and try

	var g int64 = 700000000000000000
	var s int = int(g)
	fmt.Println(s)

	// TODO 9: Save float32 in float64 and try

	var t float32 = 77787.95678698 //rounding and approximation that occurs during the conversion and representation of the floating-point numbers. o/p 77787.953125
	var u float64 = float64(t)
	fmt.Println(u)

	// TODO 10: Check what is size of float32 and float64

	fmt.Printf("size of float64: %d bytes \n", unsafe.Sizeof(float64(0)))
	fmt.Printf("Size of float32: %d bytes \n", unsafe.Sizeof(float32(0)))

	// TODO 11: intialize an int or float with math operation

	e := 10 + 20
	c := 55.8 * 67.0

	// TODO 12: intialize an int or float with math operation between two variables like define i,j and then k=i*j
	fmt.Println(float64(e), c)
	r := float64(e) * c
	fmt.Println(r)

}
