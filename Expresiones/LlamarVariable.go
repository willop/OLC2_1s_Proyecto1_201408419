package Expresiones

import (
	"proyecto1/Estructura"
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"
)

type Llamadavariable struct {
	Id string
}

func NewLlamarvariable(id string) Llamadavariable {
	exp := Llamadavariable{Id: id}
	return exp
}

func (p Llamadavariable) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo {

	result := env.(Estructura.Entorno).GetVariable(p.Id)
	return result
}
