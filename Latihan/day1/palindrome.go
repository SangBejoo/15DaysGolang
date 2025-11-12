package main

import (
	"fmt"
	"strings"
)

/*
palindrome adalah
kata atau kalimat yang jika dibaca dari depan maupun belakang tetap sama
contoh:
- katak
- kasur rusak
- A nut for a jar of tuna
*/

func isPalindrome(y string) bool {
	// hilangkan spasi dan ubah ke lowercase
	/*
		kenapa harus dihilangkan spasi dan di lowercase?
		karena agar pengecekan palindrome lebih akurat
		contoh: "A nut for a jar of tuna"
		jika tidak dihilangkan spasi dan di lowercase, maka hasilnya tidak akan sama
		"A nut for a jar of tuna" != "anutforajaroftuna"
	*/

	y = strings.ReplaceAll(y, " ", "")
	y = strings.ToLower(y)

	// bandingkan dengan string yang dibalik
	for i := 0; i < len(y)/2; i++ {
		if y[i] != y[len(y)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal)"))
	fmt.Println(isPalindrome("Kasur Rusak"))
	fmt.Println(isPalindrome("Hello World"))
}
