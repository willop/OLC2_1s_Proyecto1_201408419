package Expresiones

import (
	"fmt"

	//"proyecto1/Expresion"

	"proyecto1/Enum"
	"proyecto1/Interfaces"
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"
)

type Asf64 struct {
	Expresion Interfaces.Expresion
}

func NewAsf64(_expresion Interfaces.Expresion) Asf64 {
	return Asf64{_expresion}
}

func (a Asf64) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo {
	fmt.Println("Esta dentro de Ejecutar asf64")
	valor := a.Expresion.Ejecutar(env, recolector)
	if valor.Tipo == Enum.INTEGER {
		return Simbolo.Simbolo{Id: "", Tipo: Enum.FLOAT, Valor: float64(valor.Valor.(int)), Mut: false}
	}
	return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.ERROREXPRESION, Valor: "", Mut: false}
}
