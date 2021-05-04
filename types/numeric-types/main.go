package main

import "fmt"

func main() {
	n := 42
	fmt.Printf("%v, %T\n", n, n)

	//int8    -128 ~ 127
	//int16   -32768 ~ 32767
	//int32   -2147482648 ~ 2147482647
	//int64   -9223372036854775808 ~ 9223372036854775807

	a := 10             //1010
	b := 3              //0011
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  //1001
	fmt.Println(a &^ b) //0100
	// values
	// 2
	// 11
	// 9
	// 8

	k := 8              // 2^3
	fmt.Println(k << 3) // 2^3 * 2^3 = 2^6 ;8 三次方 = 64
	fmt.Println(k >> 3) // 2^3 / 2^3 = 2^0 ;8 開根號 3 = 1
	// values
	// 64
	// 1
}
