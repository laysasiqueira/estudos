/*2- maps: heterogêneos
pode misturar tipos
estrutura - valor
[key] = value
chave tem um tipo, e o valor pode ter outro
map[k]v -> k = chave, v = valor


map[string]int
{"laysa": 20, "siqueira": 18}
map[string]string
{"laysa": "siqueira", "siqueira": "laysa"}
*/

package main

import "fmt"

func main() {

	idade := map[string]int{} //começa vazio e depois atribui valores a ele
	idade["laysa"] = 20
	idade["siqueira"] = 18
	fmt.Println(idade)
	fmt.Println(idade["laysa"])
	fmt.Println(idade["siqueira"])

	anoNasc := map[string]int{ //crio direto
		"laysa":    2000,
		"siqueira": 2002,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["laysa"])
	fmt.Println(anoNasc["siqueira"])
	idade["amaral"] = 10 //adicionando mais um elemento ao meu map
	fmt.Println(idade)

}
