package Operaciones

import (
	"math"
	"proyecto1/Enum"
	"proyecto1/Simbolo"
)

func Pow(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Enum.INTEGER {
		if derecha.Tipo == Enum.INTEGER {
			return powintint(izquierda, derecha)
		}
	}
	return Simbolo.Simbolo{"", "", false, Enum.ERRORTIPOEXPRESION}
}

func powintint(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", math.Pow(float64(izquierda.Valor.(int)), float64(derecha.Valor.(int))), false, Enum.INTEGER}
}
