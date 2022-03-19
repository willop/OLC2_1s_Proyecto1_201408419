package Instruccion

import (
	"fmt"

	"proyecto1/Enum"
	"proyecto1/Estructura"
	"proyecto1/Excepciones"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"

	"reflect"

	arrayList "github.com/colegno/arraylist"
)

type Whilee struct {
	Expresion  Interfaces.Expresion
	Bloqueinst *arrayList.List
}

func NewWhile(expresion Interfaces.Expresion, bloqueinst *arrayList.List) Whilee {
	return Whilee{expresion, bloqueinst}
}

func (w Whilee) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {

	fmt.Println("Ingresa a ejecutar el while")
	var condicion = w.Expresion.Ejecutar(env, recolector)
	fmt.Println("Condicion:", condicion)

	if condicion.Tipo != Enum.BOOLEAN {
		recolector.ListaErrores.Add(Excepciones.ErrorGeneral("La expresion del while no es booleana", env))
		return nil
	}

	fmt.Println("Instrucciones while", w.Bloqueinst)
	for condicion.Valor.(bool) {
		//creo el entorno para el while
		var NuevoEntorno Estructura.Entorno
		var num Estructura.Entorno = env.(Estructura.Entorno)
		NuevoEntorno = Estructura.NuevoEntorno(NuevoEntorno, "While", num.GetNumEntorno()+1)
		var banddd bool = true
		if banddd {
			for j := 0; j < w.Bloqueinst.Len(); j++ {
				instr := w.Bloqueinst.GetValue(j).(Interfaces.Instruccion)
				fmt.Println("Dentro del while: ", instr)
				b := NewBreak()
				if reflect.TypeOf(instr) == reflect.TypeOf(b) {
					banddd = false
					return nil
				} else {
					retorno := instr.Ejecutar(env, recolector)
					fmt.Println("Despues de ejecutar")
					if reflect.TypeOf(retorno) == reflect.TypeOf(b) {
						banddd = false
						return nil
					}
				}
			}
			condicion = w.Expresion.Ejecutar(env, recolector)
			fmt.Println("La nueva instruccion es: ", condicion)
			if !condicion.Valor.(bool) {
				return nil
			}
		}

	}
	return nil
}
