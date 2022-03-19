package Instruccion

import (
	"proyecto1/Enum"
	"proyecto1/Utilitario"
)

type Break struct {
	Tiporetorno Enum.Tipoexpresion
}

func NewBreak() Break {
	return Break{Tiporetorno: Enum.BREAK}
}

func (b Break) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {
	return b.Tiporetorno
}
