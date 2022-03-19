package Simbolo

import (
	"fmt"
	"proyecto1/Enum"
	"reflect"
)

type Simbolo struct {
	Id    interface{}
	Valor interface{}
	Mut   interface{}
	Tipo  Enum.Tipoexpresion
}

func ConstructorSimbolo(_id interface{}, _valor interface{}, _mut interface{}, _tipo Enum.Tipoexpresion) *Simbolo {
	if _tipo == Enum.SINTIPO {
		tipovar := reflect.TypeOf(_valor)
		fmt.Println("valor del sin tipo: ", tipovar)
	}

	fmt.Println("Dentro del constructro de un simbolo")
	fmt.Println("Valores recividos: ", _valor, _mut, _tipo, _id)
	return &Simbolo{Id: _id, Valor: _valor, Mut: _mut, Tipo: _tipo}

}
