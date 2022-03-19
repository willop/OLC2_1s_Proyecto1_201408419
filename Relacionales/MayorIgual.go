package Relacionales

import (
	"proyecto1/Enum"
	"proyecto1/Simbolo"
)

func MayorIgual(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	//izquierda con int y sus derivaciones
	if izquierda.Tipo == Enum.INTEGER {
		if derecha.Tipo == Enum.INTEGER {
			return mayorigualintint(izquierda, derecha)
		}
	} else if izquierda.Tipo == Enum.FLOAT {
		if derecha.Tipo == Enum.FLOAT {
			return mayorigualfloatfloat(izquierda, derecha)
		}
	} else if izquierda.Tipo == Enum.STR {
		if derecha.Tipo == Enum.STR {
			return mayorigualstringstring(izquierda, derecha)
		}
	}
	return Simbolo.Simbolo{"", "", false, Enum.SINTIPO}
}

func mayorigualintint(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(int) >= derecha.Valor.(int), false, Enum.BOOLEAN}
}

func mayorigualfloatfloat(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(float64) >= derecha.Valor.(float64), false, Enum.BOOLEAN}
}

func mayorigualstringstring(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(string) >= derecha.Valor.(string), false, Enum.BOOLEAN}
}
