package main

import "fmt"

func main() {
	/*
		var years int
		var position = "ma"
		years = 32




		switch position {
		case "manager" && years <= 23:
			fmt.Println("salary is 8000")
		case "manager" && years > 23 && years <= 45:
			fmt.Println("salary is 10000")
		case "janitor" && years < 5:
			fmt.Println("salary is 40")
		case "janitor" && years > 5 && years <= 10:
			fmt.Println("salary is 100")
		default:
			fmt.Println("salary is 1000")
		sitch does not allow use of &&
		}
	*/

	var position string
	var yearsInService int
	yearsInService = 15
	position = "janitor"
	if position == "manager" {
		if yearsInService >= 0 && yearsInService <= 23 {
			fmt.Println("salary is ", 8000)
		} else if yearsInService <= 45 {

		} else {
			// ignore / print whatever you want
		}
	} else { // janitor
		if yearsInService < 5 {
			fmt.Println("salary is", 40)
		} else if yearsInService >= 5 && yearsInService <= 10 {
			fmt.Println("salary is ", 100)
		} else {
			fmt.Println("salary is ", 1000)
		}

	}

}
