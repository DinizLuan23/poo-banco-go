package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorBoleto float64){
	conta.Sacar(valorBoleto)
}

func main() {
	clienteLuan := clientes.Titular{"Luan Diniz", "11790368910", "TI"}

	contaLuan := contas.ContaCorrente{
		Titular: clienteLuan,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}

	contaLuan.Depositar(1000)
	PagarBoleto(&contaLuan, 400)

	fmt.Println(contaLuan.ObterSaldo())
}
