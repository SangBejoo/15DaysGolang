package main

import "fmt"

func main() {

	// slice itu dibagi menjadi 3 bagian yaitu
	// pointer : menunjuk ke index pertama array yang direferensikan
	// length  : panjang slice, yaitu jumlah elemen yang ada di dalam slice
	// capacity: kapasitasnya yang dihitung dari pointer sampai akhir array yang direferensikan

	// slice ini adalah data referensi array dimana
	// slice ini menunjuk ke array yang ada di memory
	// jadi ketika slice diubah, array yang direferensikan juga ikut berubah
	// dan sebaliknya

	// contoh pembuatan slice
	var electronic = []string{"TV", "Radio", "Laptop", "Smartphone"}
	fmt.Println(electronic[0])

	// salah satu perbedaan slice dengan array adalah
	// pada deklarasi variablenya, jika jumlah element tidak ditentukan
	// maka itu adalah slice, bukan array

	var buahA = [2]string{"apple", "banana"}             // array
	var buahB = []string{"apple", "banana"}              // slice
	var buahC = [...]string{"apple", "banana", "cherry"} // array

	fmt.Println(buahA)
	fmt.Println(buahB)
	fmt.Println(buahC)

	// array itu kumpulan nilai atau elemen sedangkan slice itu refrensinya
	var celana = []string{"jeans", "chino", "cargo"}

	var celana1 = celana[0:3] // slicing artinya di mulai dari 0 sampai index ke 2 (tidak termasuk index ke 2)
	var celana2 = celana[1:3]

	var celana3 = celana1[0:1]
	var celana4 = celana2[:2]

	fmt.Println()
	fmt.Println(celana)
	fmt.Println(celana1)
	fmt.Println(celana2)
	fmt.Println(celana3)
	fmt.Println(celana4)

	fmt.Println()
	celana[1] = "short"

	fmt.Println(len(celana))
	fmt.Println(celana1)
	fmt.Println(celana2)
	fmt.Println("jumlah", len(celana3))
	fmt.Println(celana4)

	// Fungsi len() digunakan untuk menghitung jumlah elemen slice yang ada
	// Fungsi cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice

	fmt.Println(cap(celana))
	fmt.Println(cap(celana1))
	fmt.Println(cap(celana2))
	fmt.Println("kapasitas", cap(celana3))
	fmt.Println(cap(celana4))

	// Fungsi append() digunakan untuk menambahkan elemen pada slice.

	var celana5 = append(celana, "jogger", "sweatpants")
	fmt.Println(celana5)
	fmt.Println(cap(celana5))
	fmt.Println(celana)

	// Fungsi copy() digunakan untuk men-copy elements slice pada src (parameter ke-2), ke dst (parameter pertama).
	n := copy(celana2, celana5)
	fmt.Println(n)

	fmt.Print()
	dst := make([]string, 2)
	src := []string{"apple", "banana", "cherry"}
	n = copy(dst, src)
	fmt.Println("dst:", dst)
	fmt.Println("src:", src)
	fmt.Println("n:", n)

	fmt.Println()
	sumber := []int{10, 20, 30}
	tujuan := []int{30, 20, 10}
	jumlah := copy(tujuan, sumber)
	// coppy ini menyalin elemen dari slice sumber ke slice tujuan
	// dan mengembalikan jumlah elemen yang disalin
	// jadinya nilai di sumber dan tujuan akan berubah sesuai dengan hasil penyalinan

	fmt.Println("sumber:", sumber)
	fmt.Println("tujuan:", tujuan)
	fmt.Println("jumlah yang disalin:", jumlah)

	var bulan = []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	fmt.Println()
	bulanBaik := bulan[5:8]
	fmt.Println(bulanBaik)
	fmt.Println("length bulanBaik:", len(bulanBaik))
	fmt.Println("capacity bulanBaik:", cap(bulanBaik))

	fmt.Println()
	bulanAkhirTahun := bulan[10:]
	fmt.Println(bulanAkhirTahun)
	fmt.Println("length bulanAkhirTahun:", len(bulanAkhirTahun))
	fmt.Println("capacity bulanAkhirTahun:", cap(bulanAkhirTahun))

	fmt.Println()
	awalBulan := bulan[0:3]
	fmt.Println(awalBulan)
	fmt.Println("length awalBulan:", len(awalBulan))
	fmt.Println("capacity awalBulan:", cap(awalBulan))

	fmt.Println()
	semuaBulan := bulan[:]
	fmt.Println(semuaBulan)
	fmt.Println("length semuaBulan:", len(semuaBulan))
	fmt.Println("capacity semuaBulan:", cap(semuaBulan))

	fmt.Println()
	// append slice
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"

	fmt.Println(days)
	fmt.Println(daySlice1)
	// jika mengubah nilai pada slice maka array aslinya juga akan berubah

	fmt.Println()
	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Libur Lama"
	fmt.Println(daySlice2)
	fmt.Println(daySlice1)
	fmt.Println(days)
	// ketika menambahkan elemen pada slice yang melebihi kapasitasnya
	// maka akan dibuatkan array baru di memory dan menyalin elemen-elemen dari array lama ke array baru
	// sehingga array aslinya tidak akan berubah

	// make slice
	// make([]tipe data, panjang atau length, kapasitas atau capacity)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ayubs"
	newSlice[1] = "Ganteng"
	newSlice2 := append(newSlice, "Ayubs11")

	fmt.Println()
	fmt.Println(newSlice2)
	fmt.Println("length newSlice:", len(newSlice2))
	fmt.Println("capacity newSlice:", cap(newSlice2))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println()
	fmt.Println("fromSlice:", fromSlice)
	fmt.Println("toSlice  :", toSlice)

}
