package main

import "fmt"

func main()  {
	var a,b,total float64
	a = 10
	b = 20
	total = a + b
	fmt.Println("Hasil penjumlahan a dan b", total)
	total = b - a
	fmt.Println("Hasil pengurangan b dan a", total)
	total = a * b
	fmt.Println("Hasil perkalian a dan b", total)
	total = b / a
	fmt.Println("Hasil pembagian b dan a", total)
}
