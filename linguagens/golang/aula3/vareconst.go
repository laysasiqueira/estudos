package main

import "fmt"

func main() {
	//variaveis: uma variável é um objeto que pode conter e representar um valor ou expressão.
	//É essencialmente um local nomeado na memória usado para armazenar dados
	//​​são classificadas em dois tipos: globais e locais

	//var + nome_da_variavel + tipo

	var nome string
	nome = "josé"
	fmt.Println(nome)

	nome = "tereza"
	fmt.Println(nome)

	var idade int
	idade = 4
	fmt.Println(idade)

	var a = "tereza"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	f := "amor"
	fmt.Println(f)

	const idadeTeresa = 4
	fmt.Println(idadeTeresa)
}
