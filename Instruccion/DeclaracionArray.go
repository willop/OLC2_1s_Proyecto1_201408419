package Instruccion

import (
	"fmt"
	"proyecto1/Enum"
	"proyecto1/Estructura"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
	//"reflect"
	//arrayList "github.com/colegno/arraylist"
)

type DeclaracionArray struct {
	Id        string
	Tipo      Enum.Tipoexpresion
	Cant      Interfaces.Expresion
	Expresion Interfaces.Expresion
	Mutable   bool
}

// para declarar una nueva variable se recolecta toda la informacion para crear una nueva variable
func NuevaDeclaracionArray(_id string, _tipo Enum.Tipoexpresion, _cant Interfaces.Expresion, _expresion Interfaces.Expresion, _mut bool) DeclaracionArray {
	return DeclaracionArray{_id, _tipo, _cant, _expresion, _mut}
}

func (d DeclaracionArray) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {
	fmt.Println("En declaracion de variable array\nImprimiendo las variables:")

	valor := d.Expresion.Ejecutar(env, recolector)
	env.(Estructura.Entorno).GuardarSimbolo(d.Id, valor, d.Mutable, Enum.ARRAY)

	/*fmt.Println(d)
	arreglo := d.Expresion.Ejecutar(env, recolector)
	fmt.Println("Imprimiendo el arreglo:\n", arreglo.Valor)
	fmt.Println("Tipo del arreglo:\n", reflect.TypeOf(arreglo))
	lis := arrayList.New()
	if reflect.TypeOf(arreglo.Valor) == reflect.TypeOf(lis) {
		lis = arreglo.Valor.(*arrayList.List)
		for _, s := range lis.ToArray() {
			fmt.Println(s)
		}
	}
	fmt.Println("Final de for")*/

	return nil
}
