package main

import (
	"fmt"
)

const (
	MOD = 10000000000
)

/**
(a+b)%c = (a%c + b%c)%c
(a*b)%c = (a%c * b%c)%c
**/
func main() {
	var temp int
	result := 0
	for i := 1; i < 1001; i++ {
		temp = i
		for j := 1; j < i; j++ {
			temp *= i
			temp %= MOD
		}
		result += temp
		result %= MOD
	}

	fmt.Printf("%d", result)
}
