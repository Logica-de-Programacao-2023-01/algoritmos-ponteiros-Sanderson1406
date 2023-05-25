package main

import "fmt"

type Conta struct {
	saldo   float64
	titular string
}

func addValo(c *Conta, n float64) Conta {
	c.saldo += n
	return Conta{
		titular: c.titular,
		saldo:   c.saldo,
	}
}

func main() {
	var n float64
	c := Conta{saldo: 2514.4, titular: "João"}
	fmt.Println("Conta antes:", c)
	fmt.Println("Valor a adicionar: ")
	fmt.Scanln(&n)
	r := addValo(&c, n)
	fmt.Println("Conta depois:", r)
}
