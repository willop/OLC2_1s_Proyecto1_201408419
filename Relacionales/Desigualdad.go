package Relacionales

import (
	"proyecto1/Interfaces"
)

func Desigualdad(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Interfaces.INTEGER {
		if derecha.Tipo == Interfaces.INTEGER {
			return desigualintint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Interfaces.FLOAT {
		if derecha.Tipo == Interfaces.FLOAT {
			return desigualfloatfloat(izquierda, derecha)
		}
	} else if izquierda.Tipo == Interfaces.STR {
		if derecha.Tipo == Interfaces.STR {
			return desigualstringstring(izquierda, derecha)
		}
	}
	return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
}

func desigualintint(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(int) != derecha.Valor.(int), false, Interfaces.BOOLEAN}
}

func desigualfloatfloat(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(float64) != derecha.Valor.(float64), false, Interfaces.BOOLEAN}
}

func desigualstringstring(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(string) != derecha.Valor.(string), false, Interfaces.BOOLEAN}
}
