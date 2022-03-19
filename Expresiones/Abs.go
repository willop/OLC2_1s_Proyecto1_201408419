package Expresiones

import (
	"fmt"
	"math"
	"proyecto1/Enum"
	"proyecto1/Interfaces"
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"
)

type Fabs struct {
	Expresion Interfaces.Expresion
}

func Newabs(_expresion Interfaces.Expresion) Fabs {
	return Fabs{Expresion: _expresion}
}

func (f Fabs) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo {
	fmt.Println("Esta en ejecutar de abs")

	resultado := f.Expresion.Ejecutar(env, recolector)
	if resultado.Tipo == Enum.INTEGER {
		temp := float64(resultado.Valor.(int))
		temp2 := int(math.Abs(temp))
		resultado.Valor = temp2
		return resultado
	} else if resultado.Tipo == Enum.FLOAT {
		resultado.Valor = math.Abs(resultado.Valor.(float64))
		return resultado
	}
	return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.ERROREXPRESION, Valor: resultado, Mut: false}
}
