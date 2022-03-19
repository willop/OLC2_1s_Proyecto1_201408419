package Excepciones

import (
	"proyecto1/Enum"
)

type Error struct {
	Descripcion string
	Ambito      interface{}
	tipo        Enum.Tipoerror
}

func ErrorGeneral(_decripcion string, _ambito interface{}) Error {
	return Error{_decripcion, _ambito, Enum.ERRORGENERAL}
}

func ErrorOperacion(_decripcion string, _ambito interface{}) Error {
	return Error{_decripcion, _ambito, Enum.OPERACION_NO_PERMITIDA}
}

func ErrorTipoIncorrecto(_decripcion string, _ambito interface{}) Error {
	return Error{_decripcion, _ambito, Enum.TIPO_INCORRECTO}
}

func ErrorVariableYaDeclarada(_decripcion string, _ambito interface{}) Error {
	return Error{_decripcion, _ambito, Enum.VARIABLE_YA_DECLARADA}
}

func ErrorVariableNoDeclarada(_decripcion string, _ambito interface{}) Error {
	return Error{_decripcion, _ambito, Enum.VARIABLE_NO_DECLARADA}
}
