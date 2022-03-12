package Operaciones

import (
	"math"
	"proyecto1/Interfaces"
)

func Pow(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Interfaces.INTEGER {
		if derecha.Tipo == Interfaces.INTEGER {
			return powintint(izquierda, derecha)
		}
	}
	return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
}

func powintint(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", math.Pow(float64(izquierda.Valor.(int)), float64(derecha.Valor.(int))), false, Interfaces.INTEGER}
}
