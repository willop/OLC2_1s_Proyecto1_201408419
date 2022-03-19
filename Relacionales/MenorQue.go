package Relacionales

import (
	"proyecto1/Enum"
	"proyecto1/Simbolo"
)

func MenorQue(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Enum.INTEGER {
		if derecha.Tipo == Enum.INTEGER {
			return menorintint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Enum.FLOAT {
		if derecha.Tipo == Enum.FLOAT {
			return menorfloatfloat(izquierda, derecha)
		}
	} else if izquierda.Tipo == Enum.STR {
		if derecha.Tipo == Enum.STR {
			return menorstringstring(izquierda, derecha)
		}
	}
	return Simbolo.Simbolo{"", "", false, Enum.SINTIPO}
}

func menorintint(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(int) < derecha.Valor.(int), false, Enum.BOOLEAN}
}

func menorfloatfloat(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(float64) < derecha.Valor.(float64), false, Enum.BOOLEAN}
}

func menorstringstring(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(string) < derecha.Valor.(string), false, Enum.BOOLEAN}
}
