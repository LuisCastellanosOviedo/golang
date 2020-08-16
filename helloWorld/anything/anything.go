package main

import "fmt"

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

}

func foo() {
	fmt.Println("im in foo")
}
