package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const BASE_URL = "https://projecteuler.net/problem="

func fetch(problems <-chan int, done <-chan bool) {
	for {
		select {
		case i := <-problems:
			num := strconv.Itoa(i)
			url := BASE_URL + num

			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			} else {
				defer resp.Body.Close()
				doc, err := goquery.NewDocumentFromReader(resp.Body)

				if err == nil {
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
						log.Fatal(err)
					}

					filepath := path + pathSeparator + "prob" + padLeft(num, "0", 3) + ".go"

					f, err := os.Create(filepath)
					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					f.WriteString("package main \n\n\n/**\n" + url + "\n\n" + title + "\n" + problemContent + "**/\n")
					f.Sync()
				}
			}
			wg.Done()

			fmt.Printf("Finished processing problem #%d", i)
		case <-done:
			return
		}
	}
}

func padLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

var wg sync.WaitGroup
var pathSeparator string

func main() {
	pathSeparator = string(byte(os.PathSeparator))

	problems := make(chan int)
	done := make(chan bool)

	defer close(problems)
	defer close(done)

	for i := 0; i < 3; i++ {
		go fetch(problems, done)
	}

	for i := 1; i < 556; i++ {
		wg.Add(1)

		problems <- i
	}

	wg.Wait()
	done <- true
}
