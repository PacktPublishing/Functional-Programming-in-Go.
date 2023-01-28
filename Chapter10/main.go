package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PacktPublishing/Chapter10/pkg"
)

type ResponseFunc func(*http.Response)

func main() {
	/*
		generated := pkg.Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

		filtered := pkg.ConcurrentFilter(generated, func(i int) bool { return i%2 == 0 }, 3)
		fmt.Printf("%v\n", output)
		mapped = pkg.ConcurrentMap(filtered, func(i int) int { return i * 2 }, 2)
		fmt.Printf("%v\n", output)
		collected := pkg.Collect(mapped)
		fmt.Printf("%v\n", collected)

		result := pkg.ConcurrentFMap(output, func(i int) string { return "number: " + strconv.Itoa(i) }, 2)
		fmt.Printf("%v\n", result)
	*/

	fmt.Printf("\n\n\n===pipelines===\n")
	generated := pkg.Generator(1, 2, 3, 4)
	filtered := pkg.FilterNode(generated, func(i int) bool { return i%2 == 0 })
	mapped := pkg.MapNode(filtered, func(i int) int { return i * 2 })
	collected := pkg.Collector(mapped)
	fmt.Printf("%v\n", collected)

	fmt.Println("chaining")
	out := pkg.ChainPipes(pkg.Generator(1, 2, 3, 4),
		pkg.CurriedFilterNode(func(i int) bool { return i%2 == 0 }),
		pkg.CurriedMapNode(func(i int) int { return i * i }))

	fmt.Println(out)

	fmt.Println("chain pipes 2")
	out2 := pkg.ChainPipes2[string](pkg.CurriedCat("./main.go"),
		pkg.CurriedFilterNode(func(s string) bool { return strings.Contains(s, "func") }),
		pkg.CurriedMapNode(func(i string) string { return "line contains func: " + i }))

	fmt.Printf("%v\n", out2)
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
