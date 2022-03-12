package Expresiones

import (
	"fmt"

	"proyecto1/Interfaces"
	"proyecto1/Relacionales"
	"proyecto1/Utilitario"
)

type Relacional struct {
	Izquierda      Interfaces.Expresion
	Derecha        Interfaces.Expresion
	TipoRelacional Interfaces.Tiporelacional
}

func NuevaRelacional(_izquierda Interfaces.Expresion, _derecha Interfaces.Expresion, _tipo Interfaces.Tiporelacional) Relacional {
	return Relacional{Izquierda: _izquierda, Derecha: _derecha, TipoRelacional: _tipo}
}

func (l Relacional) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Interfaces.Simbolo {
	var izquierdaa Interfaces.Simbolo = l.Izquierda.Ejecutar(env, recolector)
	var derechaa Interfaces.Simbolo = l.Derecha.Ejecutar(env, recolector)

	switch l.TipoRelacional {
	case Interfaces.IGUALDAD:
		result := Relacionales.Igualdad(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	case Interfaces.MAYOR_QUE:
		result := Relacionales.MayorQue(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

	case Interfaces.MAYOR_IGUAL:
		result := Relacionales.MayorIgual(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	case Interfaces.MENOR_QUE:
		result := Relacionales.MenorQue(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	case Interfaces.MENOR_IGUAL:
		result := Relacionales.MenorIgual(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	case Interfaces.DESIGUAL:
		result := Relacionales.Desigualdad(izquierdaa, derechaa)
		fmt.Println("resultado: relacional ", result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
		//case Interfaces.NOT:
	}
	var resultado interface{}
	return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: resultado, Mut: false}

}
