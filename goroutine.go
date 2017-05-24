package main

import (
	"fmt"
	"strings"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
	finds := strings.Contains(input, "%")

	if finds != false {
		str := strings.Replace(input, "%", "", 1)
		fmt.Print("Test "+str)
	}
}
