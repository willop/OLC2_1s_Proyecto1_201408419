// Code generated from gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 67, 15, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 4, 2, 4, 2, 2, 2, 12, 2, 6, 3, 2, 2, 2, 4, 9, 3, 2, 2, 2, 6, 7,
	5, 4, 3, 2, 7, 8, 7, 2, 2, 3, 8, 3, 3, 2, 2, 2, 9, 10, 7, 51, 2, 2, 10,
	11, 7, 20, 2, 2, 11, 12, 7, 32, 2, 2, 12, 13, 7, 21, 2, 2, 13, 5, 3, 2,
	2, 2, 2,
}
var literalNames = []string{
	"", "'->'", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'{'", "'}'",
	"'['", "']'", "':'", "','", "';'", "'<'", "'>'", "'.'", "'('", "')'", "'='",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'", "'!'", "'?'", "", "",
	"", "", "", "", "", "", "", "'i64'", "'f64'", "'bool'", "'char'", "", "'usize'",
	"'let'", "'mut'", "'struct'", "'as'", "'println!'", "'true'", "'false'",
	"'fn'", "'return'", "'abs'", "'sqrt'", "'to_string()'", "'clone()'", "'new'",
	"'len'", "'push'", "'remove'", "'contains'", "'insert'", "'capacity'",
	"'witch_capacity'",
}
var symbolicNames = []string{
	"", "TK_flecha", "TK_or", "TK_and", "TK_igualacion", "TK_diferente", "TK_mayor_igual",
	"TK_menor_igual", "TK_corchete_apertura", "TK_corchete_cierre", "TK_llave_apertura",
	"TK_llave_cierre", "TK_dosPuntos", "TK_coma", "TK_pcoma", "TK_menor", "TK_mayor",
	"TK_punto", "TK_par_apertura", "TK_par_cierre", "TK_igual", "TK_suma",
	"TK_menos", "TK_por", "TK_diagonal", "TK_porcentaje", "TK_linea", "TK_amp",
	"TK_sig_admiracion", "TK_sig_interrogacion", "TK_entero", "TK_decimal",
	"TK_id", "Letra", "TK_cadena", "TK_caracter", "WHITESPACE", "TK_comentario_multi",
	"TK_comentario_lineal", "TKR_numericos_enteros", "TKR_numericos_flotantes",
	"TKR_bool", "TKR_char", "TKR_String", "TKR_usize", "TKR_let", "TKR_mut",
	"TKR_struct", "TKR_as", "TKR_println", "TKR_true", "TKR_false", "TKR_fn",
	"TKR_return", "TKR_abs", "TKR_sqrt", "TKR_to_string", "TKR_clone", "TKR_new",
	"TKR_len", "TKR_push", "TKR_remove", "TKR_contains", "TKR_insert", "TKR_capacity",
	"TKR_with_capacity",
}

var ruleNames = []string{
	"start", "expresiones",
}

type gramaticaParser struct {
	*antlr.BaseParser
}

// NewgramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *gramaticaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewgramaticaParser(input antlr.TokenStream) *gramaticaParser {
	this := new(gramaticaParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "gramatica.g4"

	return this
}

// gramaticaParser tokens.
const (
	gramaticaParserEOF                     = antlr.TokenEOF
	gramaticaParserTK_flecha               = 1
	gramaticaParserTK_or                   = 2
	gramaticaParserTK_and                  = 3
	gramaticaParserTK_igualacion           = 4
	gramaticaParserTK_diferente            = 5
	gramaticaParserTK_mayor_igual          = 6
	gramaticaParserTK_menor_igual          = 7
	gramaticaParserTK_corchete_apertura    = 8
	gramaticaParserTK_corchete_cierre      = 9
	gramaticaParserTK_llave_apertura       = 10
	gramaticaParserTK_llave_cierre         = 11
	gramaticaParserTK_dosPuntos            = 12
	gramaticaParserTK_coma                 = 13
	gramaticaParserTK_pcoma                = 14
	gramaticaParserTK_menor                = 15
	gramaticaParserTK_mayor                = 16
	gramaticaParserTK_punto                = 17
	gramaticaParserTK_par_apertura         = 18
	gramaticaParserTK_par_cierre           = 19
	gramaticaParserTK_igual                = 20
	gramaticaParserTK_suma                 = 21
	gramaticaParserTK_menos                = 22
	gramaticaParserTK_por                  = 23
	gramaticaParserTK_diagonal             = 24
	gramaticaParserTK_porcentaje           = 25
	gramaticaParserTK_linea                = 26
	gramaticaParserTK_amp                  = 27
	gramaticaParserTK_sig_admiracion       = 28
	gramaticaParserTK_sig_interrogacion    = 29
	gramaticaParserTK_entero               = 30
	gramaticaParserTK_decimal              = 31
	gramaticaParserTK_id                   = 32
	gramaticaParserLetra                   = 33
	gramaticaParserTK_cadena               = 34
	gramaticaParserTK_caracter             = 35
	gramaticaParserWHITESPACE              = 36
	gramaticaParserTK_comentario_multi     = 37
	gramaticaParserTK_comentario_lineal    = 38
	gramaticaParserTKR_numericos_enteros   = 39
	gramaticaParserTKR_numericos_flotantes = 40
	gramaticaParserTKR_bool                = 41
	gramaticaParserTKR_char                = 42
	gramaticaParserTKR_String              = 43
	gramaticaParserTKR_usize               = 44
	gramaticaParserTKR_let                 = 45
	gramaticaParserTKR_mut                 = 46
	gramaticaParserTKR_struct              = 47
	gramaticaParserTKR_as                  = 48
	gramaticaParserTKR_println             = 49
	gramaticaParserTKR_true                = 50
	gramaticaParserTKR_false               = 51
	gramaticaParserTKR_fn                  = 52
	gramaticaParserTKR_return              = 53
	gramaticaParserTKR_abs                 = 54
	gramaticaParserTKR_sqrt                = 55
	gramaticaParserTKR_to_string           = 56
	gramaticaParserTKR_clone               = 57
	gramaticaParserTKR_new                 = 58
	gramaticaParserTKR_len                 = 59
	gramaticaParserTKR_push                = 60
	gramaticaParserTKR_remove              = 61
	gramaticaParserTKR_contains            = 62
	gramaticaParserTKR_insert              = 63
	gramaticaParserTKR_capacity            = 64
	gramaticaParserTKR_with_capacity       = 65
)

// gramaticaParser rules.
const (
	gramaticaParserRULE_start       = 0
	gramaticaParserRULE_expresiones = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expresiones() IExpresionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionesContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(gramaticaParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *gramaticaParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gramaticaParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Expresiones()
	}
	{
		p.SetState(5)
		p.Match(gramaticaParserEOF)
	}

	return localctx
}

// IExpresionesContext is an interface to support dynamic dispatch.
type IExpresionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpresionesContext differentiates from other interfaces.
	IsExpresionesContext()
}

type ExpresionesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionesContext() *ExpresionesContext {
	var p = new(ExpresionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_expresiones
	return p
}

func (*ExpresionesContext) IsExpresionesContext() {}

func NewExpresionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionesContext {
	var p = new(ExpresionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_expresiones

	return p
}

func (s *ExpresionesContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionesContext) TKR_println() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_println, 0)
}

func (s *ExpresionesContext) TK_par_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_apertura, 0)
}

func (s *ExpresionesContext) TK_entero() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_entero, 0)
}

func (s *ExpresionesContext) TK_par_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_cierre, 0)
}

func (s *ExpresionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExpresiones(s)
	}
}

func (s *ExpresionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExpresiones(s)
	}
}

func (p *gramaticaParser) Expresiones() (localctx IExpresionesContext) {
	this := p
	_ = this

	localctx = NewExpresionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gramaticaParserRULE_expresiones)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(7)
		p.Match(gramaticaParserTKR_println)
	}
	{
		p.SetState(8)
		p.Match(gramaticaParserTK_par_apertura)
	}
	{
		p.SetState(9)
		p.Match(gramaticaParserTK_entero)
	}
	{
		p.SetState(10)
		p.Match(gramaticaParserTK_par_cierre)
	}

	return localctx
}
