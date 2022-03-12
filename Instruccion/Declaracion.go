package Instruccion

import (
	"fmt"

	"proyecto1/Estructura"
	"proyecto1/Excepciones"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
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
	var env interface{}
	var recolector *Utilitario.Recolector

	simb := _value.Ejecutar(env, recolector)

	if _tipo == Interfaces.SINTIPO {
		//evaluar que es la entrada
		fmt.Println("/n/n/n******* entra a sin tipo con tiponuevade: ", simb.Valor, "**********/n/n/n")
		switch simb.Valor.(type) {
		case string:
			nuevdadec = Declaracion{_id, Interfaces.STRING, _value, _ismutable, _isArray, _isStruct}
			fmt.Println("Entro al case String")
		case int:
			nuevdadec = Declaracion{_id, Interfaces.INTEGER, _value, _ismutable, _isArray, _isStruct}
			fmt.Println("Entro al case int")
		case bool:
			nuevdadec = Declaracion{_id, Interfaces.BOOLEAN, _value, _ismutable, _isArray, _isStruct}
			fmt.Println("Entro al case bool")
		case float64:
			nuevdadec = Declaracion{_id, Interfaces.FLOAT, _value, _ismutable, _isArray, _isStruct}
			fmt.Println("Entro al case float64")
		}
	}

	fmt.Println("Se realizo una declaracion de una varialbe:\n", nuevdadec)

	return nuevdadec
}

func (_dec Declaracion) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {
	fmt.Println("****** Imprimiendo declaracion", _dec)

	var resultado Interfaces.Simbolo = _dec.Expresion.Ejecutar(env, recolector)

	if _dec.Tipo == Interfaces.SINTIPO {
		fmt.Println("Entro al else de error en ejecucion")
		recolector.ListaErrores.Add(Excepciones.ErrorGeneral("Error en la expresion de la variable", env))
	} else if resultado.Tipo == _dec.Tipo {
		//si es del mismo tipo se guarda en el entorno
		fmt.Println("Se accede al ejecutar de la declaracion, y se guarda en el entorno")
		fmt.Println("Se guarda: ", _dec.Id, resultado.Valor, _dec.IsMutable, _dec.Tipo)
		env.(Estructura.Entorno).GuardarSimbolo(_dec.Id, resultado.Valor, _dec.IsMutable, _dec.Tipo)
	} else {
		fmt.Println("Entro al else de no compatible")
		recolector.ListaErrores.Add(Excepciones.ErrorTipoIncorrecto("El tipo no coincide", env))
	}

	return resultado.Valor
}
