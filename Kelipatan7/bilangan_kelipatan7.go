package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&number)

	if number%7 == 0 {
		fmt.Printf("%d adalah bilangan kelipatan 7.\n", number)
	} else {
		fmt.Printf("%d bukan bilangan kelipatan 7.\n", number)
	}
}
