package main

import (
	"fmt"
	"functional/programming/util"
)

// * การประกาศ function type
type MathFunc func(int, int) int

func main() {

	// * pure function
	number1, number2 := util.PureFunction(10, 10)
	util.Add(number1, number2)
	// * pure function

	x := util.GetFunction()
	fmt.Println(x(5, 5))

	util.RunFunction(greet)

	newFunction := util.ReturnFunction()
	newFunction()

	util.PrintHelloFP(5)
}

func greet() {
	fmt.Println("Hello!")
}
