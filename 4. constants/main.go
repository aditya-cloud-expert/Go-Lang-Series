package main

import "fmt"

func main() {
	const name = "Aditya"
	fmt.Println(name)
	const number = 1
	fmt.Println(number)

	const (
		method = "POST"
		URL    = "localhost"
	)

	fmt.Println(method)
}
