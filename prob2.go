package main

/**
Even Fibonacci numbers
Problem 2

Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
**/

import (
	"fmt"
)
const (
	LIMIT = 4000000
)

func main() {
	sum := 0
	c := fib_generator()
	for {
		fib := <- c
		if 0 == fib % 2 {
			sum += fib
		} else if fib >= LIMIT {
			break
		}
	}

	fmt.Printf("the sum of the even-valued terms is %d\n", sum)
}

func fib_generator() chan int {
  c := make(chan int)

  go func() { 
    for i, j := 0, 1; j <= LIMIT; i, j = i+j,i {
        c <- i
    }
    defer close(c)
  }()

  return c
}