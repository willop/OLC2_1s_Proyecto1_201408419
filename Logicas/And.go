package Logicas

import (
	"fmt"
	"proyecto1/Enum"
	"proyecto1/Simbolo"
)

func And(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	if izquierda.Tipo == Enum.BOOLEAN {
		if derecha.Tipo == Enum.BOOLEAN {
			fmt.Println("Entra al fonodo del if de and")
			fmt.Println(andboolbool(izquierda, derecha))
			return andboolbool(izquierda, derecha)
		}
	}
	return Simbolo.Simbolo{"", "", false, Enum.SINTIPO}
}

func andboolbool(izquierda Simbolo.Simbolo, derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", izquierda.Valor.(bool) && derecha.Valor.(bool), false, Enum.BOOLEAN}
}
