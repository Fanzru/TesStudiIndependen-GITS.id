package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var kalimat string
	fmt.Print("Masukan Kalimat : ")
	scanner := bufio.NewReader(os.Stdin)
	kalimat, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Print(err)
	}
	i := 0
	cek := 0
	j := len(kalimat) - 2

	for i < len(kalimat)-1 {
		if string(kalimat[i]) == string(kalimat[j]) {
			cek++
		}
		j--
		i++
	}

	if cek == len(kalimat)-1 {
		fmt.Println("True Palindrom")
	} else {
		fmt.Println("False Palindrom")
	}
}
