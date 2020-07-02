package main

import (
	"fmt"
)

func square(x int) int {
	return x * x
}

func cube(x int) int {
	return x * x * x
}

func paymentCalculator(dow int) func(x int) int {
	if dow >= 1 && dow <= 5 {
		return square
	}
	return cube
}

func calculatePayment(rate, dow int) int {
	return paymentCalculator(dow)(rate)
}

func x(f func(d int) int, g int) int {
	squared := f(g)
	return 2 + squared
}

func main() {
	fmt.Println(calculatePayment(3, 4))
	fmt.Println(x(square, 12))
}
