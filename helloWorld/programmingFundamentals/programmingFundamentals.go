package main

import (
	"fmt"
	"runtime"
)

var x bool

const a2 int = 22
const b2 = 33
const c2 = "constant"

const (
	a3 = 44
	b3 = 55
)

const (
	a4 = iota
	b4
	c4
)

func main() {

	fmt.Println("***************** BOOL *****************")
	// boolean
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 44
	b := 33
	// a == b :
	fmt.Println("a is equal to b ???")
	fmt.Println(a == b)

	fmt.Println("***************** NUMERIC  *****************")

	a1 := 43
	b1 := 43.3424

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", b1)

	fmt.Println("***************** PRINT RUNTIME  *****************")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println("***************** STRINGS   *****************")
	s := "helllos strings "
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println("***************** CONST    *****************")
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)
	fmt.Println(a3)
	fmt.Println(b3)

	const c3 = 66
	fmt.Println(c3)

	fmt.Println("***************** CONST  IOTA  *****************")
	fmt.Println(a4)
	fmt.Println(b4)
	fmt.Println(c4)

}
