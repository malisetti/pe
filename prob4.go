package main

/**
Largest palindrome product
Problem 4

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
**/

func main() {
	largest_palindrome := 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			prod = i * j
			if IsPalindrome() && largest_palindrome < prod {
				largest_palindrome = prod
			}
		}
	}

	fmt.Printf("Largest palindrome %d ", largest_palindrome)
}

//http://stackoverflow.com/questions/18940140/how-to-check-if-number-is-bit-palindrome-or-not
func IsPalindrome(number int) bool {
	reversed := 0
	aux := number

    while (aux > 0) {
    	reversed = (reversed  << 1) | (aux & 1)
	    aux = aux >> 1
    }

    return reversed  == number;
}

