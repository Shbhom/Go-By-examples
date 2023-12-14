package main

import (
	"fmt"
	"time"
)

func Main5() {
	i := 1

	//default is optional in GO

	switch i {
	case 1, 2:
		fmt.Println("either 1 or 2")
	case 3, 4:
		fmt.Println("either 3 or 4")
	case 5, 6:
		fmt.Println("either 5 or 6")
	}

	//switch without any expression is if-else

	q := time.Now().Weekday()
	switch {
	case q == time.Saturday || q == time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("neither bool nor int ", t)
		}
	}

	whatAmI("hi")

}
