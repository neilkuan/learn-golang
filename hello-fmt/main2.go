package main

import (
	"fmt"
	"strconv"
)

func main() {
	// convert tetsing int to float32
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	// final values
	// 42, int
	//42, float32

	// convert tetsing float32 to int.
	var k float32 = 42.5
	fmt.Printf("%v, %T\n", k, k)

	var l int
	l = int(k)
	fmt.Printf("%v, %T\n", l, l)

	// final values
	// 42.5, float32
	// 42, int

	// convert tetsing int to string.
	var a int = 42
	fmt.Printf("%v, %T\n", a, a)

	var b string
	b = string(a)
	fmt.Printf("%v, %T\n", b, b)

	// final values
	// 42, int
	// *, string <- you must to use "strconv" package to help convert int to string ...
	// convert tetsing int to string.
	var z int = 42
	fmt.Printf("%v, %T\n", z, z)

	var x string
	x = strconv.Itoa(z)
	fmt.Printf("%v, %T\n", x, x)
	// final values
	// 42, int
	// 42, string
}
