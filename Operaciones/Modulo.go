package Operaciones

import (
	"math"
	"proyecto1/Enum"
	"proyecto1/Simbolo"
)

func Modulo(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Enum.INTEGER {
		if derecha.Tipo == Enum.INTEGER {
			return modulointint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Enum.FLOAT {
		if derecha.Tipo == Enum.FLOAT {
			return modulofloatfloat(izquierda, derecha)
		}
	}
	return Simbolo.Simbolo{"", "", false, Enum.ERRORTIPOEXPRESION}
}

func modulointint(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	if derecha.Valor.(int) != 0 {
		return Simbolo.Simbolo{"", izquierda.Valor.(int) % derecha.Valor.(int), false, Enum.INTEGER}
	} else {
		return Simbolo.Simbolo{"", "", false, Enum.ERRORTIPOEXPRESION}
	}
}

func modulofloatfloat(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	if derecha.Valor.(float64) != 0.0 {
		//var res float64 = float64(math.Mod(izquierda.Valor.(float64), derecha.Valor.(float64)))
		return Simbolo.Simbolo{"", math.Mod(izquierda.Valor.(float64), derecha.Valor.(float64)), false, Enum.FLOAT}
	} else {
		return Simbolo.Simbolo{"", "", false, Enum.ERRORTIPOEXPRESION}
	}
}
