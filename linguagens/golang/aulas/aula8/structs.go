// structs
/*  forma de criar sua propria estrutura de dados
personalizar de acordo com a sua necessidade
podemos usar vários tipos diferentes*/

package main

import "fmt"

//type <nome da estrutura> struct {<campos>}
type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	AnoNasc   int
}

func main() {
	fmt.Println(Pessoa{"laysa", "siqueira", 12, 05})
	fmt.Println(Pessoa{Nome: "laysa", Sobrenome: "siqueira", Idade: 12, AnoNasc: 05})
	fmt.Println(Pessoa{Nome: "rebeca"})

	p1 := Pessoa{Nome: "carol", Idade: 3}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Sobrenome)

	p1.Idade = 8
	fmt.Println(p1.Idade)

}
