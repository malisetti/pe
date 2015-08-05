package main

import (
	//"fmt"
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
	for i := 1; i < 523; i++ {
		wg.Add(1)
		go fetch(i)
	}
	wg.Wait()
}

func fetch(i int) {
	num := strconv.Itoa(i)
	url := BASE_URL + num
	defer wg.Done()
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != http.StatusOK {
		panic(err)
	} else {
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			panic(err)
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
		  	
		  	pathSeparator := string(byte(os.PathSeparator))
		  	filepath := path + pathSeparator + pathSeparator + "prob" + num + ".go"

		  	f, err := os.Create(filepath)
		  	if err != nil {
		  		panic(err)
		  	} else {
		  		println(url)
		  	}
	  	    defer f.Close()

	  	    _, _ = f.WriteString("package main \n\n\n/**\n" + url + "\n\n" + title + "\n" + problemContent + "**/\n")
  	        f.Sync()
		}
	}
}
