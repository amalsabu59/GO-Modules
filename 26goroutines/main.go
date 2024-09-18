package main

import (
	"fmt"
	"net/http"
	"sync"
)
	var wg sync.WaitGroup
func main() {

	websitelist := []string {
		"https://google.com",
		"https://youtube.com",
		"https://fb.com",
		
	}

	for _,web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS error in endpoint: ", endpoint)
	} else {
		fmt.Printf("%d status code for %s \n",res.StatusCode, endpoint)
	}
defer wg.Done()
}