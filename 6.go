package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
}

func liv(l Livro) Livro {
	if l.autor == "Anonimo" {
		l.titulo = "Desconhecido"
	}
	return l
}

func main() {
	l := Livro{titulo: "As cronicas", autor: "Lucas"}
	r := liv(l)
	fmt.Println(r)
}
