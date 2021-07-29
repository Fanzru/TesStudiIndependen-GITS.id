package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var jam, menit, detik int
	var Keterangan string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Input was: %q\n", line)
	}
	fmt.Println("Exp : 12:00:00 PM")
	fmt.Print("Masukan Waktu : ")
	fmt.Scanf("%d:%d:%d %s", &jam, &menit, &detik, &Keterangan)

	// Error Handler
	if (jam == 12) && ((menit != 00) || (detik != 00)) {
		fmt.Println("Waktu Tidak Valid")
	} else if (jam > 12) || (jam < 0) {
		fmt.Println("Waktu Tidak Valid")
	} else if (menit >= 60) || (menit < 0) {
		fmt.Println("Waktu Tidak Valid")
	} else if (detik >= 60) || (detik < 0) {
		fmt.Println("Waktu Tidak Valid")
	} else if Keterangan != "AM" && Keterangan != "PM" {
		fmt.Println("Waktu Tidak Valid")
	} else {
		if jam == 12 {
			if Keterangan == "AM" {
				fmt.Println("00:00")
			} else {
				fmt.Println("12:00")
			}
		} else if menit <= 10 {
			fmt.Printf("%d:0%d\n", jam, menit)
		} else {
			fmt.Printf("%d:%d\n", jam, menit)
		}
	}
}
