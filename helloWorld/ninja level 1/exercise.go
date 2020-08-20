package main

import "fmt"

var a int
var b string
var c bool

var a1 int
var b1 string
var c1 bool

type hotdog int // int is the underlying type

var newtypehahaha hotdog

var x5 int

func main() {

	// exercise 1
	x := 42
	y := "James bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)

	// exercise 2 print zero values
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a, b, c)

	// exercise 3 print types

	a1 = 42
	b1 = "James bond"
	c1 = true

	s := fmt.Sprintf("%v\t%v\t%v\t", a1, b1, c1)
	fmt.Println(s)

	//exercise 4 new type

	fmt.Println(newtypehahaha)
	fmt.Printf("%T\n", newtypehahaha)

	newtypehahaha = 42
	fmt.Println(newtypehahaha)

	//exercise 5

	x5 = int(x) // not x5 := int(x)
	fmt.Println(x5)
	fmt.Printf("%T\n", x5)

}
