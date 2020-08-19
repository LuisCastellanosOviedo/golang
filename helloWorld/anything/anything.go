package main

import "fmt"

var y = 24
type hotdog int

var b hotdog

func main() {
	fmt.Println("helloo func")
	foo()

	n, _ := fmt.Println("Hello, playground 	", true, 42)
	fmt.Println(n)

	// variable declarations
	x := 42
	fmt.Println(x)

	x = 99

	fmt.Println(x)

	// a statemnet
	y := 100 + 24
	fmt.Println(y)

	fmt.Println("fmt lib examples -->")
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%T\t", y, y, y)

	// new types 
	fmt.Println("+++++++++++++++++++++++++++")
	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// y = b thorw error cause are different types 
	fmt.Println("CONVERTIONNNN NOT CASTING ")
	fmt.Println(int(b))

}

func foo() {
	fmt.Println("im in foo")
}
