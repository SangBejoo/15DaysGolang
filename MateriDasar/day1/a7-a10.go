package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// ini commentar single line
	/*
		ini commentar
		multi line
	*/
	// variable merupakan manifest typing
	var name string = "John Doe"
	age := 30
	// multi variable
	first, second := 1, 2
	fmt.Println(name, age, first, second)

	// tipe data, ada dua yaitu numerik dan desimal
	/*
		Tipe data	Cakupan bilangan
		uint8	0 ↔ 255
		uint16	0 ↔ 65535
		uint32	0 ↔ 4294967295
		uint64	0 ↔ 18446744073709551615
		uint	sama dengan uint32 atau uint64 (tergantung nilai)
		byte	sama dengan uint8
		int8	-128 ↔ 127
		int16	-32768 ↔ 32767
		int32	-2147483648 ↔ 2147483647
		int64	-9223372036854775808 ↔ 9223372036854775807
		int	sama dengan int32 atau int64 (tergantung nilai)
		rune	sama dengan int32
	*/
	// tipe data numerik desimal
	decimalNumber := 3.14 // float64
	fmt.Println(decimalNumber)

	// tipe data boolean
	exist := true
	fmt.Println(exist)

	// tipe data string
	var message string = "Hello, Go!"
	fmt.Println(message)

	// tipe data nil & zero value
}
