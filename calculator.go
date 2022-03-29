package main

import (
	"fmt"
	"math"
)

func add(a, b int) int {
	//Insert code here
	return a + b
}

func subtract(a, b int) int {
	//Insert code here
	return a - b
}

func multiply(a, b int) int {
	//Insert code here
	return a * b
}

func divide(a, b int) int {
	//Insert code here
	//consider for b = 0
	if b == 0 {
		fmt.Println("Zero Division Error")
		return int(math.Inf(a))
	}

	return a / b

}

func main() {
	var a, b int
	var process string
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&a)
	fmt.Println("Enter process: (add, sub, mul, div)")
	fmt.Scanln(&process)
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&b)

	//Insert code here
	if process == "add" {
		result := add(a, b)
		fmt.Println(result)
	} else if process == "sub" {
		result := subtract(a, b)
		fmt.Println(result)
	} else if process == "mul" {
		result := multiply(a, b)
		fmt.Println(result)
	} else if process == "div" {
		result := divide(a, b)
		fmt.Println(result)

	} else {
		fmt.Println("You did not enter the correct process! Please try again!")
	}
}
