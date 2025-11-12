package main

import (
	"fmt"
	"strings"
)

// ini string repeat
func printPyramid(height int) {
	for i := 0; i < height; i++ {
		spaces := strings.Repeat(" ", height-i-1)
		stars := strings.Repeat("*", 2*i+1)
		fmt.Println(spaces + stars)
	}
}

// penggunaan loop
func piramidLoop(height int) {
	for i := 0; i < height; i++ {
		// cetak spasi
		for j := 0; j < height-i-1; j++ {
			fmt.Print(" ")
		}
		// cetak bintang
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
}

func main() {
	// pattern Printer - Piramida Bintang
	printPyramid(5)
	fmt.Println()
	piramidLoop(5)

}
