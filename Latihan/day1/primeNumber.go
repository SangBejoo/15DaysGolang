package main

import "fmt"

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

func printPrimeInRange(start, end int) {
	fmt.Printf("Prime %d is %d\n", start, end)
	count := 0

	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
			count++
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Printf("\nTotal: %d primes\n", count)
}

func main() {
	printPrimeInRange(1, 100)
}
