package Logicas

import (
	"fmt"
	"proyecto1/Interfaces"
)

func Not(derecha Interfaces.Simbolo) Interfaces.Simbolo {

	if derecha.Tipo == Interfaces.BOOLEAN {
		fmt.Println("Entra al fonodo del if de and")
		fmt.Println(notbool(derecha))
		return notbool(derecha)
	}

	return Interfaces.Simbolo{"", "", false, Interfaces.SINTIPO}
}

func notbool(derecha Interfaces.Simbolo) Interfaces.Simbolo {
	return Interfaces.Simbolo{"", !derecha.Valor.(bool), false, Interfaces.BOOLEAN}
}
