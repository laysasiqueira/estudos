package main

import "fmt"

func main() {
	fmt.Println(soma(43, 14))

	sub := subtracao(4, 6)
	fmt.Println(sub)

}

func soma(x int, y int) int {
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

//quando a função começa com letra minuscula ela é privada
//quando a função começa com letra maiuscula ela é publica
//para utilizar em putros pacote:mainNomeFuncao(nome)
