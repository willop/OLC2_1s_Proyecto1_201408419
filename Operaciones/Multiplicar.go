package Operaciones

import (
	"proyecto1/Interfaces"
)

func Multiplicar(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Interfaces.INTEGER {
		if derecha.Tipo == Interfaces.INTEGER {
			return multiintint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Interfaces.FLOAT {
		if derecha.Tipo == Interfaces.FLOAT {
			return multifloatfloat(izquierda, derecha)
		}
	}
	return Interfaces.Simbolo{"", "", false, Interfaces.ERRORTIPOEXPRESION}
}

func multiintint(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(int) * derecha.Valor.(int), false, Interfaces.INTEGER}
}

func multifloatfloat(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", float64(izquierda.Valor.(float64) * derecha.Valor.(float64)), false, Interfaces.FLOAT}
}
