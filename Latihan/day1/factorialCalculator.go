package main

import "fmt"

// factorial calator adalah program untuk menghitung faktorial dari sebuah bilangan bulat non-negatif.

// iterative adalah metode perhitungan yang menggunakan pengulangan (looping) untuk mencapai hasil akhir.
func factorialIterative(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// recursive adalah metode perhitungan yang memanggil dirinya sendiri untuk mencapai hasil akhir.
func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func main() {
	fmt.Println(factorialIterative(7))
	fmt.Println(factorialRecursive(5))
}
