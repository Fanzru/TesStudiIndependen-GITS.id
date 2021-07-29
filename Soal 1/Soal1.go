package main

import "fmt"

func main() {

	var nilai int
	fmt.Print("Input Angka : ")
	fmt.Scanln(&nilai)
	if (nilai%3 == 0) && (nilai%5 == 0) {
		fmt.Println("Hello World")
	} else if nilai%3 == 0 {
		fmt.Println("Hello")
	} else if nilai%5 == 0 {
		fmt.Println("World")
	} else {
		fmt.Println("Error")
	}
}
