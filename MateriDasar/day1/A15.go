package main

import "fmt"

func main() {
	// array adalah kumpulan data dengan tipe yang sama dan
	// memiliki ukuran tetap dan tidak bisa di ubah setelah di deklarasikan
	// contoh array

	var names [4]string // dideklarasikan array dengan tipe data string dan ukuran 4
	names[0] = "Luffy"
	names[1] = "Ayub"
	names[2] = "Subagiya"
	names[3] = "Smith"
	fmt.Println(names[0], names[1], names[2], names[3])

	// inisialisasi array
	var buah = [3]string{"Appel", "Anggur", "coffe"}
	fmt.Println(buah[0], buah[1], buah[2])
	fmt.Println(len(buah))

	// inisialisasi array dengan vertikal
	var hewan = [3]string{
		"Kucing",
		"Anjing",
		"Kelinci",
	}
	fmt.Println(hewan[0], hewan[1], hewan[2])
	fmt.Println(len(hewan))

	// array tanpa jumlah elemen
	var angka = [...]int{10, 20, 30, 40, 50}
	fmt.Println(angka)
	fmt.Println(len(angka))
	angka[4] = 60
	fmt.Println(angka)

	// array multidimensi
	numbers1 := [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	numbers2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(numbers1)
	fmt.Println(numbers2)

	// perulangan dengan array
	for i := 0; i < len(buah); i++ {
		fmt.Println("buah ke-", i, buah[i])
	}

	// perulangan dengan for range untuk array
	for i, buah := range buah {
		fmt.Println("Baris ke-", i, buah)
	}

	// penggunaan underscore jika index tidak digunakan
	for _, hewan := range hewan {
		fmt.Println("Hewan:", hewan)
	}

	// alokasi elemen array menggunakan keyword make
	var electronics = make([]string, 3) // membuat array dengan tipe data string dan ukuran 3
	electronics[0] = "TV"
	electronics[1] = "Radio"
	electronics[2] = "Laptop"
	fmt.Println(electronics)
}
