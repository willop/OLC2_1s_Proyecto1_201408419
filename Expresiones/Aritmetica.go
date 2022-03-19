package Expresiones

import (
	"fmt"

	"proyecto1/Enum"
	"proyecto1/Interfaces"
	"proyecto1/Operaciones"
	"proyecto1/Simbolo"
	"proyecto1/Utilitario"
)

type Aritmetica struct {
	Izquierda Interfaces.Expresion
	Derecha   Interfaces.Expresion
	Tipo      Enum.Tipooperacion
	//linea
	//columna
}

func NuevaAritmetica(_izquierda Interfaces.Expresion, _derecha Interfaces.Expresion, _tipo Enum.Tipooperacion) Aritmetica {
	Arit := Aritmetica{Izquierda: _izquierda, Derecha: _derecha, Tipo: _tipo}
	return Arit
}

func (_arit Aritmetica) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Simbolo.Simbolo {
	fmt.Println("Incia expresion aritmetica")
	var tipotipo Enum.Tipooperacion = _arit.Tipo
	var izquierdaa Simbolo.Simbolo = _arit.Izquierda.Ejecutar(env, recolector)
	var derechaa Simbolo.Simbolo = _arit.Derecha.Ejecutar(env, recolector)

	fmt.Println("izquierda", izquierdaa)
	fmt.Println("derecha", derechaa)
	if derechaa.Valor == -1 {
		fmt.Println("*************---------- Entra con el unario--------------**************")
		derechaa.Tipo = izquierdaa.Tipo
		if izquierdaa.Tipo == Enum.INTEGER {
			derechaa.Valor = derechaa.Valor.(int)
		} else {
			derechaa.Valor = -1.0
		}
		fmt.Println(derechaa)
	}

	switch tipotipo {
	case Enum.SUMA:
		fmt.Printf("Estoy en la suma")
		result := Operaciones.Sumar(izquierdaa, derechaa)
		fmt.Println(result)
		return result
	case Enum.RESTA:
		fmt.Printf("Estoy en la resta")
		result := Operaciones.Restar(izquierdaa, derechaa)
		fmt.Println(result)
		if result.Tipo == Enum.ERROREXPRESION {
			return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

	case Enum.DIVISION:
		fmt.Printf("Estoy en la division")
		result := Operaciones.Dividir(izquierdaa, derechaa)
		fmt.Println(result)
		return result

	case Enum.MULTIPLICACION:
		fmt.Printf("Estoy en la multiplicacion")
		result := Operaciones.Multiplicar(izquierdaa, derechaa)
		fmt.Println(result)
		return result

	case Enum.MODULO:
		fmt.Printf("Estoy en modulo")
		result := Operaciones.Modulo(izquierdaa, derechaa)
		fmt.Println(result)
		return result

	case Enum.POW:
		fmt.Printf("Estoy en pow")
		result := Operaciones.Pow(izquierdaa, derechaa)
		fmt.Println(result)
		return result

	case Enum.POWF:
		fmt.Printf("Estoy en pow flotante")
		result := Operaciones.Powf(izquierdaa, derechaa)
		fmt.Println(result)
		return result
	}

	fmt.Println("Nueva aritmetica")
	fmt.Println("Valor del tipo es: ", _arit.Tipo)
	var resultado interface{}
	return Simbolo.Simbolo{Id: "Error_op", Tipo: Enum.ERROREXPRESION, Valor: resultado, Mut: false}
}
