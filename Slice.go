package main

import "fmt"

func main() {
	var data [5]int
	data[0] = 1
	data[1] = 2
	data[2] = 3
	data[3] = 4
	data[4] = 5
	panjang := len(data)
	top3 := data[0:3]
	fmt.Println("Tiga nilai teratas dari", panjang, "angka yang terdaftar adalah:", top3)
}
