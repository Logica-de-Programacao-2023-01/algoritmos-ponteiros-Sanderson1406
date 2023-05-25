package main

import "fmt"

func nosentido(npnt *int, n int) int {
	r := *npnt * n
	return r
}

func main() {
	var n2 int
	var npnt *int = &n2
	n := 9
	fmt.Println("Digite um numero para multiplicar por 9: ")
	fmt.Scanln(&n2)
	r := nosentido(npnt, n)
	fmt.Println("Resultado: ", r)
}
