package Expresiones

import (
	"fmt"

	"proyecto1/Enum"
	"proyecto1/Interfaces"
	"proyecto1/Logicas"
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"
)

type Logica struct {
	Izquierda  Interfaces.Expresion
	Derecha    Interfaces.Expresion
	TipoLogico Enum.Tipologico
}

func NuevaLogica(_izquierda Interfaces.Expresion, _derecha Interfaces.Expresion, _tipo Enum.Tipologico) Logica {
	return Logica{Izquierda: _izquierda, Derecha: _derecha, TipoLogico: _tipo}
}

func (l Logica) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo {
	var izquierdaa Simbolo.Simbolo = l.Izquierda.Ejecutar(env, recolector)
	var derechaa Simbolo.Simbolo = l.Derecha.Ejecutar(env, recolector)

	switch l.TipoLogico {
	case Enum.AND:
		result := Logicas.And(izquierdaa, derechaa)
		fmt.Println("resultado: logico ", result)
		if result.Tipo == Enum.SINTIPO {
			return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	case Enum.OR:
		result := Logicas.Or(izquierdaa, derechaa)
		fmt.Println("resultado: logico ", result)
		if result.Tipo == Enum.SINTIPO {
			return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

	case Enum.NOT:
		result := Logicas.Not(derechaa)
		fmt.Println("resultado: logico ", result)
		if result.Tipo == Enum.SINTIPO {
			return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

		//case Interfaces.NOT:
	}
	var resultado interface{}
	return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: resultado, Mut: false}

}
