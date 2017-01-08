package main

import "fmt"

/*
Even Fibonacci numbers
Problem 2
Each new term in the Fibonacci sequence is generated by adding the previous
two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.
*/

func main() {
	var sum int
	for _, value := range getFibonacciNumbersBelow(4000000) {
		if value%2 == 0 {
			sum += value
		}
	}
	fmt.Println("Sum: ", sum)
}

func getFibonacciNumbersBelow(max int, previousNumbers ...int) []int {
	if len(previousNumbers) < 2 {
		return getFibonacciNumbersBelow(max, []int{1, 2}...)
	}
	next := previousNumbers[len(previousNumbers)-1] + previousNumbers[len(previousNumbers)-2]
	if next > max {
		return previousNumbers
	}
	return getFibonacciNumbersBelow(max, append(previousNumbers, next)...)
}