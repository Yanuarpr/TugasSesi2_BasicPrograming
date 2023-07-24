package main

import "fmt"

func main() {
	var tinggi, panjangAlas, panjangSisiMiring float64

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scanln(&tinggi)

	fmt.Print("Masukkan panjang alas trapesium: ")
	fmt.Scanln(&panjangAlas)

	fmt.Print("Masukkan panjang sisi miring trapesium: ")
	fmt.Scanln(&panjangSisiMiring)

	luas := (panjangAlas + panjangSisiMiring) * tinggi / 2

	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}
