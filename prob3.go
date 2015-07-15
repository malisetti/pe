package main

/**
Largest prime factor
Problem 3

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
**/

import (
	"fmt";
	"math"
)

func main() {
	largest_factor := 2
	number := 600851475143
	start := 2
	limit := int(math.Ceil(math.Sqrt(float64(number))))

	for start <= limit {
		for number % start == 0 {
			number = number / start
			if start >= largest_factor {
				largest_factor = start
			}
		}

		start++
	}

	fmt.Printf("Largest prime factor is %d \n", largest_factor)
}
