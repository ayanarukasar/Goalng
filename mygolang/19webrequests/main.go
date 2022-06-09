package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.linkedin.com/in/ayana-rukasar-28a9851a3/?originalSubdomain=in"

func main() {
	fmt.Println("LCO web requests")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T", response)

	defer response.Body.Close() //Callers responsibility to close the connection

	//Reading file using ioutil

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
