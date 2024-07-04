package main

import "fmt"

func main() {

	posicao := 2
	switch posicao {
	case 1:
		fmt.Println("Primeiro")
	case 2:
		fmt.Println("Segundo")
	case 3:
		fmt.Println("Terceiro")
	}

	nome := "josé"
	switch nome {
	case "laysa":
		fmt.Println("é uma pessoa")
	case "francisco":
		fmt.Println("é um gato")
	case "jose":
		fmt.Println("é meu filho")
	}
}
