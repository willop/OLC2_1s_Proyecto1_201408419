package Operaciones

import (
	"math"
	"proyecto1/Enum"
	"proyecto1/Simbolo"
)

func Powf(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Enum.FLOAT {
		if derecha.Tipo == Enum.FLOAT {
			return powffloatfloat(izquierda, derecha)
		}
	}
	return Simbolo.Simbolo{"", "", false, Enum.ERRORTIPOEXPRESION}
}

func powffloatfloat(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", math.Pow(izquierda.Valor.(float64), derecha.Valor.(float64)), false, Enum.FLOAT}
}
