package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")

	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	// if err != nil {
	// 	panic(err) //shutdown the execution of the program will show u what err we are facing
	// }

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }

	checkNilErr(err)

	fmt.Println("Length is : ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

//Reading

func readFile(filename string) {
	datbyte, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }

	checkNilErr(err)

	// fmt.Println("Text data inside the file is :\n", datbyte)
	fmt.Println("Text data inside the file is :\n", string(datbyte))
}

//Instead of using if error so many times use

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}

}
