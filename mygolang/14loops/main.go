package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Loops in golang")

	days := []string{"Monday", "Sunday", "Friday", "Tuesday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	sort.Strings(days)
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index value is %v and day value is %v\n", index, day) //More used
	}

	rouguevalue := 1

	// for rouguevalue < 10 {
	// 	if rouguevalue == 5 {
	// 		break

	// 	}
	for rouguevalue < 10 {

		if rouguevalue == 2 {
			goto lco

		}

		if rouguevalue == 5 { //skips 5
			rouguevalue++
			continue

		}
		fmt.Println("Value is : ", rouguevalue)
		rouguevalue++
	}
lco:
	fmt.Println("Jumping to learn coding online")
}
