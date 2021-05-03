package main

import (
	"fmt"
)

var r int = 44

//var group make code clear
var (
	a string = "a"
	b string = "b"
)

func main() {
	var i int
	i = 42
	i = 1000
	// var j int
	var j int = 2000
	// var k int
	k := 20.00
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	// value = 20, float6411 <nil>
	fmt.Println(fmt.Printf("%v, %T", k, k))
	// value = 44, int
	fmt.Printf("%v, %T", r, r)
	fmt.Println("")
	// a b
	fmt.Println(a, b)
}
