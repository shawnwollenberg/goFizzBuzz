package Fizzbuzz

import (
	"fmt"
	"strconv"
)

// --- Directions
// Write a program that console logs the numbers
// from 1 to n. But for multiples of three print
// “fizz” instead of the number and for the multiples
// of five print “buzz”. For numbers which are multiples
// of both three and five print “fizzbuzz”.
// --- Example
//   fizzBuzz(5);
//   1
//   2
//   fizz
//   4
//   buzz

func Fizzbuzz(n int) (result []string) {

	for i := 1; i <= n; i++ {
		x := ""
		if i%3 == 0 {
			x += "fizz"
		}
		if i%5 == 0 {
			x += "buzz"
		}
		if x == "" {
			x = strconv.Itoa(i)
		}
		result = append(result, x)
	}
	fmt.Println(result)
	return result
}
