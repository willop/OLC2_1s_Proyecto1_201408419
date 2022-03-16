package Expresiones

import (
	"fmt"
	"math"

	//"proyecto1/Expresion"

	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
)

type Asi64 struct {
	Expresion Interfaces.Expresion
}

func NewAsi64(_expresion Interfaces.Expresion) Asi64 {
	return Asi64{_expresion}
}

func (a Asi64) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Interfaces.Simbolo {
	fmt.Println("Esta dentro de Ejecutar asi64")
	valor := a.Expresion.Ejecutar(env, recolector)
	if valor.Tipo == Interfaces.FLOAT {
		return Interfaces.Simbolo{Id: "", Tipo: Interfaces.INTEGER, Valor: int(math.Trunc(valor.Valor.(float64))), Mut: false}
	}
	return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.ERROREXPRESION, Valor: "", Mut: false}
}
