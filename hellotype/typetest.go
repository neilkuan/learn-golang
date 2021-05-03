package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	hell := rand.Intn(1)
	hellt := reflect.TypeOf(hell).Kind()
	fmt.Println(hellt)
	if hellt == reflect.Int {
		fmt.Println("hell is: ", reflect.Int)
	}
}
