package main

import (
	"fmt"
	"math/rand"
)

func main() {
	firstNames := [5]string{"Anand", "Shreya", "Deepa", "Rishya", "Chintu"}
	lastNames := [5]string{"Anand", "Rajagopal", "Narayanan", "Suresh", "Smith"}
	fullNames := make([]string, 5)
	for i := 0; i < 5; i++ {
		rand.Seed(int64(i))
		x := firstNames[rand.Intn(5)]
		y := lastNames[rand.Intn(5)]
		fullNames[i] = x + " " + y

	}
    fmt.Println(fullNames)
    fmt.Println("testing git")
}
