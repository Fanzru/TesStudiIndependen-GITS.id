package main

import (
	"fmt"
	"regexp"
)

func main() {
	var email string

	fmt.Print("Masukan Email : ")
	fmt.Scanln(&email)
	if len(email) <= 20 {
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+[.id][.co.id]$`)
		if emailRegex.MatchString(email) {
			fmt.Println("Email Valid")
		} else {
			fmt.Println("Email Tidak Valid")
		}
	} else {
		fmt.Println("Panjang Email Lebih dari 20 karakter")
	}
}
