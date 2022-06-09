package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var cosmeticList = []string{"Lakme", "Mac", "Sugar", "Maybelline"}
	fmt.Printf("The type of COSMETIC LIST is %T\n", cosmeticList)

	//Method 1
	cosmeticList = append(cosmeticList, "Purple", "Swiss beauty")
	sort.Strings(cosmeticList) //sorting
	fmt.Println(cosmeticList)

	//Method 2

	cosmeticList = append(cosmeticList[1:3])

	fmt.Println(cosmeticList)

	//Another Method

	highscores := make([]int, 4)

	highscores[0] = 98
	highscores[1] = 95
	highscores[2] = 92
	highscores[3] = 88

	highscores = append(highscores, 87, 99, 83)

	fmt.Println("Toppers with :", highscores)
	fmt.Printf("Types is : %T\n", highscores)
	fmt.Println("Length is :", len(highscores))

	sort.Ints(highscores) // to sort d array
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	//How to remove a value from slice based on index

	var courses = []string{"react", "golang", "python", "angular", "graphql"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
