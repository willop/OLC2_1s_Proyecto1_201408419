package Relacionales

import (
	"proyecto1/Interfaces"
)

func MayorQue(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Interfaces.INTEGER {
		if derecha.Tipo == Interfaces.INTEGER {
			return mayorqueintint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Interfaces.FLOAT {
		if derecha.Tipo == Interfaces.FLOAT {
			return mayorquefloatfloat(izquierda, derecha)
		}
	} else if izquierda.Tipo == Interfaces.STR {
		if derecha.Tipo == Interfaces.STR {
			return mayorquestringstring(izquierda, derecha)
		}
	}
	return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
}

func mayorqueintint(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(int) > derecha.Valor.(int), false, Interfaces.BOOLEAN}
}

func mayorquefloatfloat(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(float64) > derecha.Valor.(float64), false, Interfaces.BOOLEAN}
}

func mayorquestringstring(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(string) > derecha.Valor.(string), false, Interfaces.BOOLEAN}
}
