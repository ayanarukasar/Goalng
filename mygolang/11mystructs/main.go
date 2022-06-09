package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	Ayana := User{"AYANA", "ayana.rukasar@capgemini.com", true, 22}

	fmt.Println(Ayana)
	fmt.Printf("Ayana details are : %+v\n", Ayana) //%+v for entire structure
	fmt.Printf("Name is %v and email is %v", Ayana.Name, Ayana.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
