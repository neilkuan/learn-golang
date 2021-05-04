package main

import "fmt"

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	k := 1 == 1
	r := 1 == 2
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", r, r)

	// final values
	// true, bool
	// true, bool
	// false, bool

	var j bool
	fmt.Printf("%v, %T\n", j, j)
	// in golang if you declaration a var but do not give values, the default value will be 0 and in boolean will been false.
	// false, bool
}
