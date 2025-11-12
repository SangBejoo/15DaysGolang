package main

import (
	"fmt"
)

func getNumberInput(prompt string) float32 {
	var num float32
	fmt.Println(prompt)
	for {
		if _, err := fmt.Scan(&num); err != nil {
			fmt.Println("Invalid input. Please enter a valid number:")
			var discard string
			fmt.Scanln(&discard) // discard invalid input
		} else {
			return num
		}
	}
}

func getOperatorInput(prompt string) string {
	var operator string
	fmt.Println(prompt)
	for {
		if _, err := fmt.Scan(&operator); err != nil {
			fmt.Println("Invalid input. Please enter a valid operator (+, -, *, /):")
			var discard string
			fmt.Scanln(&discard) // discard invalid input
		} else {
			switch operator {
			case "+", "-", "*", "/":
				return operator
			default:
				fmt.Println("Invalid operator. Please enter one of (+, -, *, /):")
			}
		}
	}
}

func main() {
	var num1, num2 float32
	var operator string
	fmt.Println("calculatorCli")

	for {
		num1 = getNumberInput("Enter number 1: ")
		num2 = getNumberInput("Enter number 2: ")

		operator = getOperatorInput("Enter operator: ")

		switch operator {
		case "+":
			fmt.Printf("%.2f + %.2f = %.3f\n", num1, num2, num1+num2)
		case "-":
			fmt.Printf("%.2f - %.2f = %.3f\n", num1, num2, num1-num2)
		case "*":
			fmt.Printf("%.2f * %.2f = %.3f\n", num1, num2, num1*num2)
		case "/":
			if num2 != 0 {
				fmt.Printf("%.2f / %.2f = %.3f\n", num1, num2, num1/num2)
			} else {
				fmt.Println("Error: Division by zero")
			}
		default:
			fmt.Println("Invalid operator")
		}
		var continueCalc string
		fmt.Println("Lanjutkan kalkulasi? (y/n): ")
		fmt.Scan(&continueCalc)
		if continueCalc != "y" && continueCalc != "Y" {
			fmt.Println("Exiting calculator. Goodbye!")
			break
		} else {
			fmt.Println()
		}
	}

}
