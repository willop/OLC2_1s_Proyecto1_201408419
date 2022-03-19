package Expresiones

import (
	"fmt"
	"math"

	//"proyecto1/Expresion"

	"proyecto1/Enum"
	"proyecto1/Interfaces"
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"
)

type Asi64 struct {
	Expresion Interfaces.Expresion
}

func NewAsi64(_expresion Interfaces.Expresion) Asi64 {
	return Asi64{_expresion}
}

func (a Asi64) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo {
	fmt.Println("Esta dentro de Ejecutar asi64")
	valor := a.Expresion.Ejecutar(env, recolector)
	if valor.Tipo == Enum.FLOAT {
		return Simbolo.Simbolo{Id: "", Tipo: Enum.INTEGER, Valor: int(math.Trunc(valor.Valor.(float64))), Mut: false}
	}
	return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.ERROREXPRESION, Valor: "", Mut: false}
}
