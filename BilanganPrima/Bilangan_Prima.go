package main

import (
	"fmt"
	"math"
)

// Fungsi untuk mengecek apakah bilangan n adalah bilangan prima atau bukan
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	// Menggunakan algoritma uji bilangan prima hingga akar kuadrat n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	// Input bilangan yang akan diperiksa apakah merupakan bilangan prima
	var num int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Println(num, "adalah bilangan prima.")
	} else {
		fmt.Println(num, "bukan bilangan prima.")
	}
}
