package Expresiones

import (
	"fmt"

	"proyecto1/Enum"
	"proyecto1/Interfaces"
	"proyecto1/Relacionales"
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"
)

type Relacional struct {
	Izquierda      Interfaces.Expresion
	Derecha        Interfaces.Expresion
	TipoRelacional Enum.Tiporelacional
}

func NuevaRelacional(_izquierda Interfaces.Expresion, _derecha Interfaces.Expresion, _tipo Enum.Tiporelacional) Relacional {
	return Relacional{Izquierda: _izquierda, Derecha: _derecha, TipoRelacional: _tipo}
}

func (l Relacional) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo {
	var izquierdaa Simbolo.Simbolo = l.Izquierda.Ejecutar(env, recolector)
	var derechaa Simbolo.Simbolo = l.Derecha.Ejecutar(env, recolector)

	switch l.TipoRelacional {
	case Enum.IGUALDAD:
		result := Relacionales.Igualdad(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Enum.SINTIPO {
			return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	case Enum.MAYOR_QUE:
		result := Relacionales.MayorQue(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Enum.SINTIPO {
			return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

	case Enum.MAYOR_IGUAL:
		result := Relacionales.MayorIgual(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Enum.SINTIPO {
			return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	case Enum.MENOR_QUE:
		result := Relacionales.MenorQue(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Enum.SINTIPO {
			return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	case Enum.MENOR_IGUAL:
		result := Relacionales.MenorIgual(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Enum.SINTIPO {
			return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	case Enum.DESIGUAL:
		result := Relacionales.Desigualdad(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Enum.SINTIPO {
			return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
		//case Enum.NOT:
	}
	var resultado interface{}
	return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: resultado, Mut: false}

}
