package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["Go"] = "Golang"
	languages["JS"] = "Javascript"
	languages["React"] = "Java"
	languages["Angular"] = "Python"

	fmt.Println("List of All LANGUAGES : ", languages)
	fmt.Println("JS shorts for : ", languages["JS"])

	delete(languages, "React")
	fmt.Println("List of All LANGUAGES after deletion : ", languages)

	//loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
