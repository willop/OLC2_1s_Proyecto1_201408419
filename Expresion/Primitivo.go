package Expresion

import "proyecto1/Interfaces"

type Primitivo struct {
	Valor interface{}
	Tipo  Interfaces.Tipoexpresion
	//linea
	//columna
}

func (p Primitivo) Ejecutar(env interface{}) Interfaces.Simbolo {

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

func GetHolamundo() string {
	return "Hola desde Hola mundo"
}
