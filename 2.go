package main

import "fmt"

func verifi(npnt *int) int {
	if *npnt%2 == 0 {
		*npnt = 0
	} else if *npnt%2 != 0 {
		*npnt = 1
	}
	return *npnt
}

func main() {
	var n int
	fmt.Println("Escreva um número: ")
	fmt.Scanln(&n)
	var npnt *int = &n
	r := verifi(npnt)
	fmt.Println(r)
}
