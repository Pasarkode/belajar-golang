package main

import "fmt"

type Pegawai struct {
	name  string
	gajih int
	umur  int
}

func main() {
	var susan Pegawai
	susan.name = "Susan"
	susan.gajih = 3000000
	susan.umur = 25

	tampilPegawai(susan)
}

func tampilPegawai(pegawai Pegawai) {
	fmt.Println("Nama:", pegawai.name)
	fmt.Println("Gajih:", pegawai.gajih)
	fmt.Println("Umur", pegawai.umur)
}
