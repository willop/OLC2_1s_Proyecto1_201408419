package Expresion

import (
	"fmt"
	"proyecto1/Estructura"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
)

type Accesovariable struct {
	id string
}

func NewAccesovariable(id string) Accesovariable {
	return Accesovariable{id}
}

func (a Accesovariable) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {
	ax := env.(Estructura.Entorno).GetVariable(a.id)

	if ax.Tipo == Interfaces.SINTIPO {
		fmt.Println("No existe la variable")
		return Interfaces.Simbolo{"", "error", false, Interfaces.SINTIPO}
	}
	return ax
}
