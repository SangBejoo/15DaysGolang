package main

import "fmt"

func fizzBuzz() {
	for i := 1; i <= 110; i++ {
		if i%3 == 0 && i%5 == 0 && i%7 == 0 {
			fmt.Println("FizzBuzzRizz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	// fizzBuzz
	fizzBuzz()
}
