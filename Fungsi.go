package main

import "fmt"

func halo(nama string)  {
	fmt.Println("Halo", nama)
}

func penjumlahan(a int, b int) {
	var hasil int
	hasil = a + b
	fmt.Println(a,"+",b,"=",hasil)
}

func perkalian(a int, b int) int {
	var hasil int
	hasil = a * b
	return hasil
}

func main() {
	var hasil int
	hasil = perkalian(5,25)
	halo("Andi")
	penjumlahan(2,10)
	fmt.Println(hasil)
}
