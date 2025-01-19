package main

import (
	"fmt"

	"github.com/stevengm45/go_zero/variables"
)

func main() {
	fmt.Println("Hello World")
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1024)
	fmt.Println(estado)
	fmt.Println(texto)
}
