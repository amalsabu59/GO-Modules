package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://bonkmark.com/learn?course=react&code=awsqwe"
func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	result,_ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Printf("The Type of query parms are: %T\n",qparams)

	fmt.Println(qparams["course"])
	for _, value := range qparams{
		fmt.Println(value)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:"bonkmark.co",
		Path:"/learn",
		RawQuery:"course=react&code=awsqwe",
	}

	anotherURL := partsOfUrl.String();
	fmt.Println("anotherUrl",anotherURL)
}