package Expresion

import (
	"proyecto1/Enum"
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"
)

type Primitivo struct {
	Valor interface{}
	Tipo  Enum.Tipoexpresion
	//linea
	//columna
}

func (p Primitivo) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo {
	return Simbolo.Simbolo{
		Id:    "",
		Tipo:  p.Tipo,
		Valor: p.Valor,
	}
}

func NuevoPrimitivo(val interface{}, tipo Enum.Tipoexpresion) Primitivo {
	exp := Primitivo{val, tipo}
	return exp
}

func (p Primitivo) GetHolamundo() string {
	return "Hola desde Hola mundo"
}
