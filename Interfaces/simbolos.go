package Interfaces

//es un enum para saber el tipo de la expresion
type tipoexpresion int

//
const (
	STRING tipoexpresion = iota //variable que funciona como un switch
	INTEGER
	CHAR
	DOUBLE
	BOOLEAN
	CONTINUE
	BRAKE
	NULL
	ARRAY
	STRUCT
	RETURN
)

//creo un objeto para almacenar la informacion de una variable
type Simbolo struct {
	Id    interface{}
	Valor interface{}
	Mut   interface{}
	Tipo  tipoexpresion
}

func ContructorSimbolo(_id interface{}, _valor interface{}, _mut interface{}, _tipo tipoexpresion) *Simbolo {
	return &Simbolo{Id: _id, Valor: _valor, Mut: _mut, Tipo: _tipo}
}

//get y set id
func (_Simbolo *Simbolo) GetID() interface{} {
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
func (_Simbolo *Simbolo) GetTipo() tipoexpresion {
	return _Simbolo.Tipo
}
func (_Simbolo *Simbolo) SetTipo(_tipo tipoexpresion) {
	_Simbolo.Id = _tipo
}
