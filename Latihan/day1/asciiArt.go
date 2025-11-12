package main

import "strings"

var asciiArtMap = map[string][]string{
	"A": {
		"  A  ",
		" A A ",
		"AAAAA",
		"A   A",
		"A   A",
	},
	"B": {
		"BBBB ",
		"B   B",
		"BBBB ",
		"B   B",
		"BBBB ",
	},
	"C": {
		" CCCC",
		"C    ",
		"C    ",
		"C    ",
		" CCCC",
	},
	"D": {
		"DDDD ",
		"D   D",
		"D   D",
		"D   D",
		"DDDD ",
	},
	"E": {
		"EEEEE",
		"E    ",
		"EEE  ",
		"E    ",
		"EEEEE",
	},
	"F": {
		"FFFFF",
		"F    ",
		"FFF  ",
		"F    ",
		"F    ",
	},
	"G": {
		" GGGG",
		"G    ",
		"G GGG",
		"G   G",
		" GGG ",
	},
	"H": {
		"H   H",
		"H   H",
		"HHHHH",
		"H   H",
		"H   H",
	},
	"I": {
		"IIIII",
		"  I  ",
		"  I  ",
		"  I  ",
		"IIIII",
	},
	"J": {
		"JJJJJ",
		"    J",
		"    J",
		"J   J",
		" JJJ ",
	},
	"K": {
		"K   K",
		"K  K ",
		"KKK  ",
		"K  K ",
		"K   K",
	},
	"L": {
		"L    ",
		"L    ",
		"L    ",
		"L    ",
		"LLLLL",
	},
	"M": {
		"M   M",
		"MM MM",
		"M M M",
		"M   M",
		"M   M",
	},
	"N": {
		"N   N",
		"NN  N",
		"N N N",
		"N  NN",
		"N   N",
	},
	"O": {
		" OOO ",
		"O   O",
		"O   O",
		"O   O",
		" OOO ",
	},
	"P": {
		"PPPP ",
		"P   P",
		"PPPP ",
		"P    ",
		"P    ",
	},
	"Q": {
		" QQQ ",
		"Q   Q",
		"Q   Q",
		"Q  QQ",
		" QQQQ",
	},
	"R": {
		"RRRR ",
		"R   R",
		"RRRR ",
		"R  R ",
		"R   R",
	},
	"S": {
		" SSSS",
		"S    ",
		" SSS ",
		"    S",
		"SSSS ",
	},
	"T": {
		"TTTTT",
		"  T  ",
		"  T  ",
		"  T  ",
		"  T  ",
	},
	"U": {
		"U   U",
		"U   U",
		"U   U",
		"U   U",
		" UUU ",
	},
	"V": {
		"V   V",
		"V   V",
		"V   V",
		" V V ",
		"  V  ",
	},
	"W": {
		"W   W",
		"W   W",
		"W W W",
		"WW WW",
		"W   W",
	},
	"X": {
		"X   X",
		" X X ",
		"  X  ",
		" X X ",
		"X   X",
	},
	"Y": {
		"Y   Y",
		" Y Y ",
		"  Y  ",
		"  Y  ",
		"  Y  ",
	},
	"Z": {
		"ZZZZZ",
		"   Z ",
		"  Z  ",
		" Z   ",
		"ZZZZZ",
	},
	" ": {
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	},
}

func printAsciiArt(word string) {
	// ubah ke uppercase untuk mencocokkan dengan kunci pada asciiArtMap
	word = strings.ToUpper(word)

	// asumsi setiap huruf direpresentasikan dalam 7 baris
	lines := make([]string, 5)

	// untuk setiap huruf dalam kata, penggunaan baris perbaris
	for _, char := range word {
		letter := string(char)
		if art, exists := asciiArtMap[letter]; exists {
			for i, line := range art {
				lines[i] += line + "  "
			}
		} else {
			for i := range lines {
				lines[i] += asciiArtMap[" "][i] + "  "
			}
		}
	}
	for _, line := range lines {
		println(line)
	}
}

func main() {
	// contoh penggunaan
	printAsciiArt("HELLO WORLD")
}
