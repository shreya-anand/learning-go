package main

import (
	"fmt"
	"math/rand"
)

type Person struct {
	firstName string
	lastName string
}

func main() {
	firstNames := [5]string{"Anand", "Shreya", "Deepa", "Rishya", "Chintu"}
	lastNames := [5]string{"Anand", "Rajagopal", "Narayanan", "Suresh", "Smith"}
	people := make([]Person, 5)
	for i := 0; i < 5; i++ {
		rand.Seed(int64(i))
		x := firstNames[rand.Intn(5)]
		y := lastNames[rand.Intn(5)]
		people[i] = Person{x, y}

	}
	fmt.Println(people)
}
