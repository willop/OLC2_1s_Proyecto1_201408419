package Estructura

import (
	"fmt"
	"proyecto1/Enum"
	Interfaces "proyecto1/Enum"

	//arrayList "github.com/colegno/arraylist"

	//"proyecto1/Funcion"
	"proyecto1/Simbolo"
)

type Entorno struct {
	prev       interface{}
	nombre     string
	numero     int
	variable   map[string]Simbolo.Simbolo
	estructura map[string]Simbolo.Simbolo
	funciones  map[string]interface{}
}

//Se crea y se retorna un nuevo entorno
func NuevoEntorno(_prev interface{}, _nombre string, _numero int) Entorno {
	fmt.Printf("Nuevo Entorno creado")
	env := Entorno{prev: _prev, nombre: _nombre, numero: _numero, variable: make(map[string]Simbolo.Simbolo), estructura: make(map[string]Simbolo.Simbolo), funciones: make(map[string]interface{})}
	return env
}

//Metodo para almacenar una variable
func (env Entorno) GuardarSimbolo(_id string, _valor interface{}, _mut bool, _tipo Enum.Tipoexpresion) {
	//verificar que existe la variable en este entorno
	if variable, encontrada := env.variable[_id]; encontrada {
		fmt.Println("La variable ", variable.Id, " ya existe en este entorno")
		return
	}
	//si no pues agrego una nueva variable al entorno
	fmt.Println("Se agrego al entorno")
	env.variable[_id] = Simbolo.Simbolo{Id: _id, Valor: _valor, Mut: _mut, Tipo: _tipo}
	fmt.Println("Sale de guardar simbolo", env.variable[_id])
}

func (env Entorno) GetNumEntorno() int {
	return env.numero
}

func (env Entorno) GetVariable(id string) Simbolo.Simbolo {
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
	return Simbolo.Simbolo{"", "error", false, Interfaces.SINTIPO}
}

func (env Entorno) ActualizarSimbolo(_id string, _valor interface{}, _mut bool, _tipo Enum.Tipoexpresion) interface{} {

	var tmpEnv Entorno
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.variable[_id]; ok {
			varr := tmpEnv.variable[_id]
			tmpEnv.variable[_id] = Simbolo.Simbolo{Id: _id, Tipo: _tipo, Valor: _valor, Mut: varr.Mut}
			return variable
		}

		if tmpEnv.prev == nil {
			break
		} else {
			tmpEnv = tmpEnv.prev.(Entorno)
		}
	}

	fmt.Println("La variable no existe")
	return nil // Interfaces.Simbolo{Id: "", Tipo: Interfaces.SINTIPO, Valor: "", Mut: false}

}

func (env Entorno) GuardarFuncion(_fun interface{}, _id string) interface{} {
	//cada funcion va a tener su propio entorno
	if _, encontrada := env.funciones[_id]; encontrada {
		fmt.Println("La funcion ", _id, " ya existe en este entorno")
		return nil
	}
	env.funciones[_id] = _fun
	fmt.Println("******************************** En el entorno: ", env, " se agrego la funcion: ", _fun)
	return nil
}

func (env Entorno) ObtenerFuncion(_id string) interface{} {
	var temporal Entorno
	temporal = env
	for {
		if variable, ok := temporal.funciones[_id]; ok {
			return variable
		}
		if temporal.prev == nil {
			break
		} else {
			temporal = temporal.prev.(Entorno)
		}
	}
	fmt.Println("La variable no existe")
	return nil //   Funcion.Funcion{"", Enum.ERROREXPRESION, nil, nil}

}
