package Logicas

import (
	"proyecto1/Enum"
	"proyecto1/Simbolo"
)

func Or(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	if izquierda.Tipo == Enum.BOOLEAN {
		if derecha.Tipo == Enum.BOOLEAN {
			return orboolbool(izquierda, derecha)
		}
	}
	return Simbolo.Simbolo{"", "", false, Enum.SINTIPO}
}

func orboolbool(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(bool) || derecha.Valor.(bool), false, Enum.BOOLEAN}
}
