package main

import "fmt"

func main()  {
	var hujan bool
	hujan = true
	if (hujan) {
		fmt.Println("Sedang hujan bawa payung!")
	} else if (!hujan) {
		fmt.Println("Tidak turun hujan, cuaca cerah...")
	}

	var angka int
	angka = 10
	if (angka == 10) {
		fmt.Println("Angka sesuai dengan yang diinginkan: ", angka)
	} else {
		fmt.Println("Angka tidak sesuai dengan yang diinginkan");
	}

	var nama string
	nama = "Andi"
	if (nama == "Andi") {
		fmt.Println("Benar nama saya: ", nama)
	} else if (nama != "Andi") {
		fmt.Println("Nama saya bukan ", nama)
	}

	//nested if adalah if statement didalam if

	fmt.Println("Nested IF statement")

	if (hujan) {
		if (nama == "Andi") {
			fmt.Println("Sekarang sedang hujan",nama,"bawa payung")
		}
	}
}
