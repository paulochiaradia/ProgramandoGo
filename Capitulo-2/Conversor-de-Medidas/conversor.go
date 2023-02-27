package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Verifica se existe o número minimo de argumentos para o programa funcionar
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}
	//Atribui o valor do último elemento do slice de entrada de argumentos
	unidadeOrigem := os.Args[len(os.Args)-1]
	//Atribui o valor dos elementos entre o primeiro e ultimo elemento nao exclusivo
	valoresOrigem := os.Args[1 : len(os.Args)-1]

	//Declarando a variable de unidade destino com var, pois nao se sabe o seu valor ainda
	var unidadeDestino string
	//Bloco de controle para sabermos quais são as conversões a serem feitas
	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOrigem == "quilometros" {
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!", unidadeOrigem)
		os.Exit(1)
	}

	//Fazendo um loop nos valores de origem utilizando range
	//convertendo os valores de origem para números float com o strconv
	for i, v := range valoresOrigem {
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s na posição %d não é um número válido!\n", v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n",
			valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}
}