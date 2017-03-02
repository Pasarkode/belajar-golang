package main

import "fmt"

func main() {
	var murid [5]string

	murid[0] = "Andi"
	murid[1] = "Budi"
	murid[2] = "Chandra"
	murid[3] = "Diki"
	murid[4] = "Evan"

	for i := 0; i < 5; i++ {
		fmt.Println(murid[i])
	}

	fmt.Println("dibawah ini menggunakan fungsi range")

	for i := range murid {
		fmt.Println(murid[i])
	}
}
