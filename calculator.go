package main

import "fmt"

func main(){

	var firstNum float64
	var secondNum float64
	var operator string

	var answer float64

	fmt.Printf("Enter first number: ")
	fmt.Scan(&firstNum)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&secondNum)


	fmt.Printf("Would you liek to ( +, -, x, ÷): ")
	fmt.Scan(&operator)

	if operator == "+" {
		answer = firstNum + secondNum
	}	else if operator == "-" {
		answer = firstNum - secondNum
	} else if operator == "x" || operator == "*"{
		answer = firstNum * secondNum
	} else if operator == "/" || operator == "÷"{
		answer = firstNum / secondNum
	} else {
		fmt.Print("Error")
	}


	fmt.Printf("Calculated: %v", answer)
}