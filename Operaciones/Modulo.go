package Operaciones

import (
	"math"
	"proyecto1/Interfaces"
)

func Modulo(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Interfaces.INTEGER {
		if derecha.Tipo == Interfaces.INTEGER {
			return modulointint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Interfaces.FLOAT {
		if derecha.Tipo == Interfaces.FLOAT {
			return modulofloatfloat(izquierda, derecha)
		}
	}
	return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
}

func modulointint(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	if derecha.Valor.(int) != 0 {
		return Interfaces.Simbolo{"", izquierda.Valor.(int) % derecha.Valor.(int), false, Interfaces.INTEGER}
	} else {
		return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
	}
}

func modulofloatfloat(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	if derecha.Valor.(float64) != 0.0 {
		//var res float64 = float64(math.Mod(izquierda.Valor.(float64), derecha.Valor.(float64)))
		return Interfaces.Simbolo{"", math.Mod(izquierda.Valor.(float64), derecha.Valor.(float64)), false, Interfaces.FLOAT}
	} else {
		return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
	}
}
