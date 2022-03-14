package Interfaces

//es un enum para saber el tipo de la expresion
type Tipoexpresion int

//
const (
	STRING   Tipoexpresion = iota //variable que funciona como un switch
	INTEGER                       //1
	CHAR                          //2
	DOUBLE                        //3
	BOOLEAN                       //4
	CONTINUE                      //5
	BREAK                         //6
	NULL                          //7
	ARRAY                         //8
	STRUCT                        //9
	RETURN                        //10
	USIZE                         //11
	FLOAT                         //12
	SINTIPO                       //13
	STR
	ERROREXPRESION
	ERRORTIPOEXPRESION
)

type Tipooperacion int

const (
	SUMA Tipooperacion = iota //variable que funciona como un switch
	RESTA
	DIVISION
	MULTIPLICACION
	MODULO
	POW
	POWF
)

type Tipoerror int

const (
	ERRORGENERAL Tipoerror = iota //variable que funciona como un switch
	OPERACION_NO_PERMITIDA
	TIPO_INCORRECTO
	VARIABLE_NO_DECLARADA
	VARIABLE_YA_DECLARADA
	ERROR_LOGICO
	ERROR_RELACIONAL
	ERROR_BOOLEANO
	ERROR_TIPO_VARIABLE_ARRAY
	FUNCION_YA_DECLARADA
)

type Tipologico int

const (
	AND Tipologico = iota
	OR
	NOT
)

type Tiporelacional int

const (
	IGUALDAD Tiporelacional = iota //variable que funciona como un switch
	DESIGUAL
	MAYOR_QUE
	MENOR_QUE
	MAYOR_IGUAL
	MENOR_IGUAL
)
