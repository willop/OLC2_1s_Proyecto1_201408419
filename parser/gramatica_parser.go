// Code generated from gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "proyecto1/Interfaces"
import "proyecto1/Expresion"
import "proyecto1/Instruccion"

//import arrayList "github.com/colegno/arraylist"
//import "proyecto1/Expresion"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 70, 227,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3,
	3, 7, 3, 27, 10, 3, 12, 3, 14, 3, 30, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 40, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 76, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 90, 10, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 110, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 5, 9, 144, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 7, 9, 188, 10, 9, 12, 9, 14, 9, 191, 11, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	206, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 222, 10, 11, 12, 11, 14,
	11, 225, 11, 11, 3, 11, 2, 4, 16, 20, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 2, 2, 2, 256, 2, 22, 3, 2, 2, 2, 4, 28, 3, 2, 2, 2, 6, 39, 3, 2, 2,
	2, 8, 75, 3, 2, 2, 2, 10, 89, 3, 2, 2, 2, 12, 91, 3, 2, 2, 2, 14, 109,
	3, 2, 2, 2, 16, 143, 3, 2, 2, 2, 18, 205, 3, 2, 2, 2, 20, 207, 3, 2, 2,
	2, 22, 23, 5, 4, 3, 2, 23, 24, 7, 2, 2, 3, 24, 3, 3, 2, 2, 2, 25, 27, 5,
	6, 4, 2, 26, 25, 3, 2, 2, 2, 27, 30, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28,
	29, 3, 2, 2, 2, 29, 5, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 31, 32, 5, 16, 9,
	2, 32, 33, 8, 4, 1, 2, 33, 40, 3, 2, 2, 2, 34, 40, 5, 18, 10, 2, 35, 36,
	5, 8, 5, 2, 36, 37, 8, 4, 1, 2, 37, 40, 3, 2, 2, 2, 38, 40, 5, 12, 7, 2,
	39, 31, 3, 2, 2, 2, 39, 34, 3, 2, 2, 2, 39, 35, 3, 2, 2, 2, 39, 38, 3,
	2, 2, 2, 40, 7, 3, 2, 2, 2, 41, 42, 7, 44, 2, 2, 42, 43, 7, 45, 2, 2, 43,
	44, 7, 67, 2, 2, 44, 45, 7, 14, 2, 2, 45, 46, 5, 10, 6, 2, 46, 47, 7, 22,
	2, 2, 47, 48, 5, 16, 9, 2, 48, 49, 7, 16, 2, 2, 49, 50, 8, 5, 1, 2, 50,
	76, 3, 2, 2, 2, 51, 52, 7, 44, 2, 2, 52, 53, 7, 45, 2, 2, 53, 54, 7, 67,
	2, 2, 54, 55, 7, 22, 2, 2, 55, 56, 5, 16, 9, 2, 56, 57, 7, 16, 2, 2, 57,
	58, 8, 5, 1, 2, 58, 76, 3, 2, 2, 2, 59, 60, 7, 44, 2, 2, 60, 61, 7, 67,
	2, 2, 61, 62, 7, 14, 2, 2, 62, 63, 5, 10, 6, 2, 63, 64, 7, 22, 2, 2, 64,
	65, 5, 16, 9, 2, 65, 66, 7, 16, 2, 2, 66, 67, 8, 5, 1, 2, 67, 76, 3, 2,
	2, 2, 68, 69, 7, 44, 2, 2, 69, 70, 7, 67, 2, 2, 70, 71, 7, 22, 2, 2, 71,
	72, 5, 16, 9, 2, 72, 73, 7, 16, 2, 2, 73, 74, 8, 5, 1, 2, 74, 76, 3, 2,
	2, 2, 75, 41, 3, 2, 2, 2, 75, 51, 3, 2, 2, 2, 75, 59, 3, 2, 2, 2, 75, 68,
	3, 2, 2, 2, 76, 9, 3, 2, 2, 2, 77, 78, 7, 35, 2, 2, 78, 90, 8, 6, 1, 2,
	79, 80, 7, 36, 2, 2, 80, 90, 8, 6, 1, 2, 81, 82, 7, 42, 2, 2, 82, 90, 8,
	6, 1, 2, 83, 84, 7, 40, 2, 2, 84, 90, 8, 6, 1, 2, 85, 86, 7, 41, 2, 2,
	86, 90, 8, 6, 1, 2, 87, 88, 7, 43, 2, 2, 88, 90, 8, 6, 1, 2, 89, 77, 3,
	2, 2, 2, 89, 79, 3, 2, 2, 2, 89, 81, 3, 2, 2, 2, 89, 83, 3, 2, 2, 2, 89,
	85, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 90, 11, 3, 2, 2, 2, 91, 92, 7, 67,
	2, 2, 92, 93, 7, 22, 2, 2, 93, 94, 5, 16, 9, 2, 94, 95, 7, 16, 2, 2, 95,
	13, 3, 2, 2, 2, 96, 97, 7, 65, 2, 2, 97, 110, 8, 8, 1, 2, 98, 99, 7, 66,
	2, 2, 99, 110, 8, 8, 1, 2, 100, 101, 7, 69, 2, 2, 101, 110, 8, 8, 1, 2,
	102, 103, 7, 70, 2, 2, 103, 110, 8, 8, 1, 2, 104, 105, 7, 49, 2, 2, 105,
	110, 8, 8, 1, 2, 106, 107, 7, 50, 2, 2, 107, 110, 8, 8, 1, 2, 108, 110,
	7, 67, 2, 2, 109, 96, 3, 2, 2, 2, 109, 98, 3, 2, 2, 2, 109, 100, 3, 2,
	2, 2, 109, 102, 3, 2, 2, 2, 109, 104, 3, 2, 2, 2, 109, 106, 3, 2, 2, 2,
	109, 108, 3, 2, 2, 2, 110, 15, 3, 2, 2, 2, 111, 112, 8, 9, 1, 2, 112, 113,
	7, 24, 2, 2, 113, 144, 5, 16, 9, 23, 114, 115, 7, 37, 2, 2, 115, 116, 7,
	20, 2, 2, 116, 117, 5, 16, 9, 2, 117, 118, 7, 15, 2, 2, 118, 119, 5, 16,
	9, 2, 119, 120, 7, 21, 2, 2, 120, 144, 3, 2, 2, 2, 121, 122, 7, 39, 2,
	2, 122, 123, 7, 20, 2, 2, 123, 124, 5, 16, 9, 2, 124, 125, 7, 15, 2, 2,
	125, 126, 5, 16, 9, 2, 126, 127, 7, 21, 2, 2, 127, 144, 3, 2, 2, 2, 128,
	129, 7, 20, 2, 2, 129, 130, 5, 16, 9, 2, 130, 131, 7, 21, 2, 2, 131, 144,
	3, 2, 2, 2, 132, 133, 5, 14, 8, 2, 133, 134, 7, 47, 2, 2, 134, 135, 7,
	35, 2, 2, 135, 144, 3, 2, 2, 2, 136, 137, 5, 14, 8, 2, 137, 138, 7, 47,
	2, 2, 138, 139, 7, 36, 2, 2, 139, 144, 3, 2, 2, 2, 140, 141, 5, 14, 8,
	2, 141, 142, 8, 9, 1, 2, 142, 144, 3, 2, 2, 2, 143, 111, 3, 2, 2, 2, 143,
	114, 3, 2, 2, 2, 143, 121, 3, 2, 2, 2, 143, 128, 3, 2, 2, 2, 143, 132,
	3, 2, 2, 2, 143, 136, 3, 2, 2, 2, 143, 140, 3, 2, 2, 2, 144, 189, 3, 2,
	2, 2, 145, 146, 12, 22, 2, 2, 146, 147, 7, 25, 2, 2, 147, 188, 5, 16, 9,
	23, 148, 149, 12, 21, 2, 2, 149, 150, 7, 24, 2, 2, 150, 188, 5, 16, 9,
	22, 151, 152, 12, 20, 2, 2, 152, 153, 7, 26, 2, 2, 153, 188, 5, 16, 9,
	21, 154, 155, 12, 19, 2, 2, 155, 156, 7, 23, 2, 2, 156, 188, 5, 16, 9,
	20, 157, 158, 12, 16, 2, 2, 158, 159, 7, 27, 2, 2, 159, 188, 5, 16, 9,
	17, 160, 161, 12, 15, 2, 2, 161, 162, 7, 17, 2, 2, 162, 188, 5, 16, 9,
	16, 163, 164, 12, 14, 2, 2, 164, 165, 7, 18, 2, 2, 165, 188, 5, 16, 9,
	15, 166, 167, 12, 13, 2, 2, 167, 168, 7, 8, 2, 2, 168, 188, 5, 16, 9, 14,
	169, 170, 12, 12, 2, 2, 170, 171, 7, 9, 2, 2, 171, 188, 5, 16, 9, 13, 172,
	173, 12, 11, 2, 2, 173, 174, 7, 6, 2, 2, 174, 188, 5, 16, 9, 12, 175, 176,
	12, 10, 2, 2, 176, 177, 7, 7, 2, 2, 177, 188, 5, 16, 9, 11, 178, 179, 12,
	9, 2, 2, 179, 180, 7, 4, 2, 2, 180, 188, 5, 16, 9, 10, 181, 182, 12, 8,
	2, 2, 182, 183, 7, 5, 2, 2, 183, 188, 5, 16, 9, 9, 184, 185, 12, 7, 2,
	2, 185, 186, 7, 30, 2, 2, 186, 188, 5, 16, 9, 8, 187, 145, 3, 2, 2, 2,
	187, 148, 3, 2, 2, 2, 187, 151, 3, 2, 2, 2, 187, 154, 3, 2, 2, 2, 187,
	157, 3, 2, 2, 2, 187, 160, 3, 2, 2, 2, 187, 163, 3, 2, 2, 2, 187, 166,
	3, 2, 2, 2, 187, 169, 3, 2, 2, 2, 187, 172, 3, 2, 2, 2, 187, 175, 3, 2,
	2, 2, 187, 178, 3, 2, 2, 2, 187, 181, 3, 2, 2, 2, 187, 184, 3, 2, 2, 2,
	188, 191, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190,
	17, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 192, 193, 7, 48, 2, 2, 193, 194,
	7, 20, 2, 2, 194, 195, 5, 16, 9, 2, 195, 196, 7, 21, 2, 2, 196, 197, 7,
	16, 2, 2, 197, 198, 8, 10, 1, 2, 198, 206, 3, 2, 2, 2, 199, 200, 7, 48,
	2, 2, 200, 201, 7, 20, 2, 2, 201, 202, 5, 16, 9, 2, 202, 203, 5, 20, 11,
	2, 203, 204, 7, 16, 2, 2, 204, 206, 3, 2, 2, 2, 205, 192, 3, 2, 2, 2, 205,
	199, 3, 2, 2, 2, 206, 19, 3, 2, 2, 2, 207, 208, 8, 11, 1, 2, 208, 209,
	7, 15, 2, 2, 209, 210, 5, 16, 9, 2, 210, 211, 7, 21, 2, 2, 211, 212, 7,
	16, 2, 2, 212, 213, 8, 11, 1, 2, 213, 223, 3, 2, 2, 2, 214, 215, 12, 4,
	2, 2, 215, 216, 7, 15, 2, 2, 216, 217, 5, 16, 9, 2, 217, 218, 7, 21, 2,
	2, 218, 219, 7, 16, 2, 2, 219, 220, 8, 11, 1, 2, 220, 222, 3, 2, 2, 2,
	221, 214, 3, 2, 2, 2, 222, 225, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223,
	224, 3, 2, 2, 2, 224, 21, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 12, 28, 39,
	75, 89, 109, 143, 187, 189, 205, 223,
}
var literalNames = []string{
	"", "'->'", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'{'", "'}'",
	"'['", "']'", "':'", "','", "';'", "'<'", "'>'", "'.'", "'('", "')'", "'='",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'", "'!'", "'?'", "", "",
	"", "'i64'", "'f64'", "'pow'", "'vec'", "'powf'", "'bool'", "'char'", "",
	"'usize'", "'let'", "'mut'", "'struct'", "'as'", "'println!'", "'true'",
	"'false'", "'fn'", "'return'", "'abs'", "'sqrt'", "'to_string()'", "'clone()'",
	"'new'", "'len'", "'push'", "'remove'", "'contains'", "'insert'", "'capacity'",
	"'witch_capacity'",
}
var symbolicNames = []string{
	"", "TK_flecha", "TK_or", "TK_and", "TK_igualacion", "TK_diferente", "TK_mayor_igual",
	"TK_menor_igual", "TK_corchete_apertura", "TK_corchete_cierre", "TK_llave_apertura",
	"TK_llave_cierre", "TK_dosPuntos", "TK_coma", "TK_pcoma", "TK_menor", "TK_mayor",
	"TK_punto", "TK_par_apertura", "TK_par_cierre", "TK_igual", "TK_suma",
	"TK_menos", "TK_por", "TK_diagonal", "TK_porcentaje", "TK_linea", "TK_amp",
	"TK_sig_admiracion", "TK_sig_interrogacion", "WHITESPACE", "TK_comentario_multi",
	"TK_comentario_lineal", "TKR_numericos_enteros", "TKR_numericos_flotantes",
	"TKR_pow", "TKR_vec", "TKR_powf", "TKR_bool", "TKR_char", "TKR_String",
	"TKR_usize", "TKR_let", "TKR_mut", "TKR_struct", "TKR_as", "TKR_println",
	"TKR_true", "TKR_false", "TKR_fn", "TKR_return", "TKR_abs", "TKR_sqrt",
	"TKR_to_string", "TKR_clone", "TKR_new", "TKR_len", "TKR_push", "TKR_remove",
	"TKR_contains", "TKR_insert", "TKR_capacity", "TKR_with_capacity", "TK_entero",
	"TK_decimal", "TK_id", "Letra", "TK_cadena", "TK_caracter",
}

var ruleNames = []string{
	"start", "instrucciones", "instruccion", "declaracion", "tipovariable",
	"identificadores", "valores", "expresion", "impresion", "impresioncomas",
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

var temporal = Interfaces.SINTIPO

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
	gramaticaParserWHITESPACE              = 30
	gramaticaParserTK_comentario_multi     = 31
	gramaticaParserTK_comentario_lineal    = 32
	gramaticaParserTKR_numericos_enteros   = 33
	gramaticaParserTKR_numericos_flotantes = 34
	gramaticaParserTKR_pow                 = 35
	gramaticaParserTKR_vec                 = 36
	gramaticaParserTKR_powf                = 37
	gramaticaParserTKR_bool                = 38
	gramaticaParserTKR_char                = 39
	gramaticaParserTKR_String              = 40
	gramaticaParserTKR_usize               = 41
	gramaticaParserTKR_let                 = 42
	gramaticaParserTKR_mut                 = 43
	gramaticaParserTKR_struct              = 44
	gramaticaParserTKR_as                  = 45
	gramaticaParserTKR_println             = 46
	gramaticaParserTKR_true                = 47
	gramaticaParserTKR_false               = 48
	gramaticaParserTKR_fn                  = 49
	gramaticaParserTKR_return              = 50
	gramaticaParserTKR_abs                 = 51
	gramaticaParserTKR_sqrt                = 52
	gramaticaParserTKR_to_string           = 53
	gramaticaParserTKR_clone               = 54
	gramaticaParserTKR_new                 = 55
	gramaticaParserTKR_len                 = 56
	gramaticaParserTKR_push                = 57
	gramaticaParserTKR_remove              = 58
	gramaticaParserTKR_contains            = 59
	gramaticaParserTKR_insert              = 60
	gramaticaParserTKR_capacity            = 61
	gramaticaParserTKR_with_capacity       = 62
	gramaticaParserTK_entero               = 63
	gramaticaParserTK_decimal              = 64
	gramaticaParserTK_id                   = 65
	gramaticaParserLetra                   = 66
	gramaticaParserTK_cadena               = 67
	gramaticaParserTK_caracter             = 68
)

// gramaticaParser rules.
const (
	gramaticaParserRULE_start           = 0
	gramaticaParserRULE_instrucciones   = 1
	gramaticaParserRULE_instruccion     = 2
	gramaticaParserRULE_declaracion     = 3
	gramaticaParserRULE_tipovariable    = 4
	gramaticaParserRULE_identificadores = 5
	gramaticaParserRULE_valores         = 6
	gramaticaParserRULE_expresion       = 7
	gramaticaParserRULE_impresion       = 8
	gramaticaParserRULE_impresioncomas  = 9
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

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
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
		p.SetState(20)
		p.Instrucciones()
	}
	{
		p.SetState(21)
		p.Match(gramaticaParserEOF)
	}

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccionContext

	// SetE sets the e rule context list.
	SetE([]IInstruccionContext)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_instruccion IInstruccionContext
	e            []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetE() []IInstruccionContext { return s.e }

func (s *InstruccionesContext) SetE(v []IInstruccionContext) { s.e = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *gramaticaParser) Instrucciones() (localctx IInstruccionesContext) {
	this := p
	_ = this

	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gramaticaParserRULE_instrucciones)
	var _la int

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
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-18)&-(0x1f+1)) == 0 && ((1<<uint((_la-18)))&((1<<(gramaticaParserTK_par_apertura-18))|(1<<(gramaticaParserTK_menos-18))|(1<<(gramaticaParserTKR_pow-18))|(1<<(gramaticaParserTKR_powf-18))|(1<<(gramaticaParserTKR_let-18))|(1<<(gramaticaParserTKR_println-18))|(1<<(gramaticaParserTKR_true-18))|(1<<(gramaticaParserTKR_false-18)))) != 0) || (((_la-63)&-(0x1f+1)) == 0 && ((1<<uint((_la-63)))&((1<<(gramaticaParserTK_entero-63))|(1<<(gramaticaParserTK_decimal-63))|(1<<(gramaticaParserTK_id-63))|(1<<(gramaticaParserTK_cadena-63))|(1<<(gramaticaParserTK_caracter-63)))) != 0) {
		{
			p.SetState(23)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_expresion   IExpresionContext
	_declaracion IDeclaracionContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *InstruccionContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *InstruccionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *InstruccionContext) Impresion() IImpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpresionContext)
}

func (s *InstruccionContext) Declaracion() IDeclaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *InstruccionContext) Identificadores() IIdentificadoresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentificadoresContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentificadoresContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *gramaticaParser) Instruccion() (localctx IInstruccionContext) {
	this := p
	_ = this

	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gramaticaParserRULE_instruccion)

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

	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)

			var _x = p.expresion(0)

			localctx.(*InstruccionContext)._expresion = _x
		}
		fmt.Println("mensaje en instrucciones: ", localctx.(*InstruccionContext).Get_expresion().GetValorexpresion())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Impresion()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		fmt.Println("mensaje en declaracion: ", localctx.(*InstruccionContext).Get_declaracion().GetDecla())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(36)
			p.Identificadores()
		}

	}

	return localctx
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdd returns the idd token.
	GetIdd() antlr.Token

	// SetIdd sets the idd token.
	SetIdd(antlr.Token)

	// Get_tipovariable returns the _tipovariable rule contexts.
	Get_tipovariable() ITipovariableContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_tipovariable sets the _tipovariable rule contexts.
	Set_tipovariable(ITipovariableContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetDecla returns the decla attribute.
	GetDecla() interface{}

	// SetDecla sets the decla attribute.
	SetDecla(interface{})

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	decla         interface{}
	idd           antlr.Token
	_tipovariable ITipovariableContext
	_expresion    IExpresionContext
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracion
	return p
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) GetIdd() antlr.Token { return s.idd }

func (s *DeclaracionContext) SetIdd(v antlr.Token) { s.idd = v }

func (s *DeclaracionContext) Get_tipovariable() ITipovariableContext { return s._tipovariable }

func (s *DeclaracionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *DeclaracionContext) Set_tipovariable(v ITipovariableContext) { s._tipovariable = v }

func (s *DeclaracionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *DeclaracionContext) GetDecla() interface{} { return s.decla }

func (s *DeclaracionContext) SetDecla(v interface{}) { s.decla = v }

func (s *DeclaracionContext) TKR_let() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_let, 0)
}

func (s *DeclaracionContext) TKR_mut() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_mut, 0)
}

func (s *DeclaracionContext) TK_dosPuntos() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_dosPuntos, 0)
}

func (s *DeclaracionContext) Tipovariable() ITipovariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipovariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipovariableContext)
}

func (s *DeclaracionContext) TK_igual() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_igual, 0)
}

func (s *DeclaracionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclaracionContext) TK_pcoma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_pcoma, 0)
}

func (s *DeclaracionContext) TK_id() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_id, 0)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDeclaracion(s)
	}
}

func (s *DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDeclaracion(s)
	}
}

func (p *gramaticaParser) Declaracion() (localctx IDeclaracionContext) {
	this := p
	_ = this

	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gramaticaParserRULE_declaracion)

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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(40)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(41)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(42)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(43)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(44)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(45)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(46)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).decla = Instruccion.NuevaDeclaracion((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipovariable().GetTipovar(), localctx.(*DeclaracionContext).Get_expresion().GetValorexpresion(), true, false, false)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(50)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(51)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(52)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(53)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(54)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).decla = localctx.(*DeclaracionContext).Get_expresion().GetValorexpresion()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(58)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(59)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(60)
			p.Tipovariable()
		}
		{
			p.SetState(61)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(62)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(63)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).decla = localctx.(*DeclaracionContext).Get_expresion().GetValorexpresion()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(66)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(67)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(68)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(69)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(70)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).decla = localctx.(*DeclaracionContext).Get_expresion().GetValorexpresion()

	}

	return localctx
}

// ITipovariableContext is an interface to support dynamic dispatch.
type ITipovariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipovar returns the tipovar attribute.
	GetTipovar() Interfaces.Tipoexpresion

	// SetTipovar sets the tipovar attribute.
	SetTipovar(Interfaces.Tipoexpresion)

	// IsTipovariableContext differentiates from other interfaces.
	IsTipovariableContext()
}

type TipovariableContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	tipovar Interfaces.Tipoexpresion
}

func NewEmptyTipovariableContext() *TipovariableContext {
	var p = new(TipovariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_tipovariable
	return p
}

func (*TipovariableContext) IsTipovariableContext() {}

func NewTipovariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipovariableContext {
	var p = new(TipovariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_tipovariable

	return p
}

func (s *TipovariableContext) GetParser() antlr.Parser { return s.parser }

func (s *TipovariableContext) GetTipovar() Interfaces.Tipoexpresion { return s.tipovar }

func (s *TipovariableContext) SetTipovar(v Interfaces.Tipoexpresion) { s.tipovar = v }

func (s *TipovariableContext) TKR_numericos_enteros() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_numericos_enteros, 0)
}

func (s *TipovariableContext) TKR_numericos_flotantes() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_numericos_flotantes, 0)
}

func (s *TipovariableContext) TKR_String() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_String, 0)
}

func (s *TipovariableContext) TKR_bool() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_bool, 0)
}

func (s *TipovariableContext) TKR_char() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_char, 0)
}

func (s *TipovariableContext) TKR_usize() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_usize, 0)
}

func (s *TipovariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipovariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipovariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterTipovariable(s)
	}
}

func (s *TipovariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitTipovariable(s)
	}
}

func (p *gramaticaParser) Tipovariable() (localctx ITipovariableContext) {
	this := p
	_ = this

	localctx = NewTipovariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gramaticaParserRULE_tipovariable)

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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTKR_numericos_enteros:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Match(gramaticaParserTKR_numericos_enteros)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.INTEGER

	case gramaticaParserTKR_numericos_flotantes:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Match(gramaticaParserTKR_numericos_flotantes)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.FLOAT

	case gramaticaParserTKR_String:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)
			p.Match(gramaticaParserTKR_String)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.STRING

	case gramaticaParserTKR_bool:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(81)
			p.Match(gramaticaParserTKR_bool)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.BOOLEAN

	case gramaticaParserTKR_char:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(83)
			p.Match(gramaticaParserTKR_char)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.CHAR

	case gramaticaParserTKR_usize:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(85)
			p.Match(gramaticaParserTKR_usize)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.USIZE

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentificadoresContext is an interface to support dynamic dispatch.
type IIdentificadoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentificadoresContext differentiates from other interfaces.
	IsIdentificadoresContext()
}

type IdentificadoresContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentificadoresContext() *IdentificadoresContext {
	var p = new(IdentificadoresContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_identificadores
	return p
}

func (*IdentificadoresContext) IsIdentificadoresContext() {}

func NewIdentificadoresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentificadoresContext {
	var p = new(IdentificadoresContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_identificadores

	return p
}

func (s *IdentificadoresContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentificadoresContext) TK_id() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_id, 0)
}

func (s *IdentificadoresContext) TK_igual() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_igual, 0)
}

func (s *IdentificadoresContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *IdentificadoresContext) TK_pcoma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_pcoma, 0)
}

func (s *IdentificadoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentificadoresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentificadoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterIdentificadores(s)
	}
}

func (s *IdentificadoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitIdentificadores(s)
	}
}

func (p *gramaticaParser) Identificadores() (localctx IIdentificadoresContext) {
	this := p
	_ = this

	localctx = NewIdentificadoresContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gramaticaParserRULE_identificadores)

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
		p.SetState(89)
		p.Match(gramaticaParserTK_id)
	}
	{
		p.SetState(90)
		p.Match(gramaticaParserTK_igual)
	}
	{
		p.SetState(91)
		p.expresion(0)
	}
	{
		p.SetState(92)
		p.Match(gramaticaParserTK_pcoma)
	}

	return localctx
}

// IValoresContext is an interface to support dynamic dispatch.
type IValoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_entero returns the _TK_entero token.
	Get_TK_entero() antlr.Token

	// Get_TK_decimal returns the _TK_decimal token.
	Get_TK_decimal() antlr.Token

	// Get_TK_cadena returns the _TK_cadena token.
	Get_TK_cadena() antlr.Token

	// Get_TK_caracter returns the _TK_caracter token.
	Get_TK_caracter() antlr.Token

	// Set_TK_entero sets the _TK_entero token.
	Set_TK_entero(antlr.Token)

	// Set_TK_decimal sets the _TK_decimal token.
	Set_TK_decimal(antlr.Token)

	// Set_TK_cadena sets the _TK_cadena token.
	Set_TK_cadena(antlr.Token)

	// Set_TK_caracter sets the _TK_caracter token.
	Set_TK_caracter(antlr.Token)

	// GetLit returns the lit attribute.
	GetLit() Interfaces.Expresion

	// SetLit sets the lit attribute.
	SetLit(Interfaces.Expresion)

	// IsValoresContext differentiates from other interfaces.
	IsValoresContext()
}

type ValoresContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	lit          Interfaces.Expresion
	_TK_entero   antlr.Token
	_TK_decimal  antlr.Token
	_TK_cadena   antlr.Token
	_TK_caracter antlr.Token
}

func NewEmptyValoresContext() *ValoresContext {
	var p = new(ValoresContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_valores
	return p
}

func (*ValoresContext) IsValoresContext() {}

func NewValoresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValoresContext {
	var p = new(ValoresContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_valores

	return p
}

func (s *ValoresContext) GetParser() antlr.Parser { return s.parser }

func (s *ValoresContext) Get_TK_entero() antlr.Token { return s._TK_entero }

func (s *ValoresContext) Get_TK_decimal() antlr.Token { return s._TK_decimal }

func (s *ValoresContext) Get_TK_cadena() antlr.Token { return s._TK_cadena }

func (s *ValoresContext) Get_TK_caracter() antlr.Token { return s._TK_caracter }

func (s *ValoresContext) Set_TK_entero(v antlr.Token) { s._TK_entero = v }

func (s *ValoresContext) Set_TK_decimal(v antlr.Token) { s._TK_decimal = v }

func (s *ValoresContext) Set_TK_cadena(v antlr.Token) { s._TK_cadena = v }

func (s *ValoresContext) Set_TK_caracter(v antlr.Token) { s._TK_caracter = v }

func (s *ValoresContext) GetLit() Interfaces.Expresion { return s.lit }

func (s *ValoresContext) SetLit(v Interfaces.Expresion) { s.lit = v }

func (s *ValoresContext) TK_entero() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_entero, 0)
}

func (s *ValoresContext) TK_decimal() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_decimal, 0)
}

func (s *ValoresContext) TK_cadena() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_cadena, 0)
}

func (s *ValoresContext) TK_caracter() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_caracter, 0)
}

func (s *ValoresContext) TKR_true() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_true, 0)
}

func (s *ValoresContext) TKR_false() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_false, 0)
}

func (s *ValoresContext) TK_id() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_id, 0)
}

func (s *ValoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValoresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterValores(s)
	}
}

func (s *ValoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitValores(s)
	}
}

func (p *gramaticaParser) Valores() (localctx IValoresContext) {
	this := p
	_ = this

	localctx = NewValoresContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gramaticaParserRULE_valores)

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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTK_entero:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)

			var _m = p.Match(gramaticaParserTK_entero)

			localctx.(*ValoresContext)._TK_entero = _m
		}

		numero, err := strconv.Atoi((func() string {
			if localctx.(*ValoresContext).Get_TK_entero() == nil {
				return ""
			} else {
				return localctx.(*ValoresContext).Get_TK_entero().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*ValoresContext).lit = Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)

	case gramaticaParserTK_decimal:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)

			var _m = p.Match(gramaticaParserTK_decimal)

			localctx.(*ValoresContext)._TK_decimal = _m
		}
		decimal, err := strconv.ParseFloat((func() string {
			if localctx.(*ValoresContext).Get_TK_decimal() == nil {
				return ""
			} else {
				return localctx.(*ValoresContext).Get_TK_decimal().GetText()
			}
		}()), 8)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*ValoresContext).lit = Expresion.NuevoPrimitivo(decimal, Interfaces.FLOAT)

	case gramaticaParserTK_cadena:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)

			var _m = p.Match(gramaticaParserTK_cadena)

			localctx.(*ValoresContext)._TK_cadena = _m
		}

		str := (func() string {
			if localctx.(*ValoresContext).Get_TK_cadena() == nil {
				return ""
			} else {
				return localctx.(*ValoresContext).Get_TK_cadena().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ValoresContext).Get_TK_cadena() == nil {
				return ""
			} else {
				return localctx.(*ValoresContext).Get_TK_cadena().GetText()
			}
		}()))-1]
		localctx.(*ValoresContext).lit = Expresion.NuevoPrimitivo(str, Interfaces.STRING)

	case gramaticaParserTK_caracter:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(100)

			var _m = p.Match(gramaticaParserTK_caracter)

			localctx.(*ValoresContext)._TK_caracter = _m
		}
		str := (func() string {
			if localctx.(*ValoresContext).Get_TK_caracter() == nil {
				return ""
			} else {
				return localctx.(*ValoresContext).Get_TK_caracter().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ValoresContext).Get_TK_caracter() == nil {
				return ""
			} else {
				return localctx.(*ValoresContext).Get_TK_caracter().GetText()
			}
		}()))-1]
		localctx.(*ValoresContext).lit = Expresion.NuevoPrimitivo(str, Interfaces.CHAR)

	case gramaticaParserTKR_true:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(102)
			p.Match(gramaticaParserTKR_true)
		}
		localctx.(*ValoresContext).lit = Expresion.NuevoPrimitivo(true, Interfaces.BOOLEAN)

	case gramaticaParserTKR_false:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(104)
			p.Match(gramaticaParserTKR_false)
		}
		localctx.(*ValoresContext).lit = Expresion.NuevoPrimitivo(false, Interfaces.BOOLEAN)

	case gramaticaParserTK_id:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(106)
			p.Match(gramaticaParserTK_id)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetVall returns the vall rule contexts.
	GetVall() IValoresContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetVall sets the vall rule contexts.
	SetVall(IValoresContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetValorexpresion returns the valorexpresion attribute.
	GetValorexpresion() Interfaces.Expresion

	// SetValorexpresion sets the valorexpresion attribute.
	SetValorexpresion(Interfaces.Expresion)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	valorexpresion Interfaces.Expresion
	e1             IExpresionContext
	vall           IValoresContext
	op             antlr.Token
	e2             IExpresionContext
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) GetOp() antlr.Token { return s.op }

func (s *ExpresionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpresionContext) GetE1() IExpresionContext { return s.e1 }

func (s *ExpresionContext) GetVall() IValoresContext { return s.vall }

func (s *ExpresionContext) GetE2() IExpresionContext { return s.e2 }

func (s *ExpresionContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *ExpresionContext) SetVall(v IValoresContext) { s.vall = v }

func (s *ExpresionContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *ExpresionContext) GetValorexpresion() Interfaces.Expresion { return s.valorexpresion }

func (s *ExpresionContext) SetValorexpresion(v Interfaces.Expresion) { s.valorexpresion = v }

func (s *ExpresionContext) TK_menos() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_menos, 0)
}

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) TKR_pow() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_pow, 0)
}

func (s *ExpresionContext) TK_par_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_apertura, 0)
}

func (s *ExpresionContext) TK_coma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_coma, 0)
}

func (s *ExpresionContext) TK_par_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_cierre, 0)
}

func (s *ExpresionContext) TKR_powf() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_powf, 0)
}

func (s *ExpresionContext) Valores() IValoresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValoresContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValoresContext)
}

func (s *ExpresionContext) TKR_as() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_as, 0)
}

func (s *ExpresionContext) TKR_numericos_enteros() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_numericos_enteros, 0)
}

func (s *ExpresionContext) TKR_numericos_flotantes() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_numericos_flotantes, 0)
}

func (s *ExpresionContext) TK_por() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_por, 0)
}

func (s *ExpresionContext) TK_diagonal() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_diagonal, 0)
}

func (s *ExpresionContext) TK_suma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_suma, 0)
}

func (s *ExpresionContext) TK_porcentaje() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_porcentaje, 0)
}

func (s *ExpresionContext) TK_menor() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_menor, 0)
}

func (s *ExpresionContext) TK_mayor() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_mayor, 0)
}

func (s *ExpresionContext) TK_mayor_igual() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_mayor_igual, 0)
}

func (s *ExpresionContext) TK_menor_igual() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_menor_igual, 0)
}

func (s *ExpresionContext) TK_igualacion() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_igualacion, 0)
}

func (s *ExpresionContext) TK_diferente() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_diferente, 0)
}

func (s *ExpresionContext) TK_or() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_or, 0)
}

func (s *ExpresionContext) TK_and() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_and, 0)
}

func (s *ExpresionContext) TK_sig_admiracion() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_sig_admiracion, 0)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *gramaticaParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *gramaticaParser) expresion(_p int) (localctx IExpresionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, gramaticaParserRULE_expresion, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(110)
			p.Match(gramaticaParserTK_menos)
		}
		{
			p.SetState(111)
			p.expresion(21)
		}

	case 2:
		{
			p.SetState(112)
			p.Match(gramaticaParserTKR_pow)
		}
		{
			p.SetState(113)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(114)
			p.expresion(0)
		}
		{
			p.SetState(115)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(116)
			p.expresion(0)
		}
		{
			p.SetState(117)
			p.Match(gramaticaParserTK_par_cierre)
		}

	case 3:
		{
			p.SetState(119)
			p.Match(gramaticaParserTKR_powf)
		}
		{
			p.SetState(120)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(121)
			p.expresion(0)
		}
		{
			p.SetState(122)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(123)
			p.expresion(0)
		}
		{
			p.SetState(124)
			p.Match(gramaticaParserTK_par_cierre)
		}

	case 4:
		{
			p.SetState(126)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(127)
			p.expresion(0)
		}
		{
			p.SetState(128)
			p.Match(gramaticaParserTK_par_cierre)
		}

	case 5:
		{
			p.SetState(130)
			p.Valores()
		}
		{
			p.SetState(131)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(132)
			p.Match(gramaticaParserTKR_numericos_enteros)
		}

	case 6:
		{
			p.SetState(134)
			p.Valores()
		}
		{
			p.SetState(135)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(136)
			p.Match(gramaticaParserTKR_numericos_flotantes)
		}

	case 7:
		{
			p.SetState(138)

			var _x = p.Valores()

			localctx.(*ExpresionContext).vall = _x
		}
		localctx.(*ExpresionContext).valorexpresion = localctx.(*ExpresionContext).GetVall().GetLit()
		fmt.Println(localctx.(*ExpresionContext).valorexpresion)

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(185)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(144)

					var _m = p.Match(gramaticaParserTK_por)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(145)

					var _x = p.expresion(21)

					localctx.(*ExpresionContext).e2 = _x
				}

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(147)

					var _m = p.Match(gramaticaParserTK_menos)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(148)

					var _x = p.expresion(20)

					localctx.(*ExpresionContext).e2 = _x
				}

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(150)

					var _m = p.Match(gramaticaParserTK_diagonal)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(151)

					var _x = p.expresion(19)

					localctx.(*ExpresionContext).e2 = _x
				}

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(153)

					var _m = p.Match(gramaticaParserTK_suma)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(154)

					var _x = p.expresion(18)

					localctx.(*ExpresionContext).e2 = _x
				}

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(155)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(156)
					p.Match(gramaticaParserTK_porcentaje)
				}
				{
					p.SetState(157)
					p.expresion(15)
				}

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(159)
					p.Match(gramaticaParserTK_menor)
				}
				{
					p.SetState(160)
					p.expresion(14)
				}

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(161)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(162)
					p.Match(gramaticaParserTK_mayor)
				}
				{
					p.SetState(163)
					p.expresion(13)
				}

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(164)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(165)
					p.Match(gramaticaParserTK_mayor_igual)
				}
				{
					p.SetState(166)
					p.expresion(12)
				}

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(167)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(168)
					p.Match(gramaticaParserTK_menor_igual)
				}
				{
					p.SetState(169)
					p.expresion(11)
				}

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(171)
					p.Match(gramaticaParserTK_igualacion)
				}
				{
					p.SetState(172)
					p.expresion(10)
				}

			case 11:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(173)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(174)
					p.Match(gramaticaParserTK_diferente)
				}
				{
					p.SetState(175)
					p.expresion(9)
				}

			case 12:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(176)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(177)
					p.Match(gramaticaParserTK_or)
				}
				{
					p.SetState(178)
					p.expresion(8)
				}

			case 13:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(179)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(180)
					p.Match(gramaticaParserTK_and)
				}
				{
					p.SetState(181)
					p.expresion(7)
				}

			case 14:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(182)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(183)
					p.Match(gramaticaParserTK_sig_admiracion)
				}
				{
					p.SetState(184)
					p.expresion(6)
				}

			}

		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IImpresionContext is an interface to support dynamic dispatch.
type IImpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImpresionContext differentiates from other interfaces.
	IsImpresionContext()
}

type ImpresionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImpresionContext() *ImpresionContext {
	var p = new(ImpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_impresion
	return p
}

func (*ImpresionContext) IsImpresionContext() {}

func NewImpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpresionContext {
	var p = new(ImpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_impresion

	return p
}

func (s *ImpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpresionContext) TKR_println() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_println, 0)
}

func (s *ImpresionContext) TK_par_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_apertura, 0)
}

func (s *ImpresionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ImpresionContext) TK_par_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_cierre, 0)
}

func (s *ImpresionContext) TK_pcoma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_pcoma, 0)
}

func (s *ImpresionContext) Impresioncomas() IImpresioncomasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpresioncomasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpresioncomasContext)
}

func (s *ImpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterImpresion(s)
	}
}

func (s *ImpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitImpresion(s)
	}
}

func (p *gramaticaParser) Impresion() (localctx IImpresionContext) {
	this := p
	_ = this

	localctx = NewImpresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gramaticaParserRULE_impresion)

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

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(191)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(192)
			p.expresion(0)
		}
		{
			p.SetState(193)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(194)
			p.Match(gramaticaParserTK_pcoma)
		}
		fmt.Println("Impresion")

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(198)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(199)
			p.expresion(0)
		}
		{
			p.SetState(200)
			p.impresioncomas(0)
		}
		{
			p.SetState(201)
			p.Match(gramaticaParserTK_pcoma)
		}

	}

	return localctx
}

// IImpresioncomasContext is an interface to support dynamic dispatch.
type IImpresioncomasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImpresioncomasContext differentiates from other interfaces.
	IsImpresioncomasContext()
}

type ImpresioncomasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImpresioncomasContext() *ImpresioncomasContext {
	var p = new(ImpresioncomasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_impresioncomas
	return p
}

func (*ImpresioncomasContext) IsImpresioncomasContext() {}

func NewImpresioncomasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpresioncomasContext {
	var p = new(ImpresioncomasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_impresioncomas

	return p
}

func (s *ImpresioncomasContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpresioncomasContext) TK_coma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_coma, 0)
}

func (s *ImpresioncomasContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ImpresioncomasContext) TK_par_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_cierre, 0)
}

func (s *ImpresioncomasContext) TK_pcoma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_pcoma, 0)
}

func (s *ImpresioncomasContext) Impresioncomas() IImpresioncomasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpresioncomasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpresioncomasContext)
}

func (s *ImpresioncomasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpresioncomasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImpresioncomasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterImpresioncomas(s)
	}
}

func (s *ImpresioncomasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitImpresioncomas(s)
	}
}

func (p *gramaticaParser) Impresioncomas() (localctx IImpresioncomasContext) {
	return p.impresioncomas(0)
}

func (p *gramaticaParser) impresioncomas(_p int) (localctx IImpresioncomasContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewImpresioncomasContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IImpresioncomasContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, gramaticaParserRULE_impresioncomas, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(gramaticaParserTK_coma)
	}
	{
		p.SetState(207)
		p.expresion(0)
	}
	{
		p.SetState(208)
		p.Match(gramaticaParserTK_par_cierre)
	}
	{
		p.SetState(209)
		p.Match(gramaticaParserTK_pcoma)
	}
	fmt.Println("Impresion especial")

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewImpresioncomasContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_impresioncomas)
			p.SetState(212)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(213)
				p.Match(gramaticaParserTK_coma)
			}
			{
				p.SetState(214)
				p.expresion(0)
			}
			{
				p.SetState(215)
				p.Match(gramaticaParserTK_par_cierre)
			}
			{
				p.SetState(216)
				p.Match(gramaticaParserTK_pcoma)
			}
			fmt.Println("Impresion especial")

		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

func (p *gramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	case 9:
		var t *ImpresioncomasContext = nil
		if localctx != nil {
			t = localctx.(*ImpresioncomasContext)
		}
		return p.Impresioncomas_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *gramaticaParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *gramaticaParser) Impresioncomas_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
