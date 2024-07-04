package main

import "fmt"

// loops
//laçõs de repetição
//repetir tarefas

func main() {

	sum := 0

	//i++ -> soma 1
	//i-- -> subtrai 1
	for i := 0; i < 10; i++ { //loop finito
		// o "i" é um indice, a posição inicial
		fmt.Println("sum:", sum)
		fmt.Println("indice", i)
		sum += i
		//é a mesma coisa que: sum = sum + i
		// sum -= i -> sum = sum -i

		//é como se ao final do loop fosse adicionado 1 ao indice i
		//i++
		//i = i + 1

		//obs: váriaveis criadas dentro de um ciclo for só existem dentro do ciclo
	}

	//loop infinito
	for {
		fmt.Println("loop infinito")
		//time.Sleep(2 * time.Second)
	}

	//for range

	frutas := []string{"laranja", "acerola", "pitomba", "jaca", "pitanga"}
	for _, fruta := range frutas { // o "_" é o indice
		fmt.Println("fruta:", fruta)
	}
}
