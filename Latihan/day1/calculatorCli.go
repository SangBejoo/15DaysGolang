package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

/*
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
*/

func getFloat(prompt string) float64 {
	for {
		fmt.Print(prompt)
		input := readLine()
		if num, err := strconv.ParseFloat(input, 64); err == nil {
			return num
		}
		fmt.Println("Invalid input. Please enter a valid number.")
	}
}

func getOperator() string {
	valid := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	for {
		fmt.Print("Enter valid operator (+, -, *, /): ")
		op := strings.TrimSpace(readLine())
		if valid[op] {
			return op
		}
		fmt.Println("Invalid operator. Please try again.")
	}
}

func askContinue() bool {
	for {
		fmt.Print("Lanjutkan kalkulasi? (y/n): ")
		ans := strings.ToLower(strings.TrimSpace(readLine()))
		if ans == "y" || ans == "yes" {
			return true
		}
		if ans == "n" || ans == "no" {
			return false
		}
		fmt.Println("Invalid input. Please enter 'y' or 'n'.")
	}
}

/*
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
*/

func main() {
	fmt.Println("====Calculator CLI====")
	fmt.Println("Supported desimal, q to quit")

	for {
		num1 := getFloat("Enter first number :  ")
		num2 := getFloat("Enter second number : ")
		operator := getOperator()

		var result float64
		valid := true

		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
				valid = false
			} else {
				result = num1 / num2
			}
		}

		if valid {
			fmt.Printf("Result: %.2f %s %.2f = %.4f\n", num1, operator, num2, result)
		} else {
			fmt.Println()
		}

		if !askContinue() {
			fmt.Println("Bye Bye!")
			break
		}
		fmt.Print()
	}

}
