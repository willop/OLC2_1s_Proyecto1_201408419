package Logicas

import (
	"fmt"
	"proyecto1/Interfaces"
)

func And(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	if izquierda.Tipo == Interfaces.BOOLEAN {
		if derecha.Tipo == Interfaces.BOOLEAN {
			fmt.Println("Entra al fonodo del if de and")
			fmt.Println(andboolbool(izquierda, derecha))
			return andboolbool(izquierda, derecha)
		}
	}
	return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
}

func andboolbool(izquierda Interfaces.Simbolo, derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", izquierda.Valor.(bool) && derecha.Valor.(bool), false, Interfaces.BOOLEAN}
}
