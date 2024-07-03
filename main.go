package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque <= c.saldo && valorSaque > 0

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	}else{
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso", c.saldo
	}else {
		return "Valor do depósito menor que zero", c.saldo
	}
}

func main() {
	contaLuan := ContaCorrente{ titular: "Luan Diniz", saldo: 500 }
	fmt.Println(contaLuan)

	fmt.Println(contaLuan.Sacar(800))
	fmt.Println(contaLuan.saldo)
}