package Expresion

import (
	"fmt"
	"proyecto1/Enum"
	"proyecto1/Interfaces"
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"

	arrayList "github.com/colegno/arraylist"
)

type Vector struct {
	Exp              Interfaces.Expresion
	Exp2             Interfaces.Expresion
	Listaexpresiones *arrayList.List
	TipoVec          Enum.Tipoarray
	//Mutable          bool
}

func NewVector(_exp Interfaces.Expresion, _exp2 Interfaces.Expresion, listaexpresiones *arrayList.List, _tipo Enum.Tipoarray) Vector {
	return Vector{_exp, _exp2, listaexpresiones, _tipo}
}

func (v Vector) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo {
	fmt.Println("Estoy en ejecutar VECTOR")

	var temp *arrayList.List
	temp = arrayList.New()

	if v.TipoVec == Enum.NORMAL {
		temp.Add(v.Exp.(Interfaces.Expresion).Ejecutar(env, recolector))
		for _, s := range v.Listaexpresiones.ToArray() {
			temp.Add(s.(Interfaces.Expresion).Ejecutar(env, recolector))
		}
		//return Interfaces.Simbolo{Id: "", Valor: temp, Tipo: Interfaces.ARRAY}
	} else if v.TipoVec == Enum.MULTIPLE {
		val := v.Exp2.Ejecutar(env, recolector)
		for j := 0; j < val.Valor.(int); j++ {
			temp.Add(v.Exp.Ejecutar(env, recolector))
		}
	} else {
		fmt.Println("Es un Vector nuevo")
	}

	fmt.Println("Esta en ejecutar un VECTOR, el valor es: ", temp)
	return Simbolo.Simbolo{Id: "", Valor: temp, Tipo: Enum.VECTOR}

}
