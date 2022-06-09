package main

import "fmt"

func main() {
	fmt.Println("If Else in Golang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
		// } else {
		// 	result = "something else"

	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Result is exactly 10"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")

	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

}
