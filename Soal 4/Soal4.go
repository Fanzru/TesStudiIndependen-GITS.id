package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var kalimat, revers string
	fmt.Print("Masukan Kalimat : ")
	scanner := bufio.NewReader(os.Stdin)
	kalimat, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Print(err)
	}
	i := len(kalimat) - 1
	for i >= 0 {
		revers = revers + string(kalimat[i])
		i--
	}
	fmt.Println("Hasil Revers di Bawah ", revers)
}
