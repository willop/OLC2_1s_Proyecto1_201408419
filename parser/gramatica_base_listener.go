// Code generated from gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasegramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type BasegramaticaListener struct{}

var _ gramaticaListener = &BasegramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BasegramaticaListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BasegramaticaListener) ExitStart(ctx *StartContext) {}

// EnterFunciones is called when production funciones is entered.
func (s *BasegramaticaListener) EnterFunciones(ctx *FuncionesContext) {}

// ExitFunciones is called when production funciones is exited.
func (s *BasegramaticaListener) ExitFunciones(ctx *FuncionesContext) {}

// EnterMain is called when production main is entered.
func (s *BasegramaticaListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BasegramaticaListener) ExitMain(ctx *MainContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BasegramaticaListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BasegramaticaListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BasegramaticaListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BasegramaticaListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterControl is called when production control is entered.
func (s *BasegramaticaListener) EnterControl(ctx *ControlContext) {}

// ExitControl is called when production control is exited.
func (s *BasegramaticaListener) ExitControl(ctx *ControlContext) {}

// EnterDeclaracion is called when production declaracion is entered.
func (s *BasegramaticaListener) EnterDeclaracion(ctx *DeclaracionContext) {}

// ExitDeclaracion is called when production declaracion is exited.
func (s *BasegramaticaListener) ExitDeclaracion(ctx *DeclaracionContext) {}

// EnterTipovariable is called when production tipovariable is entered.
func (s *BasegramaticaListener) EnterTipovariable(ctx *TipovariableContext) {}

// ExitTipovariable is called when production tipovariable is exited.
func (s *BasegramaticaListener) ExitTipovariable(ctx *TipovariableContext) {}

// EnterIdentificadores is called when production identificadores is entered.
func (s *BasegramaticaListener) EnterIdentificadores(ctx *IdentificadoresContext) {}

// ExitIdentificadores is called when production identificadores is exited.
func (s *BasegramaticaListener) ExitIdentificadores(ctx *IdentificadoresContext) {}

// EnterValores is called when production valores is entered.
func (s *BasegramaticaListener) EnterValores(ctx *ValoresContext) {}

// ExitValores is called when production valores is exited.
func (s *BasegramaticaListener) ExitValores(ctx *ValoresContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BasegramaticaListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BasegramaticaListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterImpresion is called when production impresion is entered.
func (s *BasegramaticaListener) EnterImpresion(ctx *ImpresionContext) {}

// ExitImpresion is called when production impresion is exited.
func (s *BasegramaticaListener) ExitImpresion(ctx *ImpresionContext) {}

// EnterImpresionexpresion is called when production impresionexpresion is entered.
func (s *BasegramaticaListener) EnterImpresionexpresion(ctx *ImpresionexpresionContext) {}

// ExitImpresionexpresion is called when production impresionexpresion is exited.
func (s *BasegramaticaListener) ExitImpresionexpresion(ctx *ImpresionexpresionContext) {}

// EnterExpcoma is called when production expcoma is entered.
func (s *BasegramaticaListener) EnterExpcoma(ctx *ExpcomaContext) {}

// ExitExpcoma is called when production expcoma is exited.
func (s *BasegramaticaListener) ExitExpcoma(ctx *ExpcomaContext) {}

// EnterCondicionales is called when production condicionales is entered.
func (s *BasegramaticaListener) EnterCondicionales(ctx *CondicionalesContext) {}

// ExitCondicionales is called when production condicionales is exited.
func (s *BasegramaticaListener) ExitCondicionales(ctx *CondicionalesContext) {}

// EnterFuncionif is called when production funcionif is entered.
func (s *BasegramaticaListener) EnterFuncionif(ctx *FuncionifContext) {}

// ExitFuncionif is called when production funcionif is exited.
func (s *BasegramaticaListener) ExitFuncionif(ctx *FuncionifContext) {}

// EnterFuncionelseif is called when production funcionelseif is entered.
func (s *BasegramaticaListener) EnterFuncionelseif(ctx *FuncionelseifContext) {}

// ExitFuncionelseif is called when production funcionelseif is exited.
func (s *BasegramaticaListener) ExitFuncionelseif(ctx *FuncionelseifContext) {}

// EnterListaelseif is called when production listaelseif is entered.
func (s *BasegramaticaListener) EnterListaelseif(ctx *ListaelseifContext) {}

// ExitListaelseif is called when production listaelseif is exited.
func (s *BasegramaticaListener) ExitListaelseif(ctx *ListaelseifContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BasegramaticaListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BasegramaticaListener) ExitBloque(ctx *BloqueContext) {}

// EnterBucles is called when production bucles is entered.
func (s *BasegramaticaListener) EnterBucles(ctx *BuclesContext) {}

// ExitBucles is called when production bucles is exited.
func (s *BasegramaticaListener) ExitBucles(ctx *BuclesContext) {}

// EnterFwhile is called when production fwhile is entered.
func (s *BasegramaticaListener) EnterFwhile(ctx *FwhileContext) {}

// ExitFwhile is called when production fwhile is exited.
func (s *BasegramaticaListener) ExitFwhile(ctx *FwhileContext) {}

// EnterFloop is called when production floop is entered.
func (s *BasegramaticaListener) EnterFloop(ctx *FloopContext) {}

// ExitFloop is called when production floop is exited.
func (s *BasegramaticaListener) ExitFloop(ctx *FloopContext) {}
