package main

import "fmt"

func main() {
	fmt.Println("Methods in golang")

	Ayana := User{"AYANA", "ayana.rukasar@capgemini.com", true, 22}

	fmt.Println(Ayana)
	fmt.Printf("Ayana details are : %+v\n", Ayana) //%+v for entire structure
	fmt.Printf("Name is %v and email is %v", Ayana.Name, Ayana.Email)
	Ayana.GetStatus()
	Ayana.NewMail()
	fmt.Printf("Name is %v and email is %v", Ayana.Name, Ayana.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("\nIs user is active : ", u.Status)
}
func (u User) NewMail() {
	u.Email = "ayana.rukasar@hpe.com"
	fmt.Println("\nEmail of this user is : ", u.Email)
}
