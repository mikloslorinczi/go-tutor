package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Add(a, b int) int {
	return a + b
}

func Incrase(a int) {
	fmt.Println("Incrase got argument: ", a)
	a++
}

func GetFileContent(filePath string) (content string, err error) {
	var (
		file      *os.File
		fileBytes []byte
	)

	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	fileBytes, err = ioutil.ReadAll(file)

	content = string(fileBytes)

	return
}

func main() {
	fmt.Println(Add(12, 33))

	content, err := GetFileContent("test.txt")
	if err != nil {
		fmt.Println("An error occured while getting the content of the file")
		panic(err)
	}
	fmt.Println(content)

	var I int

	I = 11

	fmt.Println("Value of I: ", I)

	Incrase(I)

	fmt.Println("Value of I after Incrase(I): ", I)
}
