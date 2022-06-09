package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URL's in golang")
	fmt.Println(myurl)

	//Parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("qparam is :", val)
	}

	//creating a url

	partsofUrl := &url.URL{
		Scheme:   "https",
		Host:     "www.slideshare.net",
		Path:     "/AyanaRukasar",
		RawQuery: "/edit_my_uploads",
	}
	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)
}

//ctrl+hit on url to open url
