package Instruccion

import (
	"fmt"

	"proyecto1/Estructura"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"

	"reflect"

	arrayList "github.com/colegno/arraylist"
)

type Loop struct {
	Bloqueinst *arrayList.List
}

func NewLoop(bloqueinst *arrayList.List) Loop {
	return Loop{bloqueinst}
}

func (l Loop) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {

	fmt.Println("Ingresa a ejecutar el loop")
	var bandera bool = true
	for bandera {
		fmt.Printf("La condicion del loop es verdadera")
		//creo el entorno para el while
		var NuevoEntorno Estructura.Entorno
		var num Estructura.Entorno = env.(Estructura.Entorno)
		NuevoEntorno = Estructura.NuevoEntorno(NuevoEntorno, "Loop", num.GetNumEntorno()+1)
		for j := 0; j < l.Bloqueinst.Len(); j++ {
			instr := l.Bloqueinst.GetValue(j).(Interfaces.Instruccion)
			fmt.Println("Dentro del Loop: ", instr)

			b := NewBreak()
			if reflect.TypeOf(instr) == reflect.TypeOf(b) {
				fmt.Println("Entra en break----------------------")
				bandera = false
				return nil
			} else {
				instr.Ejecutar(env, recolector)
				fmt.Println("Despues de ejecutar")
			}

		}

	}
	return nil
}
