package main

import "fmt"

func main() {
	var alive bool
	alive = false
	if alive {
		fmt.Println("you are alive ")
	} else {
		fmt.Println("you are dead")
	}
	var age int
	age = 18
	if age < 18 {
		fmt.Println("Service denied ")
	} else {
		fmt.Println("How may I be of help")
	}

	var  gender = "female"
	age = 45
	if gender == "female" && age >= 18 {
		fmt.Println("you are hired")
	}else {
		fmt.Println("Try everywhere")	
	}
	


	
	age = 65

	if age < 18 {
		fmt.Println("Service denied ")
	} else if age <= 55 {
		fmt.Println("How may I be of help")

	}else { 
		fmt.Println("Janet,will assist you in amoment")
	}


	

}
