package Estructura

import (
	"fmt"
)

type Entorno struct {
	prev   interface{}
	nombre string
	numero int
	//variables = lista
	//funciones = lista
	//estructura = lista
}

func NuevoEntorno(_prev interface{}, _nombre string, _numero int) Entorno {
	fmt.Printf("Nuevo Entorno creado")
	return Entorno{prev: _prev, nombre: _nombre, numero: _numero}
}
