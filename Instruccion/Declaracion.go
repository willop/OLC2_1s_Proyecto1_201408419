package Instruccion

import (
	"fmt"
	"proyecto1/Interfaces"
)

type Declaracion struct {
	Id        string
	Tipo      Interfaces.Tipoexpresion
	Expresion Interfaces.Expresion
	IsArray   bool
	IsStruct  bool
	IsMutable bool
}

// para declarar una nueva variable se recolecta toda la informacion para crear una nueva variable
func NuevaDeclaracion(_id string, _tipo Interfaces.Tipoexpresion, _value Interfaces.Expresion, _ismutable bool, _isArray bool, _isStruct bool) Declaracion {
	nuevdadec := Declaracion{_id, _tipo, _value, _ismutable, _isArray, _isStruct}
	fmt.Println("Se realizo una declaracion de una varialbe:\n", nuevdadec)

	return nuevdadec
}

/*
func (_dec Declaracion) Ejecutar(env interface{}) interface{} {
	var resultado Interfaces.Simbolo
	resultado = _dec.Expresion.Ejecutar(env)


}
*/
