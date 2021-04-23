package main

import "fmt"

func main() {
	var age int
	var gender = "female"
	age = 34
	if gender == "male" && age >= 18 && age < 36 {
		fmt.Println("you are ready for war")

	} else if gender == "male" && age > 35 {
		fmt.Println("You are ready for babies")

	} else {
		fmt.Println("welcome lady")

	}

}
