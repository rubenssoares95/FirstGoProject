package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
	cpf   int
}

type funcionario struct {
	pessoa
	matricula    string
	cargo        string
	dataAdmissao uint16
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", 21, 42938457393}
	fmt.Println(p1)

	f1 := funcionario{p1, "1234", "gerente", 01012}
	fmt.Println(f1.nome)
	fmt.Println(f1.idade)
	fmt.Println(f1.cpf)

}
