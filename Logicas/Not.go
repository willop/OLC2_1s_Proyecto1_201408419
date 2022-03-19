package Logicas

import (
	"fmt"
	"proyecto1/Enum"
	"proyecto1/Simbolo"
)

func Not(derecha Simbolo.Simbolo) Simbolo.Simbolo {

	if derecha.Tipo == Enum.BOOLEAN {
		fmt.Println("Entra al fonodo del if de and")
		fmt.Println(notbool(derecha))
		return notbool(derecha)
	}

	return Simbolo.Simbolo{"", "", false, Enum.SINTIPO}
}

func notbool(derecha Simbolo.Simbolo) Simbolo.Simbolo {
	return Simbolo.Simbolo{"", !derecha.Valor.(bool), false, Enum.BOOLEAN}
}
