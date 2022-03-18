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
	return Declaracion{_id, _tipo, _value, _ismutable, _isArray, _isStruct}
}

func (_dec Declaracion) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {

	if _dec.IsArray {
		fmt.Println("Es un array con los siguientes valores")
		fmt.Println(_dec)
		fmt.Println("Termina")
		return nil
	} else {
		simb := _dec.Expresion.Ejecutar(env, recolector)

		if _dec.Tipo == Interfaces.SINTIPO {
			//evaluar que es la entrada
			fmt.Println("/n/n/n******* entra a sin tipo con tiponuevade: ", simb.Valor, "**********/n/n/n")
			switch simb.Valor.(type) {
			case string:
				_dec.Tipo = Interfaces.STR
				fmt.Println("Entro al case String")
			case int:
				_dec.Tipo = Interfaces.INTEGER
				fmt.Println("Entro al case int")
			case bool:
				_dec.Tipo = Interfaces.BOOLEAN
				fmt.Println("Entro al case bool")
			case float64:
				_dec.Tipo = Interfaces.FLOAT
				fmt.Println("Entro al case float64")
			}
			env.(Estructura.Entorno).GuardarSimbolo(_dec.Id, simb.Valor, _dec.IsMutable, _dec.Tipo)
			return simb.Valor
		} else {
			fmt.Println("****** Imprimiendo declaracion", _dec)
			if simb.Tipo == _dec.Tipo {
				//si es del mismo tipo se guarda en el entorno
				fmt.Println("Se guarda en el entorno: ", _dec.Id, simb.Valor, _dec.IsMutable, _dec.Tipo)
				env.(Estructura.Entorno).GuardarSimbolo(_dec.Id, simb.Valor, _dec.IsMutable, _dec.Tipo)
				return simb.Valor
			} else if _dec.Tipo == Interfaces.ERROREXPRESION {
				fmt.Println("Entro al else de error en ejecucion")
				recolector.ListaErrores.Add(Excepciones.ErrorGeneral("Error en la expresion de la variable", env))
				return simb.Valor
			} else {
				fmt.Println("Entro al else de no compatible")
				recolector.ListaErrores.Add(Excepciones.ErrorTipoIncorrecto("El tipo no coincide", env))
				return simb.Valor
			}
		}
	}
}
