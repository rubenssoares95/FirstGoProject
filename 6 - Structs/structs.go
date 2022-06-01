package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Rubens"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Tres", 120}

	usuario2 := usuario{"Renan", 27, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 29}
	fmt.Println(usuario3)
}
