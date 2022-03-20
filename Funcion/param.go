package Funcion

import (
	"proyecto1/Enum"
)

type Param struct {
	ID   string
	TIpo Enum.Tipoexpresion
}

func NewParam(id string, Tipo Enum.Tipoexpresion) Param {
	return Param{id, Tipo}
}
