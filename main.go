package main

import (
	"fmt"
	"banco/contas"
)

func main() {
	contaLuan := contas.ContaCorrente{ Titular: "Luan Diniz", Saldo: 500 }
	contaLuna := contas.ContaCorrente{ Titular: "Luna Diniz", Saldo: 500 }

	fmt.Println(contaLuan)
	fmt.Println(contaLuna)
}