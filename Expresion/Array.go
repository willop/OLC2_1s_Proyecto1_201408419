package Expresion

import (
	"fmt"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"

	arrayList "github.com/colegno/arraylist"
)

type Array struct {
	Exp              Interfaces.Expresion
	Exp2             Interfaces.Expresion
	Listaexpresiones *arrayList.List
	TipoArray        Interfaces.Tipoarray
	//Mutable          bool
}

func NewArray(_exp Interfaces.Expresion, _exp2 Interfaces.Expresion, listaexpresiones *arrayList.List, _tipo Interfaces.Tipoarray) Array {
	return Array{_exp, _exp2, listaexpresiones, _tipo}
}

func (a Array) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Interfaces.Simbolo {
	fmt.Println("Estoy en ejecutar array")

	var temp *arrayList.List
	temp = arrayList.New()

	if a.TipoArray == Interfaces.NORMAL {
		temp.Add(a.Exp.(Interfaces.Expresion).Ejecutar(env, recolector))
		for _, s := range a.Listaexpresiones.ToArray() {
			temp.Add(s.(Interfaces.Expresion).Ejecutar(env, recolector))
		}
		//return Interfaces.Simbolo{Id: "", Valor: temp, Tipo: Interfaces.ARRAY}
	} else {
		val := a.Exp2.Ejecutar(env, recolector)
		for j := 0; j < val.Valor.(int); j++ {
			temp.Add(a.Exp.Ejecutar(env, recolector))
		}
	}

	fmt.Println("Esta en ejecutar un array, el valor es: ", temp)

	return Interfaces.Simbolo{Id: "", Valor: temp, Tipo: Interfaces.ARRAY}
}
