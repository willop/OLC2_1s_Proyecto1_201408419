package Expresiones

import (
	"fmt"
	"proyecto1/Enum"
	"proyecto1/Interfaces"

	"proyecto1/Simbolo"
	"proyecto1/Utilitario"

	arrayList "github.com/colegno/arraylist"
)

type AccesoArray struct {
	Variable Interfaces.Expresion
	Indice   Interfaces.Expresion
}

func NewAccesoArray(_id Interfaces.Expresion, _indice Interfaces.Expresion) AccesoArray {
	return AccesoArray{_id, _indice}
}

func (a AccesoArray) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo {
	fmt.Println("Ejecutar accesoarray")
	temp := a.Variable.Ejecutar(env, recolector)
	ind := a.Indice.Ejecutar(env, recolector)
	fmt.Println("Imprimiendo termporal asignacion array: ", temp)
	fmt.Println("Esta es llamado accesoarray: ", temp.Valor.(Simbolo.Simbolo).Valor)
	fmt.Println("El valor del indice", ind.Valor, " y el tamanio es: ", temp.Valor.(Simbolo.Simbolo).Valor.(*arrayList.List).Len())
	if ind.Valor.(int) < temp.Valor.(Simbolo.Simbolo).Valor.(*arrayList.List).Len() {
		varr := temp.Valor.(Simbolo.Simbolo).Valor.(*arrayList.List).GetValue(ind.Valor.(int))
		fmt.Println("Valor de temp desde acceso array", varr, "El id de temporal es:", temp.Id)
		return Simbolo.Simbolo{Id: temp.Id, Valor: varr, Tipo: Enum.ARRAY}
	} else {
		return Simbolo.Simbolo{Id: temp.Id, Valor: "Error", Tipo: Enum.ERROREXPRESION}
	}
}
