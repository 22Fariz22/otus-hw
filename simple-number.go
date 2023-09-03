package main

import "fmt"

func main() {
	fmt.Println(isPrime(7))
}

func isPrime(n int) bool {
	if n == 0 {
		return false
	}

	for i := 1; i < n+1; i++ {
		if n%2 == 0 {
			return false
		}
	}

	return true
}
