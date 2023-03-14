package main

import "fmt"

func main() {
	kalimat := "selamat malam"

	for _, char := range kalimat {
		fmt.Printf("%c\n", char)
	}

	charMap := make(map[string]int)

	for _, char := range kalimat {
		charMap[string(char)]++
	}

	fmt.Println(charMap)
}
