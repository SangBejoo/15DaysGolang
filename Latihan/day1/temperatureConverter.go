package main

import "fmt"

/*
Celsius ke Fahrenheit: F = (C * 9/5) + 32

Fahrenheit ke Celsius: C = (F - 32) * 5/9

Celsius ke Kelvin: K = C + 273.15

Kelvin ke Celsius: C = K - 273.15

Fahrenheit ke Kelvin: K = (F - 32) * 5/9 + 273.15

Kelvin ke Fahrenheit: F = (K - 273.15) * 9/5 + 32
*/

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}
func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}
func fahrenheitToKelvin(f float64) float64 {
	return (f-32)*5/9 + 273.15
}
func kelvinToFahrenheit(k float64) float64 {
	return (k-273.15)*9/5 + 32
}

func main() {
	fmt.Println(celsiusToFahrenheit(273.15))
	fmt.Println(fahrenheitToCelsius(332))
	fmt.Println(celsiusToKelvin(34))
	fmt.Println(kelvinToCelsius(550))
	fmt.Println(fahrenheitToKelvin(32))
	fmt.Println(kelvinToFahrenheit(273.15))

}
