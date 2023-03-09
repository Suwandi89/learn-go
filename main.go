package main

import "fmt"

func main() {
	i := 21
	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Println("%")

	j := true
	fmt.Println(j)

	fmt.Printf("\n%b \n", i)
	fmt.Println("?")
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%X \n", 15)
	fmt.Printf("%U \n\n", 'Я')

	var k float64 = 123.456

	fmt.Printf("%f \n", k)
	fmt.Printf("%E \n", k)

}
