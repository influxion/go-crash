package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" || p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Maya", lastName: "Baker", city: "Winnipeg", gender: "f", age: 18}
	// Alt
	person2 := Person{"Bob", "Johnson", "New Jersery", "m", 28}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1.age)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Nichols")
	person2.getMarried("Bast")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
