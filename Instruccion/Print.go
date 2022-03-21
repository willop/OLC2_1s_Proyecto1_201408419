package Instruccion

import (
	"fmt"
	"proyecto1/Enum"
	"proyecto1/Interfaces"
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"
	"reflect"
	"strings"

	arrayList "github.com/colegno/arraylist"
)

type Print struct {
	Expresion  Interfaces.Expresion
	Bloqueinst *arrayList.List
	//linea int
	//columna int
}

func NuevoPrint(_exp Interfaces.Expresion, _bloqueinst *arrayList.List) Print {
	return Print{_exp, _bloqueinst}
}

func (p Print) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {

	resultado := p.Expresion.Ejecutar(env, recolector) //ejecuto la expresion la cual me retorna el simbolo con id,valor,mut,tipo
	//si viene solo para imprimir texto sin parametros
	if p.Bloqueinst == nil {
		fmt.Println("Luego de ejecutar ", resultado)
		fmt.Println("Consolav: ", resultado.Valor)
		tmpDato := fmt.Sprintf("%v", resultado.Valor)
		recolector.Consolavirtual.Add(tmpDato + "\n")
		return resultado.Valor
	} else {
		//si vienen parametros
		var arrtemp *arrayList.List
		st := resultado.Valor
		for j := 0; j < p.Bloqueinst.Len(); j++ {
			instr := p.Bloqueinst.GetValue(j).(Interfaces.Expresion)
			val := instr.Ejecutar(env, recolector)
			//verifico si ese parametro es un arreglo
			if val.Tipo == Enum.ARRAY || val.Tipo == Enum.VECTOR {
				simb := val.Valor
				a := simb.(Simbolo.Simbolo).Valor
				fmt.Println("Tipo de a: ", reflect.TypeOf(a))
				fmt.Println("Imprimiendo a:", a)
				var recol string = ""
				if reflect.TypeOf(a) != reflect.TypeOf(arrtemp) {
					fmt.Println("Esta dentro del if ya que es un simbolo y no un arraylist por que es: ", reflect.TypeOf(a))
					tmpDato := fmt.Sprintf("%v", a)
					st = strings.Replace(st.(string), "{:?}", tmpDato, 1)
					recolector.Consolavirtual.Add(st.(string) + "\n")
					return st
				} else {
					//recorro todo el arreglo y lo almaceno en una temrporal recol donde se almacena el valor de cada posicion del arreglo
					for _, s := range a.(*arrayList.List).ToArray() {
						fmt.Println("s: ", s.(Simbolo.Simbolo).Valor)
						tmpDato := fmt.Sprintf("%v", s.(Simbolo.Simbolo).Valor)
						recol += tmpDato
					}
				}
				//recol += "]"
				st = strings.Replace(st.(string), "{:?}", recol, 1)
				recolector.Consolavirtual.Add(st.(string) + "\n")
				return nil
			} else {
				//si no lo es lo asigno tal cual fuera una variable
				fmt.Println("Val: ", val.Valor)
				tmpDato := fmt.Sprintf("%v", val.Valor)
				st = strings.Replace(st.(string), "{}", tmpDato, 1)
			}
		}
		recolector.Consolavirtual.Add(st.(string) + "\n")
		fmt.Println("************El resultado de la sustitucion es: ", st)
		return st
	}
}

func nousar() {
	list := arrayList.New()
	list.Add("hola mundo")
	fmt.Println(list)
}
