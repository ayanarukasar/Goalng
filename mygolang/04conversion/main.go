package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our cosmetic app")
	fmt.Println("Enter the ratings for our cosmetics")

	reader := bufio.NewReader(os.Stdin)

	//comma ok || err ok
	//Try input , catch err or _

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the ratings :", input)
	fmt.Printf("Type of the cosmetics is %T\n", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your ratings: ", numRating+1)
	}
}
