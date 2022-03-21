package Instruccion

import (
	"fmt"
	//"proyecto1/Expresiones"
	"proyecto1/Estructura"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
	"reflect"
)

type AsignacionArray struct {
	Id  Interfaces.Expresion
	Exp Interfaces.Expresion
}

func NewAsignacionArray(_id Interfaces.Expresion, _exp Interfaces.Expresion) AsignacionArray {
	return AsignacionArray{_id, _exp}
}

func (a AsignacionArray) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {
	fmt.Println("En ejecutar Asignacion array")

	fmt.Println("a.id **************", a.Id)
	fmt.Println("a.exp **************", a.Exp)
	valor := a.Id.Ejecutar(env, recolector)
	result := a.Exp.Ejecutar(env, recolector)
	fmt.Println("Imprimiendo el valor: ", valor)
	fmt.Println("El tipo de valor: ", reflect.TypeOf(valor))
	fmt.Println("Result es igual: ", result)
	fmt.Println("Valor de valor.Valor: ", valor.Valor)
	fmt.Println("Valor de valor: ", valor.Id)
	valor.Valor = result
	fmt.Println("nuevo valor", valor)
	env.(Estructura.Entorno).ActualizarSimbolo(valor.Id.(string), valor, true, valor.Tipo)
	return nil
}
