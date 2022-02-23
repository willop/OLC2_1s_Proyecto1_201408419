// Code generated from gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// gramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type gramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExpresiones is called when entering the expresiones production.
	EnterExpresiones(c *ExpresionesContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExpresiones is called when exiting the expresiones production.
	ExitExpresiones(c *ExpresionesContext)
}
