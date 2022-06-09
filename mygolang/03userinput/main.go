package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to the User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the ratings for our cosmetics : ")

	//comma ok || err ok
	//Try input , catch err or _

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the ratings :", input)
	fmt.Printf("Type of the cosmetics is %T", input)
}
