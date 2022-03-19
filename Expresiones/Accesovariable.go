package Expresiones

import (
	"fmt"
	"proyecto1/Enum"
	"proyecto1/Estructura"
	"proyecto1/Simbolo"
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

	if ax.Tipo == Enum.SINTIPO {
		fmt.Println("No existe la variable")
		return Simbolo.Simbolo{"", "error", false, Enum.SINTIPO}
	}
	return ax
}
