//listas
/* 1- array e slices: homogêneos
	todos os elemetos tem o mesmo tipo
	[1, 2, 3, 4, 5] - int
	["laysa", "siqueira"] - string

   2- maps: heterogêneos
    pode misturar tipos
	estrutura - valor
	[key] = value
	chave tem um tipo, e o valor pode ter outro
	map[string]int
	{1: "laysa", 3: "siqueira"} - int, string
	map[string]string
	{"laysa": "siqueira", "josé": "fernando"}
*/

//o indice é sempre o tamanho do array -1. o indice começa sempre no 0

// array

/* tamanho fixo, de zero ou mais elementos do mesmo tipo.
   acessamos os valores com indice: a[0], a[1]...
   funçãos embutida len() retorna o tamanho do array
   por conta do tamanho fixo, não é tão usado, só em casos específicos.
*/

//slice

/*tipo o array, mas sem tamanho fixo
  acessamos os valores com indice: a[0], a[1]...
  função embutida len() retorna o tamanho do array
  função embutida append() usada para adicionar valor
*/

package main

import "fmt"

func main() {

	//array- tamanho fixo
	var array [2]string //aqui é onde eu declaro o tamanho do array
	array[0] = "laysa"
	array[1] = "siqueira"
	fmt.Println(array[1], array[0]) //print da posição
	fmt.Println(array)              //print do array inteiro

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:3]) //da posição 0 (:=até) a 3.
	fmt.Println(numPrimos[:5])  //mostra o que tá antes da posição 5
	fmt.Println(numPrimos[5:])  //mostra o que tá depois da posição 5

	//var slice []string
	// nesse primeiro caso criei o slice sem valores
	slice := make([]string, 5) //determina um valor inicial para ele e depois pode adicionar mais
	slice[0] = "laysa"
	slice[1] = "siqueira"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)
	fmt.Println(len(slice)) //tamanho do slice
	slice[2] = "josé"
	fmt.Println(slice[2])
	fmt.Println(slice[3])

	// nesse segundo caso criei com valore
	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	numPares = append(numPares, 14) //adicionando mais um valor ou mais no slice
	fmt.Println(numPares)

}
