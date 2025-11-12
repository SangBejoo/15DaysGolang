package main

import "fmt"

func main() {
	// seleksi kondisi menggunakan keyword if else
	// sintaks if else
	/*
		if kondisi {
			// blok kode jika kondisi bernilai true
		} else {
			// blok kode jika kondisi bernilai false
		}
	*/
	point := 99

	if point >= 90 {
		fmt.Println("Selamat, Anda mendapatkan nilai A")
	} else if point <= 90 {
		fmt.Println("Anda mendapatkan nilai B")
	} else {
		fmt.Println("Point tidak valid")
	}

	// variable temporary pada if else
	if percent := point / 100; percent >= 80 {
		fmt.Println("Anda lulus dengan predikat sangat baik")
	} else if percent <= 80 {
		fmt.Println("Anda lulus dengan predikat baik")
	} else {
		fmt.Println("point and tidak valid")
	}

	// penggunaan switch case

	switch point {
	case 99:
		fmt.Println("Selamat, Anda mendapatkan nilai A")
	case 80:
		fmt.Println("Anda mendapatkan nilai B")
	default:
		fmt.Println("Point tidak valid")
	}

	// switch case dengan multiple case
	switch point {
	case 91, 92, 93, 94, 95, 96, 97, 98, 99, 100:
		fmt.Println("Selamat, Anda mendapatkan nilai A")
	case 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90:
		fmt.Println("Anda mendapatkan nilai B")
	default:
		fmt.Println("Point tidak valid")
	}

	// kurung kurawal pada case dan default
	switch point {
	case 99, 98, 97, 96, 95:
		{
			fmt.Println("Selamat, Anda mendapatkan nilai A")
			fmt.Println("Anda lulus dengan predikat sangat baik")
		}
	case 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90:
		{
			fmt.Println("Anda mendapatkan nilai B")
			fmt.Println("Anda lulus dengan predikat baik")
		}
	default:
		{
			fmt.Println("Point tidak valid")
			fmt.Println("Silakan coba lagi")
		}
	}
	// switch dengan if else
	switch {
	case point >= 90:
		fmt.Println("Selamat, Anda mendapatkan nilai A")
	case (point >= 80) && (point < 90):
		fmt.Println("Anda mendapatkan nilai B")
	default:
		fmt.Println("Point tidak valid")
	}

	// penggunaan fallthrough pada switch case
	switch point {
	case 99:
		fmt.Println("Selamat, Anda mendapatkan nilai A")
		fallthrough
	case 80:
		fmt.Println("Anda mendapatkan nilai B")
	default:
		fmt.Println("Point tidak valid")
	}

	// nested seleksi kondisi
	if point >= 90 {
		switch point {
		case 99, 98, 97, 96, 95:
			fmt.Println("Selamat, Anda mendapatkan nilai A")
		default:
			fmt.Println("Anda mendapatkan nilai A-")
		}
	} else if point >= 80 {
		switch point {
		case 89, 88, 87, 86, 85:
			fmt.Println("Anda mendapatkan nilai B+")
			fallthrough
		default:
			fmt.Println("Anda mendapatkan nilai B")
		}
	} else {
		fmt.Println("Point tidak valid")
	}
}
