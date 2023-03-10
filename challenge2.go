package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	for j := 0; j < 11; j++ {
		if j == 5 {
			var character = [7]rune{'С', 'А', 'Ш', 'А', 'Р', 'В', 'О'}
			for k := 0; k < len(character); k++ {
				fmt.Printf("character %U '%c' starts at byte position %d\n", character[k], character[k], 2*k)
			}
			continue
		}
		fmt.Println("Nilai j = ", j)
	}
}
