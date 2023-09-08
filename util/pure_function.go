package util

import (
	"fmt"
)

type Test func(int, int) int

// * pure function

func PureFunction(value1 int, value2 int) (int, int) {
	return value1, value2
}

func Add(value1 int, value2 int) int {
	sum := value1 + value2
	fmt.Println(sum)
	return sum
}

// * pure function

func GetFunction() Test {
	return func(a int, b int) int {
		return a + b
	}
}

func RunFunction(fn func()) {
	fn()
}

func ReturnFunction() func() {
	return func() {
		fmt.Println("Returned function.")
	}
}

func PrintHelloFP(times int) {
	for i := 1; i <= times; i++ {
		fmt.Printf("Hello FP {ครั้งที่ %d}\n", i)
	}
}
