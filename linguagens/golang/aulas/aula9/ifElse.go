package main

import "fmt"

//if - else
//se - se não

//sequência de passos

func main() {

	a := 1
	// if (condição) { <ação>}
	if a == 1 { //true
		fmt.Println("valor é igual a 1")
	} else {
		fmt.Println("valor não é igual a 1")
	}

	b := 1
	if b == 3 { //true
		fmt.Println("valor é igual a 1")
	} else if b == 2 {
		fmt.Println("valor é igual a 2")
	} else {
		fmt.Println("valor é diferentte de 1 e de 2")
	}

	fmt.Println(1 + 1)
	fmt.Println(2 - 1)
	fmt.Println(2 * 2)
	fmt.Println(2 / 2)
	fmt.Println(2 % 1)

	//testando a paridade
	numero := 7
	if numero%2 == 0 {
		fmt.Printf("%d é par", numero) // o %d substitui o valor da variavel número. quando usado com o printf
	} else {
		fmt.Printf("%d é impar", numero)
	}

	// ----operações----
	// == igual
	// != diferente
	// > maior que
	// < menor que
	// >= maior ou igual
	// <= menor ou igual
	// && e
	// || ou
	// ! negação
	// soma: 1 + 1
	// subtração: 1 - 1
	// multiplicação: 1 * 1
	// divisão: 1 / 1
	// resto da divisão: 1 % 1 (resto da divisão por 1)

}
