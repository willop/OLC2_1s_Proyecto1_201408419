package Operaciones

import (
	"proyecto1/Interfaces"
)

func Restar(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Interfaces.INTEGER {
		if derecha.Tipo == Interfaces.INTEGER {
			return restarintint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Interfaces.FLOAT {
		if derecha.Tipo == Interfaces.FLOAT {
			return restarfloatfloat(izquierda, derecha)
		}
	}
	return Interfaces.Simbolo{"", "", false, Interfaces.ERRORTIPOEXPRESION}
}

func restarintint(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(int) - derecha.Valor.(int), false, Interfaces.INTEGER}
}

func restarfloatfloat(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(float64) - derecha.Valor.(float64), false, Interfaces.FLOAT}
}
