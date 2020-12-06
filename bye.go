package main

import "fmt"

func bye() {
	fmt.Println("Byeeee!")
}

func hey() {
	fmt.Println("Hey!")
}

func main() {
	fmt.Println("Hello!")
	bye()
	hey()
}
