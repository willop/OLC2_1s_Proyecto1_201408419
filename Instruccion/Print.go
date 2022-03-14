package Instruccion

import (
	"fmt"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
	"reflect"
	"strconv"

	arrayList "github.com/colegno/arraylist"
)

type Print struct {
	Expresion Interfaces.Expresion
	//linea int
	//columna int
}

func NuevoPrint(_exp Interfaces.Expresion) Print {
	valor := _exp
	fmt.Println("Consolav: ", valor)
	return Print{_exp}
}

func (p Print) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {

	resultado := p.Expresion.Ejecutar(env, recolector) //ejecuto la expresion la cual me retorna el simbolo con id,valor,mut,tipo
	fmt.Println("Luego de ejecutar ", resultado)
	fmt.Println("Consolav: ", resultado.Valor)
	var st int = 50
	if reflect.TypeOf(resultado.Valor) == reflect.TypeOf(st) {
		recolector.Consolavirtual.Add(strconv.Itoa(resultado.Valor.(int)))
	} else {
		recolector.Consolavirtual.Add(resultado.Valor)
	}

	return resultado.Valor
}

func nousar() {
	list := arrayList.New()
	list.Add("hola mundo")
	fmt.Println(list)
}
