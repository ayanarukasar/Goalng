package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Math in Golang")

	// var mynumberOne int = 2
	// var mynumberTwo float32= 4

	// fmt.Println("The sum of numbers is :",mynumberOne+int(mynumberTwo))

	//random number from math
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	//random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
