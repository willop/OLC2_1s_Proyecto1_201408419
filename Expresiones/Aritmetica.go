package Expresiones

import (
	"fmt"

	"proyecto1/Interfaces"
	"proyecto1/Operaciones"
	"proyecto1/Utilitario"
)

type Aritmetica struct {
	Izquierda Interfaces.Expresion
	Derecha   Interfaces.Expresion
	Tipo      Interfaces.Tipooperacion
	//linea
	//columna
}

func NuevaAritmetica(_izquierda Interfaces.Expresion, _derecha Interfaces.Expresion, _tipo Interfaces.Tipooperacion) Aritmetica {
	Arit := Aritmetica{Izquierda: _izquierda, Derecha: _derecha, Tipo: _tipo}
	return Arit
}

func (_arit Aritmetica) Ejecutar(env interface{}, recolector *Utilitario.Recolector) Interfaces.Simbolo {
	fmt.Println("Incia expresion aritmetica")
	var tipotipo Interfaces.Tipooperacion = _arit.Tipo
	var izquierdaa Interfaces.Simbolo = _arit.Izquierda.Ejecutar(env, recolector)
	var derechaa Interfaces.Simbolo = _arit.Derecha.Ejecutar(env, recolector)
	if derechaa.Valor == -1 {
		fmt.Println("*************---------- Entra con el unario--------------**************")
		derechaa.Tipo = izquierdaa.Tipo
		derechaa.Valor = float64(derechaa.Valor.(int))
		fmt.Println(derechaa)
	}

	switch tipotipo {
	case Interfaces.SUMA:
		fmt.Printf("Estoy en la suma")
		result := Operaciones.Sumar(izquierdaa, derechaa)
		fmt.Println(result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	case Interfaces.RESTA:
		fmt.Printf("Estoy en la resta")
		result := Operaciones.Restar(izquierdaa, derechaa)
		fmt.Println(result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

	case Interfaces.DIVISION:
		fmt.Printf("Estoy en la division")
		result := Operaciones.Dividir(izquierdaa, derechaa)
		fmt.Println(result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

	case Interfaces.MULTIPLICACION:
		fmt.Printf("Estoy en la multiplicacion")
		result := Operaciones.Multiplicar(izquierdaa, derechaa)
		fmt.Println(result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

	case Interfaces.MODULO:
		fmt.Printf("Estoy en modulo")
		result := Operaciones.Modulo(izquierdaa, derechaa)
		fmt.Println(result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

	case Interfaces.POW:
		fmt.Printf("Estoy en pow")
		result := Operaciones.Pow(izquierdaa, derechaa)
		fmt.Println(result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result

	case Interfaces.POWF:
		fmt.Printf("Estoy en pow flotante")
		result := Operaciones.Powf(izquierdaa, derechaa)
		fmt.Println(result)
		if result.Tipo == Interfaces.SINTIPO {
			return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: result.Valor, Mut: false}
		}
		return result
	}

	fmt.Println("Nueva aritmetica")
	fmt.Println("Valor del tipo es: ", _arit.Tipo)
	var resultado interface{}
	return Interfaces.Simbolo{Id: "Error_op", Tipo: Interfaces.SINTIPO, Valor: resultado, Mut: false}
}
