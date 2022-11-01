package main

import "fmt"



func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 
	// byte - alias for uint8
	//rune - alias for int32
	// float32 float64
	// complex64 complex128
	
	// Using var
	var age = 18
	const isCool = true
	var size float32 = 2.3
	
	// Shorthand
	// name := "Jordon"
	// email := "jordon@gmail.com"

	name, email := "Jordon", "jordon@influxed.io"

	fmt.Println(name, age, email, isCool)
	fmt.Printf("%T\n", size)
}