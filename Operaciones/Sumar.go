package Operaciones

import (
	"proyecto1/Enum"
	"proyecto1/Simbolo"
)

func Sumar(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Enum.INTEGER {
		if derecha.Tipo == Enum.INTEGER {
			return sumarintint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Enum.FLOAT {
		if derecha.Tipo == Enum.FLOAT {
			return sumarfloatfloat(izquierda, derecha)
		}
	} else if izquierda.Tipo == Enum.STRING {
		if derecha.Tipo == Enum.STRING {
			return sumarstringstring(izquierda, derecha)
		}
	}
	return Simbolo.Simbolo{"", "", false, Enum.ERRORTIPOEXPRESION}
}

func sumarintint(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(int) + derecha.Valor.(int), false, Enum.INTEGER}
}

func sumarfloatfloat(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(float64) + derecha.Valor.(float64), false, Enum.FLOAT}
}

func sumarstringstring(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(string) + derecha.Valor.(string), false, Enum.STR}
}
