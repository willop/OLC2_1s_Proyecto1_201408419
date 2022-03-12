package Instruccion

import (
	"fmt"
	"proyecto1/Estructura"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"

	arrayList "github.com/colegno/arraylist"
)

type Instrucciones struct {
	Instrucciones *arrayList.List
	Nombre        string
	Crearentorno  bool
}

func NewInstrucciones(_istrucciones *arrayList.List, _nombre string, _crearentorno bool) Instrucciones {
	return Instrucciones{Instrucciones: _istrucciones, Nombre: _nombre, Crearentorno: true}
}

func (i *Instrucciones) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {

	var NuevoEntorno Estructura.Entorno
	if i.Crearentorno {
		var num Estructura.Entorno = env.(Estructura.Entorno)
		NuevoEntorno = Estructura.NuevoEntorno(env, i.Nombre, num.GetNumEntorno()+1)
	} else {
		NuevoEntorno = env.(Estructura.Entorno)
	}
	for _, v := range i.Instrucciones.ToArray() {
		fmt.Println("----------------------FOR INSTRUCCIONES------------------")
		fmt.Println(v.(Interfaces.Instruccion))
		fmt.Println(v.(Interfaces.Instruccion).Ejecutar(NuevoEntorno, recolector))
		v.(Interfaces.Instruccion).Ejecutar(NuevoEntorno, recolector)
		fmt.Println("---------------------------------------------------------")
	}

	return nil
}
