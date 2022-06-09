package main

import "fmt"

func main() {
	defer fmt.Println("Welcome")
	defer fmt.Println("to")
	defer fmt.Println("Golang")

	fmt.Println("Hello")

	mydefer() //function call will execute first

	//defer is like reverse

}
func mydefer() {
	for i := 0; i < 5; i++ {

		defer fmt.Println(i)

	}
}
