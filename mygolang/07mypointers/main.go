package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int  //creating a pointer
	// fmt.Println("Value of this pointer is : ", ptr)

	mynumber := 23

	var ptr = &mynumber //creating a pointer + referencing a pointer

	fmt.Println("Value of actual pointer is : ", ptr)
	fmt.Println("Value of actual pointer is : ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is : ", mynumber)

}
