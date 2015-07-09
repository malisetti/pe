package main
/**
Multiples of 3 and 5
Problem 1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
**/
import(
	"fmt"
)
const LIMIT = 1000

func main() {
	sum := 0
	for i := 3; i < LIMIT; i++ {
		if 0 == i % 3 || 0 == i % 5 {
			sum += i
		}
}	

fmt.Printf("sum of all the multiples of 3 or 5 below %d is : %d\n", LIMIT, sum)
}