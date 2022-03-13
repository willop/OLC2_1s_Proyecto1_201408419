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
	env := Entorno{prev: _prev, nombre: _nombre, numero: _numero, variable: make(map[string]Interfaces.Simbolo), estructura: make(map[string]Interfaces.Simbolo)}
	return env
}

//Metodo para almacenar una variable
func (env Entorno) GuardarSimbolo(_id string, _valor interface{}, _mut bool, _tipo Interfaces.Tipoexpresion) {
	//verificar que existe la variable en este entorno
	if variable, encontrada := env.variable[_id]; encontrada {
		fmt.Println("La variable ", variable.Id, " ya existe en este entorno")
		return
	}
	//si no pues agrego una nueva variable al entorno
	fmt.Println("Se agrego al entorno")
	env.variable[_id] = Interfaces.Simbolo{Id: _id, Valor: _valor, Mut: _mut, Tipo: _tipo}
	fmt.Println("Sale de guardar simbolo", env.variable[_id])
}

func (env Entorno) GetNumEntorno() int {
	return env.numero
}

func (env Entorno) GetVariable(id string) Interfaces.Simbolo {
	var temporal Entorno
	temporal = env
	for {
		if variable, ok := temporal.variable[id]; ok {
			return variable
		}
		if temporal.prev == nil {
			break
		} else {
			temporal = temporal.prev.(Entorno)
		}
	}
	fmt.Println("La variable no existe")
	return Interfaces.Simbolo{"", "error", false, Interfaces.SINTIPO}
}

func (env Entorno) ActualizarSimbolo(_id string, _valor interface{}, _mut bool, _tipo Interfaces.Tipoexpresion) interface{} {

	var tmpEnv Entorno
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.variable[_id]; ok {
			tmpEnv.variable[_id] = Interfaces.Simbolo{Id: _id, Tipo: _tipo, Valor: _valor, Mut: _mut}
			return variable
		}

		if tmpEnv.prev == nil {
			break
		} else {
			tmpEnv = tmpEnv.prev.(Entorno)
		}
	}

	fmt.Println("La variable no existe")
	return Interfaces.Simbolo{Id: "", Tipo: Interfaces.SINTIPO, Valor: "", Mut: false}

}
