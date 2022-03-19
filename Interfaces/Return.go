package Interfaces

import (
	"proyecto1/Enum"
)

type Return struct {
	Valor interface{}
	Tipo  Enum.Tipoexpresion
	//auxtipo Tipoexpresion
}

func NewReturn(_valor interface{}, _tipo Enum.Tipoexpresion) Return {
	return Return{Valor: _valor, Tipo: _tipo}
}
