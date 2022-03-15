package Instruccion

import (
	"fmt"
	"proyecto1/Estructura"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
	"reflect"

	arrayList "github.com/colegno/arraylist"
)

type If struct {
	Expresion    Interfaces.Expresion
	Bloqueinst   *arrayList.List
	Bloqueifelse *arrayList.List
	Bloqueelse   *arrayList.List
}

func NewIf(condicion Interfaces.Expresion, bloque *arrayList.List, bloque2 *arrayList.List, bloque3 *arrayList.List) If {
	return If{condicion, bloque, bloque2, bloque3}
}

func (iff If) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {

	fmt.Println("Ingresa a ejecutar del if")
	var condicion = iff.Expresion.Ejecutar(env, recolector)
	fmt.Println("Condicion: ", condicion)

	if condicion.Tipo != Interfaces.BOOLEAN {
		fmt.Println("El resultado no es una booleanda: ", condicion.Tipo)
		return nil
	}
	/*
		fmt.Println("Condicion if: ", iff.Bloqueinst)
		fmt.Println("Condicion elseif: ", iff.Bloqueifelse)
		fmt.Println("Condicion else: ", iff.Bloqueelse)
	*/
	if condicion.Valor.(bool) {
		var NuevoEntorno Estructura.Entorno
		var num int = env.(Estructura.Entorno).GetNumEntorno()
		NuevoEntorno = Estructura.NuevoEntorno(env, "IF", num+1)
		fmt.Println("If si es true")
		for j := 0; j < iff.Bloqueinst.Len(); j++ {
			instr := iff.Bloqueinst.GetValue(j).(Interfaces.Instruccion)
			fmt.Println("Dentro del for-if: ", instr)
			b := NewBreak()
			if reflect.TypeOf(instr) == reflect.TypeOf(b) {
				fmt.Println("Entra en break----------------------")
				fmt.Println("Entorno: ", NuevoEntorno)
				fmt.Println("Entorno padre: ", env)
				condicion.Valor = false
				return NewBreak()
			} else {
				instr.Ejecutar(NuevoEntorno, recolector)
			}
		}
	} else {
		fmt.Println("No entra la if, entra al else o else if", iff)
		if iff.Bloqueifelse != nil {
			fmt.Println(" Si entra al IF ELSE ")
			for _, elseif_instruccion := range iff.Bloqueifelse.ToArray() {
				nuevoif2 := elseif_instruccion.(If)
				retornonuevo := nuevoif2.Expresion.Ejecutar(env, recolector)
				if retornonuevo.Tipo != Interfaces.BOOLEAN {
					return nil
				}
				if retornonuevo.Valor.(bool) {
					fmt.Println(" Se evalua el else if y es verdadero ")
					//var NuevoEntorno1 Estructura.Entorno
					var num Estructura.Entorno = env.(Estructura.Entorno)
					entnuevoelseif := Estructura.NuevoEntorno(env, "ElseIF", num.GetNumEntorno()+1)
					for j := 0; j < iff.Bloqueifelse.Len(); j++ {
						instr2 := iff.Bloqueifelse.GetValue(j).(Interfaces.Instruccion)
						instr2.Ejecutar(entnuevoelseif, recolector)
					}
					return nil
				}

			}
		}
		if iff.Bloqueelse != nil {
			fmt.Println(" Si entra al else ")
			//var NuevoEntorno2 Estructura.Entorno
			var num Estructura.Entorno = env.(Estructura.Entorno)
			entnuevoelseif := Estructura.NuevoEntorno(env, "Else final", num.GetNumEntorno()+1)
			for j := 0; j < iff.Bloqueelse.Len(); j++ {
				instr := iff.Bloqueelse.GetValue(j).(Interfaces.Instruccion)
				instr.Ejecutar(entnuevoelseif, recolector)
			}
		}
	}
	return nil
}
