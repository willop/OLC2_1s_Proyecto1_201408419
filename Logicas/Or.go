package Logicas

import (
	"proyecto1/Interfaces"
)

func Or(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	if izquierda.Tipo == Interfaces.BOOLEAN {
		if derecha.Tipo == Interfaces.BOOLEAN {
			return orboolbool(izquierda, derecha)
		}
	}
	return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
}

func orboolbool(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(bool) || derecha.Valor.(bool), false, Interfaces.BOOLEAN}
}
