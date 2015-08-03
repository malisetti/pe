package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	BASE_URL = "https://projecteuler.net/problem="
)

var wg sync.WaitGroup

func main() {
	done := make(chan bool)
	d := true
	i := 1
	for d != false {
		url := BASE_URL + strconv.Itoa(i)
		wg.Add(1)
		go fetch(url, done)
		i++
		d = <- done
	}
	wg.Wait()
}

func fetch(url string, done chan bool) {
	done <- true
	defer wg.Done()
	resp, err := http.Get(url)

	if err != nil {
		done <- false
		fmt.Println("http transport error is:", err)
	} else {
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println("read error is:", err)
		} else {
			doc.Find("#content h2").Each(func(i int, s *goquery.Selection) {
		    	title := s.Text()
		    	fmt.Printf("%s Title - %s\n", url, title)
		  	})
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println(string(body))
		}
	}
}
