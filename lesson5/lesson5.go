package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func reverse(n, k int) int {
	if n < 10 {
		return n + k
	}
	rem := n % 10
	n = n / 10
	k = k * 10 + rem * 10
	return reverse(n, k)

}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(reverse(12345, 0))
}
