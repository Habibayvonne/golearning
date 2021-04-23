package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var age int
	var name string
	var gender rune
	var salary float32
	var alive bool
	age = 25
	name = "Grinch"
	gender = 'T'
	salary = 123.50
	alive = true
	fmt.Println("my name is", name, "I am", age, "I earn", salary, "my status is", alive, "And ", gender)
	fmt.Printf("myname is %s I am %d I earn %f \tmy status is %t and  %c\n", name, age, salary, alive, gender)
}
