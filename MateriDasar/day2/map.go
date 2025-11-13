package main

import "fmt"

// map adalah tipe data assosiatif di go yang berbentuk key-value pair

func main() {
	// penggunaan map
	var ayam map[string]int
	ayam = map[string]int{}

	ayam["dada"] = 20
	ayam["paha"] = 15
	ayam["sayap"] = 10

	fmt.Println(ayam)
	fmt.Println("januari:", ayam["dada"])

	// A17.2 Inisialisasi nilai map
	// var data map[string]int
	// data["one"] = 1
	// akan error karena map belum diinisialisasi
	// fmt.Println(data)

	// solusi
	data := map[string]int{}
	data["one"] = 1
	fmt.Println(data)

	// map itu menggabungkan key dan value

	var buku = map[string]int{"novel": 50000, "komik": 30000, "ensiklopedia": 100000}
	fmt.Println(buku)

	// cara vertical
	var gelas = map[string]int{
		"plastik": 5000,
		"kaca":    20000,
		"keramik": 15000,
	}
	fmt.Println(gelas)

	for key, val := range gelas {
		fmt.Println(key, "harga:", val)
	}

	for key, val := range ayam {
		fmt.Println(key, "jumlah:", val)
	}

	// menghapus elemen pada map
	fmt.Println()
	fmt.Println(len(buku))
	fmt.Println("sebelum dihapus:", buku)

	delete(buku, "novel")
	fmt.Println()
	fmt.Println("setelah dihapus:", buku)
	fmt.Println(len(buku))

	fmt.Println()
	// mengecek apakah key ada di map
	var value, exists = gelas["kaca"]
	if exists {
		fmt.Println("harga gelas kaca adalah", value)
	} else {
		fmt.Println("gelas kaca tidak ditemukan")
	}

	fmt.Println()
	// kombinasikan slice dan map
	var parentSiswa = []map[string]string{
		{"Ayah": "Pardjo", "Ibu": "Surati", "NamaSiswa": "Budi"},
		{"Ayah": "Hakimi", "Ibu": "Surkini", "NamaSiswa": "Ani"},
		{"Ayah": "Joko", "Ibu": "Siti", "NamaSiswa": "Sari"},
		{"Ayah": "Solerah", "Ibu": "Solehah", "NamaSiswa": "Udin"},
	}

	for key, val := range parentSiswa {
		fmt.Println("Data Siswa ke-", key+1)
		fmt.Println("Ayah:", val["Ayah"])
		fmt.Println("Ibu:", val["Ibu"])
		fmt.Println("Nama Siswa:", val["NamaSiswa"])
		fmt.Println()
	}

}
