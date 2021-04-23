package main

import "fmt"

func main() {

var gender = "female"
switch gender {
case "female":
	fmt.Println("I am a woman")
case "male":
	fmt.Println("I am a man")	
default:
	fmt.Println("I am non conforming")
}

}