package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int

	array1[0] = 01
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6} //DEIXA COM O TAMANHO BASEADO NA QUANTIDADE DE VALORES PASSADO NA DECLARAÇÃO
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println(slice)

	slice[1] = 25
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice)) //TYPEOF DEVOLVE O TIPO DE UMA VARIÁVEL
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 20)
	fmt.Println(slice) //ADICIONA VALOR NUM SLICE

	slice2 := array2[0:3] //CRIA UM SLICE COM UM PEDAÇO DO ARRAY, NO CASO DO INDICE 0 AO 3 (não inclui o 3)
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	//Arrays Internos

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
