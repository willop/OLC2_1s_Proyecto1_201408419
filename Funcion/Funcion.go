package Funcion

import (
	//"fmt"
	"proyecto1/Enum"
	"proyecto1/Estructura"
	"proyecto1/Utilitario"

	arrayList "github.com/colegno/arraylist"
)

type Funcion struct {
	Id            string
	Tipo          Enum.Tipoexpresion
	Parametros    *arrayList.List
	Instrucciones *arrayList.List
}

func NewFuncion(_id string, _tipo Enum.Tipoexpresion, _parametros *arrayList.List, _instrucciones *arrayList.List) Funcion {
	fun := Funcion{_id, _tipo, _parametros, _instrucciones}
	return fun
}

func (f Funcion) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {
	env.(Estructura.Entorno).GuardarFuncion(f, f.Id)
	return nil
}
