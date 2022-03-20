// Code generated from gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// gramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type gramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterFunciones is called when entering the funciones production.
	EnterFunciones(c *FuncionesContext)

	// EnterDeclararParametros is called when entering the declararParametros production.
	EnterDeclararParametros(c *DeclararParametrosContext)

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterControl is called when entering the control production.
	EnterControl(c *ControlContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterTipovariable is called when entering the tipovariable production.
	EnterTipovariable(c *TipovariableContext)

	// EnterIdentificadores is called when entering the identificadores production.
	EnterIdentificadores(c *IdentificadoresContext)

	// EnterValores is called when entering the valores production.
	EnterValores(c *ValoresContext)

	// EnterArreglo is called when entering the arreglo production.
	EnterArreglo(c *ArregloContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterImpresion is called when entering the impresion production.
	EnterImpresion(c *ImpresionContext)

	// EnterImpresionexpresion is called when entering the impresionexpresion production.
	EnterImpresionexpresion(c *ImpresionexpresionContext)

	// EnterExpcoma is called when entering the expcoma production.
	EnterExpcoma(c *ExpcomaContext)

	// EnterCondicionales is called when entering the condicionales production.
	EnterCondicionales(c *CondicionalesContext)

	// EnterFuncionif is called when entering the funcionif production.
	EnterFuncionif(c *FuncionifContext)

	// EnterFuncionelseif is called when entering the funcionelseif production.
	EnterFuncionelseif(c *FuncionelseifContext)

	// EnterListaelseif is called when entering the listaelseif production.
	EnterListaelseif(c *ListaelseifContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterBucles is called when entering the bucles production.
	EnterBucles(c *BuclesContext)

	// EnterFwhile is called when entering the fwhile production.
	EnterFwhile(c *FwhileContext)

	// EnterFloop is called when entering the floop production.
	EnterFloop(c *FloopContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitFunciones is called when exiting the funciones production.
	ExitFunciones(c *FuncionesContext)

	// ExitDeclararParametros is called when exiting the declararParametros production.
	ExitDeclararParametros(c *DeclararParametrosContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitControl is called when exiting the control production.
	ExitControl(c *ControlContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitTipovariable is called when exiting the tipovariable production.
	ExitTipovariable(c *TipovariableContext)

	// ExitIdentificadores is called when exiting the identificadores production.
	ExitIdentificadores(c *IdentificadoresContext)

	// ExitValores is called when exiting the valores production.
	ExitValores(c *ValoresContext)

	// ExitArreglo is called when exiting the arreglo production.
	ExitArreglo(c *ArregloContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitImpresion is called when exiting the impresion production.
	ExitImpresion(c *ImpresionContext)

	// ExitImpresionexpresion is called when exiting the impresionexpresion production.
	ExitImpresionexpresion(c *ImpresionexpresionContext)

	// ExitExpcoma is called when exiting the expcoma production.
	ExitExpcoma(c *ExpcomaContext)

	// ExitCondicionales is called when exiting the condicionales production.
	ExitCondicionales(c *CondicionalesContext)

	// ExitFuncionif is called when exiting the funcionif production.
	ExitFuncionif(c *FuncionifContext)

	// ExitFuncionelseif is called when exiting the funcionelseif production.
	ExitFuncionelseif(c *FuncionelseifContext)

	// ExitListaelseif is called when exiting the listaelseif production.
	ExitListaelseif(c *ListaelseifContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitBucles is called when exiting the bucles production.
	ExitBucles(c *BuclesContext)

	// ExitFwhile is called when exiting the fwhile production.
	ExitFwhile(c *FwhileContext)

	// ExitFloop is called when exiting the floop production.
	ExitFloop(c *FloopContext)
}
