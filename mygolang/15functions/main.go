package main

import "fmt"

func main() {

	fmt.Println("Functions in Golang")
	greeter()
	greeterTwo()

	result, Message := adder(3, 4)
	fmt.Println("Result is : ", result)
	fmt.Println("Message is : ", Message)

	proResult := proAdder(4, 5, 1, 1, 2, 3, 5, 6, 8, 6, 6)
	fmt.Println("Pro Result is : ", proResult)
}

func adder(valueOne int, valueTwo int) (int, string) {
	return valueOne + valueOne, "Ayana W."
}
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Namastey from golang")
}

func greeterTwo() {
	fmt.Println("Another Method")
}
