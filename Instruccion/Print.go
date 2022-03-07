package Instruccion

import (
	"fmt"
	"proyecto1/Interfaces"
)

type Print struct {
	Expresion Interfaces.Expresion
	//linea int
	//columna int
}

func NuevoPrint(_exp Interfaces.Expresion) Print {
	valor := _exp
	fmt.Println("Consolav: ", valor)
	return Print{_exp}
}

func (p Print) Ejecutar(env interface{}) interface{} {
	resultado := p.Expresion.Ejecutar(env)
	fmt.Println("Consolav: ", resultado.Valor)
	return resultado.Valor
}
