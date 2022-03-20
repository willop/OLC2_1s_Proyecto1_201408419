package Funcion

import (
	"fmt"
	"proyecto1/Estructura"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"

	arrayList "github.com/colegno/arraylist"
)

type Llamadafuncion struct {
	Id         string
	Llamadaexp bool
	Parametros *arrayList.List
}

func NewLlamadafuncion(_id string, _llamadaexp bool, _param *arrayList.List) Llamadafuncion {
	return Llamadafuncion{_id, _llamadaexp, _param}
}

func (ll Llamadafuncion) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {
	fmt.Println("Esta en ejecutar NewLlamadafuncion")

	funcion := env.(Estructura.Entorno).ObtenerFuncion(ll.Id)

	//verifico si obtuve bien la funcion del entorno
	if funcion != nil {
		fmt.Println("Funcion: ", funcion)

		//verifico si existen parametros
		if ll.Parametros != nil {
			fmt.Println("Si tiene parametros")
			//linea de 45 al 75
		} else {
			fmt.Println("*****Se crea un nuevo entorno")
			fmt.Println(env)

			aux := funcion.(Funcion).Instrucciones
			fmt.Println("Las instrucciones son:", aux)
			for _, f := range funcion.(Funcion).Instrucciones.ToArray() {
				f.(Interfaces.Instruccion).Ejecutar(env, recolector)
			}
		}
	}
	return nil
}
