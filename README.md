# This repo just for learn golang for beginner learn-golang

```bash
# test.go 
# `package`
# go run: cannot run non-main package
# func format need complete with "ansi c" .


go run test.go

---
Hello world
```



if else to check object Type.
```go
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
```