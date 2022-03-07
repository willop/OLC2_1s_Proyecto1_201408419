package Instruccion

import (
	"fmt"
	"proyecto1/Interfaces"

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

func (p Print) Ejecutar(env interface{}) interface{} {
	resultado := p.Expresion.Ejecutar(env) //ejecuto la expresion la cual me retorna el simbolo con id,valor,mut,tipo
	fmt.Println("Luego de ejecutar ", resultado)
	fmt.Println("Consolav: ", resultado.Valor)
	return resultado.Valor
}

func nousar() {
	list := arrayList.New()
	list.Add("hola mundo")
	fmt.Println(list)
}
