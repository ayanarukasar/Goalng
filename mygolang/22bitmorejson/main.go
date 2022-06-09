package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //alias
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` //remove d password from response data
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON in golang")
	// EncodeJson()
	DecodeJson()
}
func EncodeJson() {
	lcoCourses := []course{
		{"React", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Golang", 0, "LearnCodeOnline.in", "@12ert", nil},
		{"Angular", 799, "LearnCodeOnline.in", "root", []string{"full-stack", "js"}},
	}

	//package this data as json data

	// finalJson, err := json.Marshal(lcoCourses)

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)

	}
	fmt.Printf("%s\n", finalJson)
}

//Decoding

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React",
        "Price": 299,
        "website": "LearnCodeOnline.in",
        "tags": ["web-dev","js"]
    }
	`)

	var lcoCourses course

	chcekValid := json.Valid(jsonDataFromWeb)

	if chcekValid {
		fmt.Println("Json was valid!")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses) //interface is another name for struct
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//Some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData) //interface is another name for struct
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is : %T\n", key, value, value)
	}
}
