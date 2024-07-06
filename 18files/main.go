package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang");
	content:= "This needs to go in a file - learn code"

	file, err := os.Create("./mylcogofile.txt")

	checkNilErr(err)

	length,err := io.WriteString(file, content)

	checkNilErr(err)
	
	fmt.Println("Length is: ",length)
	defer file.Close()
	readFile("./mylcogofile.txt")

}


func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	
	checkNilErr(err)
	
	fmt.Println("Text data inside the fileis \n",databyte)
}


func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}