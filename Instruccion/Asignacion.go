package Instruccion

import (
	"fmt"
	"proyecto1/Estructura"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
)

type Asignacion struct {
	Expresion Interfaces.Expresion
	Id        string
}

func NewAsignacion(_id string, _expresion Interfaces.Expresion) Asignacion {
	return Asignacion{Id: _id, Expresion: _expresion}
}

func (a Asignacion) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {

	fmt.Println("Ejecuta asignacion")
	variable := env.(Estructura.Entorno).GetVariable(a.Id)
	valor := a.Expresion.Ejecutar(env, recolector)

	fmt.Println("variable: ", variable)
	fmt.Println("valor: ", valor)

	if variable.Tipo != valor.Tipo {
		fmt.Println("No coinciden los tipos")
		//recolector.ListaErrores.Add(Excepciones.ErrorTipoIncorrecto("El nuevo valor no es igual al asignado", env))
		return variable
	} else if variable.Mut == false {
		fmt.Println("La variable no es mutable")
		//recolector.ListaErrores.Add(Excepciones.ErrorTipoIncorrecto("La variable no es mutable", env))
		return variable
	} else {
		return env.(Estructura.Entorno).ActualizarSimbolo(a.Id, valor.Valor, true, variable.Tipo)
	}
}
