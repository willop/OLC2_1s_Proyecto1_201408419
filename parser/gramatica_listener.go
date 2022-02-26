// Code generated from gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// gramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type gramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterTipovariable is called when entering the tipovariable production.
	EnterTipovariable(c *TipovariableContext)

	// EnterIgualacion is called when entering the igualacion production.
	EnterIgualacion(c *IgualacionContext)

	// EnterIdentificadores is called when entering the identificadores production.
	EnterIdentificadores(c *IdentificadoresContext)

	// EnterValores is called when entering the valores production.
	EnterValores(c *ValoresContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterImpresion is called when entering the impresion production.
	EnterImpresion(c *ImpresionContext)

	// EnterImpresioncomas is called when entering the impresioncomas production.
	EnterImpresioncomas(c *ImpresioncomasContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitTipovariable is called when exiting the tipovariable production.
	ExitTipovariable(c *TipovariableContext)

	// ExitIgualacion is called when exiting the igualacion production.
	ExitIgualacion(c *IgualacionContext)

	// ExitIdentificadores is called when exiting the identificadores production.
	ExitIdentificadores(c *IdentificadoresContext)

	// ExitValores is called when exiting the valores production.
	ExitValores(c *ValoresContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitImpresion is called when exiting the impresion production.
	ExitImpresion(c *ImpresionContext)

	// ExitImpresioncomas is called when exiting the impresioncomas production.
	ExitImpresioncomas(c *ImpresioncomasContext)
}
