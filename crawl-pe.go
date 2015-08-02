package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

const (
	BASE_URL = "https://projecteuler.net/problem="
)

var wg sync.WaitGroup

func main() {
	for i := 1; i < 2; i++ {
		url := BASE_URL + strconv.Itoa(i)
		wg.Add(1)
		go fetch(url)
	}
	wg.Wait()
}

func fetch(url string) {
	defer wg.Done()
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("http transport error is:", err)
	} else {
		body, err := ioutil.ReadAll(resp.Body)

		fmt.Println("read error is:", err)

		fmt.Println(string(body))
	}
}
