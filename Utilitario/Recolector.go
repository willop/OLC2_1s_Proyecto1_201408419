package Utilitario

import (
	arrayList "github.com/colegno/arraylist"
)

type Recolector struct {
	Consolavirtual *arrayList.List
	ListaErrores   *arrayList.List
}

func NuevoRecolector() Recolector {
	return Recolector{Consolavirtual: arrayList.New(), ListaErrores: arrayList.New()}
}
