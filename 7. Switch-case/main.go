package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5

	switch i {
	case 1:
		fmt.Println("number is 1")

	case 5:
		fmt.Println("Number is 5")

	default:
		fmt.Println("You Look like a drug dealer")
	}

	switch time.Now().Weekday() {

	case time.Saturday, time.Sunday:
		fmt.Println("Shut the f*ck Up and Sleep")

	default:
		fmt.Println("Time to go to work you matrix enslaved brokies")
	}

	WhoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer", t)
		case string:
			fmt.Println("String")

		}
	}

	WhoAmI("golang")
	WhoAmI(3)
}
