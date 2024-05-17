package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println(a + b)
	c := add(a, b)
	fmt.Println("Hello, World!!", c)
}

func add(a int, b int) int {
	return a + b
}
