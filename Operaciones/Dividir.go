package Operaciones

import (
	"proyecto1/Enum"
	"proyecto1/Simbolo"
)

func Dividir(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Enum.INTEGER {
		if derecha.Tipo == Enum.INTEGER {
			return dividirintint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Enum.FLOAT {
		if derecha.Tipo == Enum.FLOAT {
			return dividirfloatfloat(izquierda, derecha)
		}
	}
	return Simbolo.Simbolo{"", "", false, Enum.ERRORTIPOEXPRESION}
}

func dividirintint(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	if derecha.Valor.(int) >= 0 {
		return Simbolo.Simbolo{"", izquierda.Valor.(int) / derecha.Valor.(int), false, Enum.INTEGER}
	} else {
		return Simbolo.Simbolo{"", "", false, Enum.ERROREXPRESION}
	}
}

func dividirfloatfloat(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	if derecha.Valor.(int) >= 0 {
		return Simbolo.Simbolo{"", izquierda.Valor.(float64) / derecha.Valor.(float64), false, Enum.FLOAT}
	} else {
		return Simbolo.Simbolo{"", "", false, Enum.ERROREXPRESION}
	}
}
