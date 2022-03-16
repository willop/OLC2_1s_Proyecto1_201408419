package Expresiones

import (
	"fmt"

	//"proyecto1/Expresion"

	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
)

type Asf64 struct {
	Expresion Interfaces.Expresion
}

func NewAsf64(_expresion Interfaces.Expresion) Asf64 {
	return Asf64{_expresion}
}

func (a Asf64) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Interfaces.Simbolo {
	fmt.Println("Esta dentro de Ejecutar asf64")
	valor := a.Expresion.Ejecutar(env, recolector)
	if valor.Tipo == Interfaces.INTEGER {
		return Interfaces.Simbolo{Id: "", Tipo: Interfaces.FLOAT, Valor: float64(valor.Valor.(int)), Mut: false}
	}
	return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.ERROREXPRESION, Valor: "", Mut: false}
}
