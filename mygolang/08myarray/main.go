package main

import "fmt"

func main() {

	fmt.Println("Welcome to array in golang")

	var cosmeticList [4]string

	cosmeticList[0] = "Sugar"
	cosmeticList[1] = "Lakme"
	cosmeticList[2] = "Mac"
	cosmeticList[3] = "Maybelline"

	fmt.Println("Cosmetic List is :", cosmeticList)
	fmt.Println("Cosmetic List is :", len(cosmeticList))

	var colorList = [3]string{"Red", "Pink", "Brown"}
	fmt.Println("Color List is :", colorList)
	fmt.Println("Color List is :", len(colorList))

}
