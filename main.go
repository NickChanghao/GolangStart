package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	fmt.Println("The result =", Sum(a, b))
}

func Sum(a int, b int) int {
	return a + b
}
