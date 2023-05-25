package main

import (
	"fmt"
	"math"
)

func media(npnt *float64) float64 {
	var soma float64 = 0
	var resu float64 = 0
	soma = *npnt + math.Pi
	resu = soma / 2
	return resu
}

func main() {
	var n float64
	fmt.Println("Escreva um numero: ")
	fmt.Scanln(&n)
	fmt.Println("Numero antes: ", n)
	var npnt *float64 = &n
	r := media(npnt)
	fmt.Println("Numero depois: ", r)
}
