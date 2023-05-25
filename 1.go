package main

import "fmt"

func somar(npnt *int, n int) int {
	var soma int
	for i := 1; i <= n; i++ {
		soma += i
	}
	*npnt = soma
	return *npnt
}

func main() {
	var n int
	fmt.Println("Digie n: ")
	fmt.Scanln(&n)
	var npnt *int = &n
	r := somar(npnt, n)
	fmt.Println("A soma dos n primeiros números natuaris é: ", r)
}
