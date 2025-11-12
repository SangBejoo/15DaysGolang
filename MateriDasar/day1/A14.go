package main

import "fmt"

func main() {
	// perulangan menggunakan for
	for i := 10; i > 5; i++ { // temporary variable
		fmt.Println("Perulangan ke-", i)
		if i >= 15 {
			break
		}
	}
	// perulangan dengan kondisi saja dan argumen
	i := 1
	for i < 5 {
		fmt.Println("Perulangan ke-", i)
		i++
	}

	// perulangan tanpa argumen
	i = 10
	for {
		fmt.Println("Perulangan ke-", i)
		i++
		if i > 15 {
			break
		}
	}

	// penggunaan for range
	var xs = "123"
	for i, v := range xs {
		fmt.Printf("Index: %d, Value: %c\n", i, v)
	}

	var ys = [5]int{1, 2, 3, 4, 5} // array
	for _, y := range ys {
		fmt.Println("value", y)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("value", v)
	}

	var kvs = map[byte]int{'a': 1, 'b': 2, 'c': 3} // map
	for k, v := range kvs {
		fmt.Println("value", k, v)
	}

	// jika k dan atau v di abaikan
	for range kvs {
		fmt.Println("done")
	}

	// cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Println("Perulangan ke-", i)
	}

	// penggunaan break and continue
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 20 {
			break
		}
		fmt.Println("Angka", i)
	}

	// perulangan bersarang atau nested loop
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// pemanfaatan label pada for loop
void:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 3 {
				break void
			}
			fmt.Println("matriks [", i, "][", j, "] ")
		}
	}

}
