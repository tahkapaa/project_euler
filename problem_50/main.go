package main

import "fmt"

/*
The prime 41, can be written as the sum of six consecutive primes:
41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.
The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.
Which prime, below one-million, can be written as the sum of the most consecutive primes?
*/

type primes []int
type solutions []primes

func (p primes) sum() int {
	sum := 0
	for _, val := range p {
		sum += val
	}
	return sum
}

func main() {
	p := getPrimesUnder(1000)
	fmt.Println(p)

	for i := 0; i < len(p); i++ {
		fmt.Println("Finding target: ", p[i])
		var s solutions
		ok := getSum(p[:i], primes{}, p[i], &s)
		fmt.Printf("Finding: %d, Success: %v, Count(Solutions): %v\n", p[i], ok, len(s))
		printLongestSolution(s)
	}
}
func printLongestSolution(s solutions) {
	var maxPrimes primes
	for _, p := range s {
		if len(p) > len(maxPrimes) {
			maxPrimes = p
		}
	}
	fmt.Printf("Max primes length: %d, Example: %v\n", len(maxPrimes), maxPrimes)
}

func getPrimesUnder(n int) primes {
	var p primes
	for i := 0; i < n; i++ {
		if isPrime(i) {
			p = append(p, i)
		}
	}
	return p
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getSum(p primes, sumParts primes, target int, sol *solutions) bool {
	for i := len(p) - 1; i >= 0; i-- {
		if p[i]+sumParts.sum() < target {
			newParts := append(sumParts, p[i])
			if i > 0 {
				getSum(p[:i], newParts, target, sol)
			}
		} else if p[i]+sumParts.sum() == target {
			// fmt.Println("***** FOUND SUM *****")
			newSolution := append(sumParts, p[i])
			*sol = append(*sol, newSolution)
			return true
		}
	}
	return len(*sol) > 0
}
