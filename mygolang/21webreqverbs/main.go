package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verbs in golang")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)

	}
	defer response.Body.Close()
	fmt.Println("Status code : ", response.StatusCode)
	fmt.Println("Content length is : ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

	//Another method
	// var responseString strings.Builder
	// content, _ := ioutil.ReadAll(response.Body)
	// byteCount, _ := responseString.Write(content)

	// fmt.Println("Bytecount is : ", byteCount)
	// fmt.Println(responseString.String())
}
func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"learnCodeOnline.in"
		}
	
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)

	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "Aamir")
	data.Add("lastname", "Waris")
	data.Add("email", "aamir.waris@infosys.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)

	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
