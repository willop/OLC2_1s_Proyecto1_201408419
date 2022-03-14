package Instruccion

import (
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
)

type Break struct {
	Tiporetorno Interfaces.Tipoexpresion
}

func NewBreak() Break {
	return Break{Tiporetorno: Interfaces.BREAK}
}

func (b Break) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {
	return b.Tiporetorno
}
