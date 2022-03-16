package Instruccion

import (
	"fmt"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
	"reflect"
	"strconv"
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

	if p.Bloqueinst == nil {
		fmt.Println("Luego de ejecutar ", resultado)
		fmt.Println("Consolav: ", resultado.Valor)
		st := 45
		stt := 455.55
		if reflect.TypeOf(resultado.Valor) == reflect.TypeOf(st) {
			st := strconv.Itoa(resultado.Valor.(int)) + "\n"
			recolector.Consolavirtual.Add(st)
		} else if reflect.TypeOf(resultado.Valor) == reflect.TypeOf(stt) {
			st := strconv.FormatFloat(resultado.Valor.(float64), 'f', 2, 64) + "\n"
			recolector.Consolavirtual.Add(st)
		} else {
			recolector.Consolavirtual.Add(resultado.Valor.(string) + "\n")
		}
		return resultado.Valor
	} else {
		st := resultado.Valor
		for j := 0; j < p.Bloqueinst.Len(); j++ {
			instr := p.Bloqueinst.GetValue(j).(Interfaces.Expresion)
			val := instr.Ejecutar(env, recolector)
			fmt.Println("Val: ", val.Valor)
			stt := 50
			stt2 := 50.80
			if reflect.TypeOf(val.Valor) == reflect.TypeOf(stt) {
				st = strings.Replace(st.(string), "{}", strconv.Itoa(val.Valor.(int)), 1)
			} else if reflect.TypeOf(val.Valor) == reflect.TypeOf(stt2) {
				st = strings.Replace(st.(string), "{}", strconv.FormatFloat(val.Valor.(float64), 'f', 2, 64), 1)
			} else {
				st = strings.Replace(st.(string), "{}", val.Valor.(string), 1)
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
