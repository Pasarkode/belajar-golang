package main

import "fmt"

var global string

func main()  {
	//integer untuk bilangan bulat
	var i int
	//float untuk bilangan pecahan atau desimal
	var f float64
	//string digunakan untuk variable yang berisi angka dan huruf
	var s string
	//variable integer tidak memerlukan tanda kutip("")
	i = 10
	//didalam pemrograman tanda koma(,) digantikan dengan tanda titik(.)
	f = 3.14
	//untuk variable string dibutuhkan tanda kutip("")
	s = "Saya Taufik R. Firdaus"
	//ini merupakan variable global. variable yang dapat dipanggil difungsi  mana saja
	global = "Saya variable global"
	//untuk menampilkan isi variable tidak memerlukan tanda kutip("")
	fmt.Println(i)

	fmt.Println(f)
	//untuk menampilkan dua variable dalam satu kali keluaran digunakan tanda koma(,)
	fmt.Println(i, f)

	fmt.Println("Halo dunia... ",s)

	fmt.Println(global)
}
