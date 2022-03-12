package Interfaces

type Return struct {
	Valor interface{}
	Tipo  Tipoexpresion
	//auxtipo Tipoexpresion
}

func NewReturn(_valor interface{}, _tipo Tipoexpresion) Return {
	return Return{Valor: _valor, Tipo: _tipo}
}
