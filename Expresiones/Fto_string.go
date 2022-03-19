package Expresiones

import (
	"fmt"
	"proyecto1/Enum"
	"proyecto1/Interfaces"
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"
)

type Fto_string struct {
	Expresion Interfaces.Expresion
}

func NewFto_string(_expresion Interfaces.Expresion) Fto_string {
	return Fto_string{_expresion}
}

func (f Fto_string) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo {
	fmt.Printf("En ejecutar de to_string()")
	resultado := f.Expresion.Ejecutar(env, recolector)

	if resultado.Tipo == Enum.STRING {
		resultado.Tipo = Enum.STR
		return resultado
	} else {
		//error de tipo
		return Simbolo.Simbolo{Id: "Error_del_to_string()_op", Tipo: Enum.ERROREXPRESION, Valor: resultado, Mut: false}
	}

}
