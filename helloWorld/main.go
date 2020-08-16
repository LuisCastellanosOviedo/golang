package main

import "fmt"

var yy = 88

//declared z as string it means cant assign a int value ! beacuse is STRING !!
var z = "this is a string"


func main() {

	//short declaration operator
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)

	var z = "this is a string"
	fmt.Println(z)

	fmt.Println(yy, "<-- yy var variable , outside the function")
	foo()

	//z = 88 this is a string !!!!!! it wil throw an error

}

func foo() {
	fmt.Println("again yy from foo --->", yy)

}
