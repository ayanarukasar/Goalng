package main

import "fmt"

const LoginToken string = "gjsdhfsjh" //public

func main() {
	var username string = "Ayana"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style
	numberOfUser := 300
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T\n", LoginToken)
}
