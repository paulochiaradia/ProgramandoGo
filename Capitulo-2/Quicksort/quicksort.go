package main

import (
	"fmt"
	"os"
	"strconv"
)

func quicksort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}
	n := make([]int, len(numeros))
	copy(n, numeros)
	indicePivo := len(n) / 2
	pivo := n[indicePivo]
	n = append(n[:indicePivo], n[indicePivo+1:]...)
	menores, maiores := particionar(n, pivo)
	return append(
		append(quicksort(menores), pivo), quicksort(maiores)...)
}

func particionar(
	numeros []int,
	pivo int) (menores []int, maiores []int) {
	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}
	return menores, maiores
}
func main() {
	//atribui os valores lidos pela linha de comando a variável entrada
	entrada := os.Args[1:]
	//valor é lido no formato ‘string’
	/*	for _, k := range entrada {
		fmt.Printf("%s e %T\n", k, k)
	}*/
	//cria um slice de inteiros do tamanho da entrada, não tem nenhum valor ainda
	numeros := make([]int, len(entrada))
	//‘loop’ para converter os dados do formato de ‘string’ para o formato int com a funcao strconv.Atoi
	for i, v := range entrada {
		//atribui o valor de v no range de entrada para a variavel numero e converte para inteiro
		numero, err := strconv.Atoi(v)
		//checa possíveis erros
		if err != nil {
			fmt.Printf("%s nao e um numero valido", v)
			os.Exit(1)
		}
		//atribui o número convertido para o indice da slice numeros criada
		numeros[i] = numero
	}
	/*	for _, l := range numeros {
		fmt.Printf("%d e %T\n", l, l)
	}*/
	fmt.Println(quicksort(numeros))
}
