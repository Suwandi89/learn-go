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
			character := "САШАРВО"
			for pos, char := range character {
				fmt.Printf("character %U '%c' starts at byte position %d\n", char, char, pos)
			}
			continue
		}
		fmt.Println("Nilai j = ", j)
	}
}
