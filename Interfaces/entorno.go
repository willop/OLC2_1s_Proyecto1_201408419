package Interfaces

import (
	"fmt"
)

type Entorno struct {
	prev        interface{}
	nombre      interface{}
	numero      interface{}
	Variables   map[interface{}]*Simbolo
	funciones   interface{}
	estructuras interface{}
}

//metodo constructor para un entorno
func ConstructorEntorno(_prev interface{}, _nombre interface{}, _numero interface{}, _variables interface{}, _funciones interface{}, estructuras interface{}) *Entorno {
	return &Entorno{}
}

func (_Entorno Entorno) GuardarSimbolo(_valor interface{}, _identificador interface{}, _tipo int, _mutable interface{}) {
	sim := ConstructorSimbolo(_valor, _identificador, _tipo, 5) //("","","",5)
	fmt.Println("El simbolo es: ", sim.GetValor())
	//luego de la validacion
	_Entorno.Variables[sim.Id] = sim
	//sim := ConstructorSimbolo{_identificador, _valor, _mutable, _tipo}
}
