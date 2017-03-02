package main

import "fmt"

func main()  {
	var nilai string
	nilai = "A"

	switch {
		case nilai == "A":
			fmt.Println("Hebat sekali...")
		case nilai == "B":
			fmt.Println("Kerja bagus...")
		case nilai == "C":
			fmt.Println("Kamu pasti bisa...")
		case nilai == "A":
			fmt.Println("Sedikit lagi berhasil...")
		default:
			fmt.Println("nilai tidak terdaftar")
	}
}
