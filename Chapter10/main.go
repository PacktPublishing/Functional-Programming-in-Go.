package main

import (
	"fmt"
	"net/http"

	"github.com/PacktPublishing/Chapter10/pkg"
)

type ResponseFunc func(*http.Response)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	output := pkg.ConcurrentFilter(ints, func(i int) bool { return i%2 == 0 }, 11)
	fmt.Printf("%v\n", output)
}

/*
func main() {
	success := func(response *http.Response) {
		fmt.Println("success")
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", string(content))
	}
	failure := func(response *http.Response) {
		fmt.Printf("something went wrong, received: %d\n", response.StatusCode)
	}
	go getURL("https://news.ycombinator.com", success, failure)
	go getURL("https://news.ycombinator.com/ThisPageDoesNotExist", success, failure)
	done := make(chan bool)
	<-done // keep main alive
}

func getURL(url string, onSuccess, onFailure ResponseFunc) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		onSuccess(resp)
	} else {
		onFailure(resp)
	}
}
*/