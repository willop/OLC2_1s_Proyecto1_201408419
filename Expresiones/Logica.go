package Expresiones

import (
	"fmt"

	"proyecto1/Interfaces"
	"proyecto1/Logicas"
	"proyecto1/Utilitario"
)

type Logica struct {
	Izquierda  Interfaces.Expresion
	Derecha    Interfaces.Expresion
	TipoLogico Interfaces.Tipologico
}

func NuevaLogica(_izquierda Interfaces.Expresion, _derecha Interfaces.Expresion, _tipo Interfaces.Tipologico) Logica {
	return Logica{Izquierda: _izquierda, Derecha: _derecha, TipoLogico: _tipo}
}

func (l Logica) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Interfaces.Simbolo {
	var izquierdaa Interfaces.Simbolo = l.Izquierda.Ejecutar(env, recolector)
	var derechaa Interfaces.Simbolo = l.Derecha.Ejecutar(env, recolector)

	switch l.TipoLogico {
	case Interfaces.AND:
		result := Logicas.And(izquierdaa, derechaa)
		fmt.Println("resultado: logico ", result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	case Interfaces.OR:
		result := Logicas.Or(izquierdaa, derechaa)
		fmt.Println("resultado: logico ", result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

	case Interfaces.NOT:
		result := Logicas.Not(derechaa)
		fmt.Println("resultado: logico ", result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

		//case Interfaces.NOT:
	}
	var resultado interface{}
	return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: resultado, Mut: false}

}
