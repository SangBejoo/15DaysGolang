package main

import "fmt"

func main() {
	// Operators Aritmetika
	/*
		Tanda	Penjelasan
		+	penjumlahan
		-	pengurangan
		*	perkalian
		/	pembagian
		%	modulus / sisa hasil pembagian
	*/
	// contoh penggunaan operator aritmetika
	a, b, c, d := 10, 5, 15, 20
	a = b + c
	b = d - a
	c = a * b
	d = d / b
	a = c % b
	fmt.Println(a, b, c, d)
	// operasi perbandingan
	/*
		Tanda	Penjelasan
		==	apakah nilai kiri sama dengan nilai kanan
		!=	apakah nilai kiri tidak sama dengan nilai kanan
		<	apakah nilai kiri lebih kecil daripada nilai kanan
		<=	apakah nilai kiri lebih kecil atau sama dengan nilai kanan
		>	apakah nilai kiri lebih besar dari nilai kanan
		>=	apakah nilai kiri lebih besar atau sama dengan nilai kanan
	*/
	// contoh penggunaan operator perbandingan
	x, y := 10, 20
	fmt.Println(x == y) // false
	fmt.Println(x != y) // true
	fmt.Println(x < y)  // true
	fmt.Println(x <= y) // true
	fmt.Println(x > y)  // false
	fmt.Println(x >= y) // false
	// operasi logika
	/*
		Tanda	Penjelasan
		&&	kiri dan kanan
		||	kiri atau kanan
		!	negasi / nilai kebalikan
	*/
	// contoh penggunaan operator logika
	m, n := true, false
	fmt.Println(m && n) // false
	fmt.Println(m || n) // true
	fmt.Println(!m)     // false
}
