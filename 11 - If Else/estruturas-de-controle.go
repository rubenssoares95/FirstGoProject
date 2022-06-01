package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//INICIANDO VARIAVEL NO IF (IF INIT) MAS A VARIÁVEL FICA LIMITADA AO ESCOPO DO IF E DO ELSE
	if outroNumero := 25; outroNumero > 10 {
		fmt.Println("Número é maior que 5")
	} else if outroNumero < 7 {
		fmt.Println("Número é menor que zero")
	} else {
		fmt.Println("Está entre 7 e 10")
	}

}
