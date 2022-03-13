package Operaciones

import (
	//"fmt"

	"proyecto1/Interfaces"
	//"proyecto1/Utilitario"
)

type ExpComa struct {
	izquierda Interfaces.Expresion
	derecha   Interfaces.Expresion
}

func NewExpComa(izquierda Interfaces.Expresion, derecha Interfaces.Expresion) ExpComa {
	return ExpComa{izquierda, derecha}
	//return izquierda
}

/*
func (p ExpComa) Ejecutar(env interface{}, recolector *Utilitario.Recolector) interface{} {
	fmt.Println("Ejecutar expcoma")
	expder := p.derecha.Ejecutar(env, recolector)
	expizq := p.izquierda.Ejecutar(env, recolector)

	fmt.Println("Imprimiendo la expresion derecha: ", expder)
	fmt.Println("Imprimiendo la expresion izquierda: ", expizq)

	//return Interfaces.Return{"Retornoexpcoma", Interfaces.STRING}
	//return p.izquierda
}
*/
