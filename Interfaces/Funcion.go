package Interfaces

import (
	//"fmt"
	//"proyecto1/Estructura"
	"proyecto1/Utilitario"
	//"proyecto1/Interfaces"

	arrayList "github.com/colegno/arraylist"
)

type Funcion struct {
	Id            string
	Tipo          Tipoexpresion
	Parametros    *arrayList.List
	Instrucciones *arrayList.List
}

func NewFuncion(_id string, _tipo Tipoexpresion, _parametros *arrayList.List, _instrucciones *arrayList.List) Funcion {
	fun := Funcion{_id, _tipo, _parametros, _instrucciones}
	return fun
}

func (f Funcion) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {
	//env.(Estructura.Entorno).GuardarFuncion(f)
	return nil
}
