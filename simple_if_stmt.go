package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var x int = 1

	if x < 5 {
		fmt.Println("bigger")
		x = x + 1

	}
}
