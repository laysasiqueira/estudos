package main

import (
	"fmt"
)

func main() {

	//%T significa o tipo da variavel
	//%v significa o valor armazenado

	// boolean (true/false)
	fmt.Printf("type: %T -value: %v\n", true, true)

	// str - sequência de bytes
	fmt.Printf("type: %T -value: %v\n", "laysa", "laysa")
	fmt.Printf("type: %T -value: %v\n", "1", "1")

	// int
	fmt.Printf("type: %T -value: %v\n", 1, 1)

	// float (decimal)
	fmt.Printf("type: %T -value: %v\n", 1.244, 1.244)

}
