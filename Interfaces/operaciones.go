package Interfaces

import (
	"fmt"
	"strconv"
)

func OperacionAritmetica(der string, op string, iz string) string {
	dere, err := strconv.Atoi(der)
	izq, err := strconv.Atoi(iz)
	var resultado int
	var retorno string

	if err != nil {
		//no hacer nada solo retornemos null
		return ""
	}
	fmt.Println("El signo de la operacion es: ", op)
	fmt.Println("derecha es: ", der)
	fmt.Println("izquierda es: ", izq)
	fmt.Println()
	fmt.Println()
	switch op {
	case "+":
		resultado = dere + izq
		retorno = strconv.Itoa(resultado)
		fmt.Println("Suma: ", resultado)

	case "-":
		resultado = dere - izq
		retorno = strconv.Itoa(resultado)
		fmt.Println("Resta: ", resultado)

	case "*":
		resultado = dere * izq
		retorno = strconv.Itoa(resultado)
		fmt.Println("Multiplicacion: ", resultado)

	case "/":
		resultado = dere / izq
		retorno = strconv.Itoa(resultado)
		fmt.Println("Division: ", resultado)
	}

	fmt.Println("el retorno es: ", retorno)
	return retorno

}
