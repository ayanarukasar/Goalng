package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the study of time in golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//Date
	createdDate := time.Date(2021, time.March, 15, 11, 0, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday")) //by default
}
