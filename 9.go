package main

import "fmt"

type Livr struct {
	titulo string
	autor  string
	valor  float64
}

func descont(l *Livr) Livr {
	l.valor = l.valor - (l.valor * 0.1)
	return Livr{
		titulo: l.titulo,
		autor:  l.autor,
		valor:  l.valor,
	}
}

func main() {
	l := Livr{titulo: "As cronicas", autor: "Sanderson", valor: 524.4}
	fmt.Println("Dados originais:", l)
	r := descont(&l)
	fmt.Println("Dados com o desconto:", r)
}
