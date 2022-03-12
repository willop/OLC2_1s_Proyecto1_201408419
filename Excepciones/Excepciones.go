package Excepciones

import (
	"proyecto1/Interfaces"
)

type Error struct {
	Descripcion string
	Ambito      interface{}
	tipo        Interfaces.Tipoerror
}

func ErrorGeneral(_decripcion string, _ambito interface{}) Error {
	return Error{_decripcion, _ambito, Interfaces.ERRORGENERAL}
}

func ErrorOperacion(_decripcion string, _ambito interface{}) Error {
	return Error{_decripcion, _ambito, Interfaces.OPERACION_NO_PERMITIDA}
}

func ErrorTipoIncorrecto(_decripcion string, _ambito interface{}) Error {
	return Error{_decripcion, _ambito, Interfaces.TIPO_INCORRECTO}
}

func ErrorVariableYaDeclarada(_decripcion string, _ambito interface{}) Error {
	return Error{_decripcion, _ambito, Interfaces.VARIABLE_YA_DECLARADA}
}

func ErrorVariableNoDeclarada(_decripcion string, _ambito interface{}) Error {
	return Error{_decripcion, _ambito, Interfaces.VARIABLE_NO_DECLARADA}
}
