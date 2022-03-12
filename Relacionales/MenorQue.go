package Relacionales

import (
	"proyecto1/Interfaces"
)

func MenorQue(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Interfaces.INTEGER {
		if derecha.Tipo == Interfaces.INTEGER {
			return menorintint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Interfaces.FLOAT {
		if derecha.Tipo == Interfaces.FLOAT {
			return menorfloatfloat(izquierda, derecha)
		}
	} else if izquierda.Tipo == Interfaces.STR {
		if derecha.Tipo == Interfaces.STR {
			return menorstringstring(izquierda, derecha)
		}
	}
	return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
}

func menorintint(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(int) < derecha.Valor.(int), false, Interfaces.BOOLEAN}
}

func menorfloatfloat(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(float64) < derecha.Valor.(float64), false, Interfaces.BOOLEAN}
}

func menorstringstring(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(string) < derecha.Valor.(string), false, Interfaces.BOOLEAN}
}
