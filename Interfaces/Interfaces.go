package Interfaces

import (
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"
)

//creo un objeto para almacenar la informacion de una variable

type Expresion interface {
	Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo
}

type Instruccion interface {
	Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{}
}

/*
//get y set id
func (_Simbolo *Simbolo.Simbolo) GetID() interface{} {
	return _Simbolo.Id
}
func (_Simbolo *Simbolo) SetID(_id interface{}) {
	_Simbolo.Id = _id
}

//get y set valor
func (_Simbolo *Simbolo) GetValor() interface{} {
	return _Simbolo.Valor
}
func (_Simbolo *Simbolo) SetValor(_valor interface{}) {
	_Simbolo.Id = _valor
}

//get y set valor
func (_Simbolo *Simbolo) GetMut() interface{} {
	return _Simbolo.Mut
}
func (_Simbolo *Simbolo) SetMut(_mut interface{}) {
	_Simbolo.Id = _mut
}

//get y set tipo
func (_Simbolo *Simbolo) GetTipo() Tipoexpresion {
	return _Simbolo.Tipo
}
func (_Simbolo *Simbolo) SetTipo(_tipo Tipoexpresion) {
	_Simbolo.Id = _tipo
}
*/
