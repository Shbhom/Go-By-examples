package main

import "fmt"

func Main4() {
	if 7 > 2 || 6 < 4 {
		fmt.Println("Kanye Best")
	}

	if 6 < 4 && 7 > 2 { //short circuiting: since first condition fails and we are using && so won't be checking the next condition
		fmt.Println("the greatest: taylor Swift")
	}

	if num := 32; num > 31 {
		fmt.Println("Kanye Best")
	} else if num < 31 {
		fmt.Println("taylor Trash")
	} else if num == 32 {
		fmt.Println("both best")
	}
}
