package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"os"

	"github.com/PuerkitoBio/goquery"
)

const (
	BASE_URL = "https://projecteuler.net/problem="
)

var wg sync.WaitGroup

func main() {
	for i := 1; i <= 523; i++ {
		wg.Add(1)
		go fetch(i)
		i++
	}
	wg.Wait()
}

func fetch(i int) {
	num := strconv.Itoa(i)
	url := BASE_URL + num
	defer wg.Done()
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("http transport error is:", err)
	} else {
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println("read error is:", err)
		} else {
			var title string = ""
			var problemContent string = ""
			doc.Find("#content h2").Each(func(i int, s *goquery.Selection) {
		    	title = s.Text()
		  	})

		  	doc.Find("#content .problem_content").Each(func(i int, s *goquery.Selection) {
		    	problemContent = s.Text()
		  	})

		  	//write to file
		  	path, err := os.Getwd()
		  	if err != nil {
		  		panic(err)
		  	}
		  	
		  	filepath := path + string(byte(os.PathSeparator)) + "prob" + num + ".go"

		  	f, err := os.Create(filepath)
		  	if err != nil {
		  		panic(err)
		  	}
	  	    defer f.Close()

	  	    _, _ = f.WriteString(title + "\n" + problemContent)
  	        f.Sync()
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println(string(body))
		}
	}
}
