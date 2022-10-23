package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta VerificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoMarcos := contas.ContaCorrente{
		Titular:       clientes.Titular{Nome: "Marcos Paulo", CPF: "12312312311", Profissao: "Desenvolvedor"},
		NumeroAgencia: 1234,
		NumeroConta:   4321,
	}

	contaDaDiana := contas.ContaPoupanca{
		Titular:       clientes.Titular{Nome: "Diana", CPF: "00000012311", Profissao: "Designer de Moda"},
		NumeroAgencia: 1212,
		NumeroConta:   4343,
	}

	contaDaDiana.Depositar(100)
	PagarBoleto(&contaDaDiana, 30)

	contaDoMarcos.Depositar(200)
	PagarBoleto(&contaDoMarcos, 180)

	fmt.Println(contaDoMarcos.Titular.Nome, contaDoMarcos.ObterSaldo())
	fmt.Println(contaDaDiana.Titular.Nome, contaDaDiana.ObterSaldo())

}
