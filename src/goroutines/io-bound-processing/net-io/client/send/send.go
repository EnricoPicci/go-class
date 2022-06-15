package send

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sendRequest() string {
	defer func() {
		if r := recover(); r != nil {

			_, ok := r.(*url.Error)
			if ok {
				fmt.Println("You probably need to start the server.\n Error:\n", r)
				os.Exit(1)
			}
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	resp, err := http.Get("http://localhost:8080/dostuff")
	check(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	check(err)
	_body := string(body)
	return _body
}

func sendRequests(howMany int, res chan []string) {
	responses := primSendRequests(howMany)

	res <- responses
}
func primSendRequests(howMany int) []string {
	responses := make([]string, howMany)

	for i := 0; i < howMany; i++ {
		resp := sendRequest()
		responses[i] = resp
	}

	return responses
}

func split(howMany int, concurrent int) []int {
	split := make([]int, concurrent)
	for i := 0; i < concurrent-1; i++ {
		split[i] = howMany / concurrent
	}
	split[concurrent-1] = howMany - (concurrent-1)*(howMany/concurrent)
	return split
}

func CallServer(howManyTimes int, concurrent int) int {
	split := split(howManyTimes, concurrent)

	var requestsprocessed int = 0

	result := make(chan []string, concurrent)

	for _, s := range split {
		go sendRequests(s, result)
	}

	for i := 0; i < concurrent; i++ {
		requestsprocessed += len(<-result)
	}

	return requestsprocessed
	//fmt.Println("Program terminating")
}

func GetSleepTime() string {
	defer func() {
		if r := recover(); r != nil {

			_, ok := r.(*url.Error)
			if ok {
				fmt.Println("You probably need to start the server.\n Error:\n", r)
				os.Exit(1)
			}
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	resp, err := http.Get("http://localhost:8080/sleep")
	check(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	check(err)
	sleep := string(body)
	return sleep
}
