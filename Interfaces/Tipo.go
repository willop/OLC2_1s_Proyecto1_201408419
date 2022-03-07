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
	BRAKE                         //6
	NULL                          //7
	ARRAY                         //8
	STRUCT                        //9
	RETURN                        //10
	USIZE                         //11
	FLOAT                         //12
	SINTIPO                       //13
)
