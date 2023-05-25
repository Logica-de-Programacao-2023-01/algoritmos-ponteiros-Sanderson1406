package main

import "fmt"

func soma(npnt *int) int {
	digs := []int{}
	num := *npnt
	soma := 0
	if *npnt < 0 {
		*npnt = -(*npnt)
	}
	for num > 0 && len(digs) < 2 {
		dig := num % 10
		digs = append(digs, dig)
		num /= 10
	}
	for _, san := range digs {
		soma += san
	}
	*npnt = soma
	return *npnt
}

func main() {
	var n int
	fmt.Println("Digite um número: ")
	fmt.Scanln(&n)
	var npnt *int = &n
	r := soma(npnt)
	fmt.Println("A soma dos 2 ultimos digitos é: ", r)
}
