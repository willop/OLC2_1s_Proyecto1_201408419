package Estructura

import (
	"fmt"
	"proyecto1/Interfaces"
)

type Entorno struct {
	prev       interface{}
	nombre     string
	numero     int
	variable   map[string]Interfaces.Simbolo
	estructura map[string]Interfaces.Simbolo
	//funciones map[string]Interfaces.Simbolo
}

//Se crea y se retorna un nuevo entorno
func NuevoEntorno(_prev interface{}, _nombre string, _numero int) Entorno {
	fmt.Printf("Nuevo Entorno creado")
	return Entorno{_prev, _nombre, _numero, make(map[string]Interfaces.Simbolo), make(map[string]Interfaces.Simbolo)}
}

//Metodo para almacenar una variable
func (env Entorno) GuardarSimbolo(_id string, _valor Interfaces.Simbolo, _mut bool, _tipo Interfaces.Tipoexpresion) {
	//verificar que existe la variable en este entorno
	if variable, encontrada := env.variable[_id]; encontrada {
		fmt.Println("La variable ", variable.Id, " ya existe en este entorno")
		return
	}
	//si no pues agrego una nueva variable al entorno
	env.variable[_id] = Interfaces.Simbolo{Id: _id, Valor: _valor, Mut: _mut, Tipo: _tipo}
}
