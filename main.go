package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v", alex)
}
