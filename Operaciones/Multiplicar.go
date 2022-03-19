package Operaciones

import (
	"proyecto1/Enum"
	"proyecto1/Simbolo"
)

func Multiplicar(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Enum.INTEGER {
		if derecha.Tipo == Enum.INTEGER {
			return multiintint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Enum.FLOAT {
		if derecha.Tipo == Enum.FLOAT {
			return multifloatfloat(izquierda, derecha)
		}
	}
	return Simbolo.Simbolo{"", "", false, Enum.ERRORTIPOEXPRESION}
}

func multiintint(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(int) * derecha.Valor.(int), false, Enum.INTEGER}
}

func multifloatfloat(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", float64(izquierda.Valor.(float64) * derecha.Valor.(float64)), false, Enum.FLOAT}
}
