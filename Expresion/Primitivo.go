package Expresion

import (
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
)

type Primitivo struct {
	Valor interface{}
	Tipo  Interfaces.Tipoexpresion
	//linea
	//columna
}

func (p Primitivo) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Interfaces.Simbolo {

	return Interfaces.Simbolo{
		Id:    "",
		Tipo:  p.Tipo,
		Valor: p.Valor,
	}
}

func NuevoPrimitivo(val interface{}, tipo Interfaces.Tipoexpresion) Primitivo {
	exp := Primitivo{val, tipo}
	return exp
}

func (p Primitivo) GetHolamundo() string {
	return "Hola desde Hola mundo"
}
