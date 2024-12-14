package main

import "fmt"

// only construct in go for looping, there is no while loop in go

func main() {
	i := 0

	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for {
		i *= i

		println(i)
	}
}
