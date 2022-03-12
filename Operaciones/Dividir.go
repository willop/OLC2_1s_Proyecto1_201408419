package Operaciones

import (
	"proyecto1/Interfaces"
)

func Dividir(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Interfaces.INTEGER {
		if derecha.Tipo == Interfaces.INTEGER {
			return dividirintint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Interfaces.FLOAT {
		if derecha.Tipo == Interfaces.FLOAT {
			return dividirfloatfloat(izquierda, derecha)
		}
	}
	return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
}

func dividirintint(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	if derecha.Valor.(int) >= 0 {
		return Interfaces.Simbolo{"", izquierda.Valor.(int) / derecha.Valor.(int), false, Interfaces.INTEGER}
	} else {
		return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
	}
}

func dividirfloatfloat(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	if derecha.Valor.(int) >= 0 {
		return Interfaces.Simbolo{"", izquierda.Valor.(float64) / derecha.Valor.(float64), false, Interfaces.FLOAT}
	} else {
		return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
	}
}
