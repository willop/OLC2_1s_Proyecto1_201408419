// Code generated from gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "proyecto1/Interfaces"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 70, 199,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 7, 2, 24, 10,
	2, 12, 2, 14, 2, 27, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 37, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 65, 10, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 108, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 160, 10, 9, 12, 9, 14, 9,
	163, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 178, 10, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 7, 11, 194, 10, 11, 12, 11, 14, 11, 197, 11, 11, 3, 11, 2, 4, 16, 20,
	12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 4, 4, 2, 35, 36, 40, 43, 5,
	2, 49, 50, 65, 67, 69, 70, 2, 217, 2, 25, 3, 2, 2, 2, 4, 36, 3, 2, 2, 2,
	6, 64, 3, 2, 2, 2, 8, 66, 3, 2, 2, 2, 10, 68, 3, 2, 2, 2, 12, 71, 3, 2,
	2, 2, 14, 75, 3, 2, 2, 2, 16, 107, 3, 2, 2, 2, 18, 177, 3, 2, 2, 2, 20,
	179, 3, 2, 2, 2, 22, 24, 5, 4, 3, 2, 23, 22, 3, 2, 2, 2, 24, 27, 3, 2,
	2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 28, 3, 2, 2, 2, 27, 25,
	3, 2, 2, 2, 28, 29, 7, 2, 2, 3, 29, 3, 3, 2, 2, 2, 30, 31, 5, 16, 9, 2,
	31, 32, 8, 3, 1, 2, 32, 37, 3, 2, 2, 2, 33, 37, 5, 18, 10, 2, 34, 37, 5,
	6, 4, 2, 35, 37, 5, 12, 7, 2, 36, 30, 3, 2, 2, 2, 36, 33, 3, 2, 2, 2, 36,
	34, 3, 2, 2, 2, 36, 35, 3, 2, 2, 2, 37, 5, 3, 2, 2, 2, 38, 39, 7, 44, 2,
	2, 39, 40, 7, 45, 2, 2, 40, 41, 7, 67, 2, 2, 41, 42, 7, 14, 2, 2, 42, 43,
	5, 8, 5, 2, 43, 44, 5, 10, 6, 2, 44, 45, 7, 16, 2, 2, 45, 65, 3, 2, 2,
	2, 46, 47, 7, 44, 2, 2, 47, 48, 7, 45, 2, 2, 48, 49, 7, 67, 2, 2, 49, 50,
	5, 10, 6, 2, 50, 51, 7, 16, 2, 2, 51, 65, 3, 2, 2, 2, 52, 53, 7, 44, 2,
	2, 53, 54, 7, 67, 2, 2, 54, 55, 7, 14, 2, 2, 55, 56, 5, 8, 5, 2, 56, 57,
	5, 10, 6, 2, 57, 58, 7, 16, 2, 2, 58, 65, 3, 2, 2, 2, 59, 60, 7, 44, 2,
	2, 60, 61, 7, 67, 2, 2, 61, 62, 5, 10, 6, 2, 62, 63, 7, 16, 2, 2, 63, 65,
	3, 2, 2, 2, 64, 38, 3, 2, 2, 2, 64, 46, 3, 2, 2, 2, 64, 52, 3, 2, 2, 2,
	64, 59, 3, 2, 2, 2, 65, 7, 3, 2, 2, 2, 66, 67, 9, 2, 2, 2, 67, 9, 3, 2,
	2, 2, 68, 69, 7, 22, 2, 2, 69, 70, 5, 16, 9, 2, 70, 11, 3, 2, 2, 2, 71,
	72, 7, 67, 2, 2, 72, 73, 5, 10, 6, 2, 73, 74, 7, 16, 2, 2, 74, 13, 3, 2,
	2, 2, 75, 76, 9, 3, 2, 2, 76, 15, 3, 2, 2, 2, 77, 78, 8, 9, 1, 2, 78, 79,
	7, 24, 2, 2, 79, 108, 5, 16, 9, 23, 80, 81, 7, 37, 2, 2, 81, 82, 7, 20,
	2, 2, 82, 83, 5, 16, 9, 2, 83, 84, 7, 15, 2, 2, 84, 85, 5, 16, 9, 2, 85,
	86, 7, 21, 2, 2, 86, 108, 3, 2, 2, 2, 87, 88, 7, 39, 2, 2, 88, 89, 7, 20,
	2, 2, 89, 90, 5, 16, 9, 2, 90, 91, 7, 15, 2, 2, 91, 92, 5, 16, 9, 2, 92,
	93, 7, 21, 2, 2, 93, 108, 3, 2, 2, 2, 94, 95, 7, 20, 2, 2, 95, 96, 5, 16,
	9, 2, 96, 97, 7, 21, 2, 2, 97, 108, 3, 2, 2, 2, 98, 99, 5, 14, 8, 2, 99,
	100, 7, 47, 2, 2, 100, 101, 7, 35, 2, 2, 101, 108, 3, 2, 2, 2, 102, 103,
	5, 14, 8, 2, 103, 104, 7, 47, 2, 2, 104, 105, 7, 36, 2, 2, 105, 108, 3,
	2, 2, 2, 106, 108, 5, 14, 8, 2, 107, 77, 3, 2, 2, 2, 107, 80, 3, 2, 2,
	2, 107, 87, 3, 2, 2, 2, 107, 94, 3, 2, 2, 2, 107, 98, 3, 2, 2, 2, 107,
	102, 3, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 161, 3, 2, 2, 2, 109, 110,
	12, 22, 2, 2, 110, 111, 7, 23, 2, 2, 111, 112, 5, 16, 9, 23, 112, 113,
	8, 9, 1, 2, 113, 160, 3, 2, 2, 2, 114, 115, 12, 21, 2, 2, 115, 116, 7,
	24, 2, 2, 116, 117, 5, 16, 9, 22, 117, 118, 8, 9, 1, 2, 118, 160, 3, 2,
	2, 2, 119, 120, 12, 20, 2, 2, 120, 121, 7, 25, 2, 2, 121, 122, 5, 16, 9,
	21, 122, 123, 8, 9, 1, 2, 123, 160, 3, 2, 2, 2, 124, 125, 12, 19, 2, 2,
	125, 126, 7, 26, 2, 2, 126, 127, 5, 16, 9, 20, 127, 128, 8, 9, 1, 2, 128,
	160, 3, 2, 2, 2, 129, 130, 12, 16, 2, 2, 130, 131, 7, 27, 2, 2, 131, 160,
	5, 16, 9, 17, 132, 133, 12, 15, 2, 2, 133, 134, 7, 17, 2, 2, 134, 160,
	5, 16, 9, 16, 135, 136, 12, 14, 2, 2, 136, 137, 7, 18, 2, 2, 137, 160,
	5, 16, 9, 15, 138, 139, 12, 13, 2, 2, 139, 140, 7, 8, 2, 2, 140, 160, 5,
	16, 9, 14, 141, 142, 12, 12, 2, 2, 142, 143, 7, 9, 2, 2, 143, 160, 5, 16,
	9, 13, 144, 145, 12, 11, 2, 2, 145, 146, 7, 6, 2, 2, 146, 160, 5, 16, 9,
	12, 147, 148, 12, 10, 2, 2, 148, 149, 7, 7, 2, 2, 149, 160, 5, 16, 9, 11,
	150, 151, 12, 9, 2, 2, 151, 152, 7, 4, 2, 2, 152, 160, 5, 16, 9, 10, 153,
	154, 12, 8, 2, 2, 154, 155, 7, 5, 2, 2, 155, 160, 5, 16, 9, 9, 156, 157,
	12, 7, 2, 2, 157, 158, 7, 30, 2, 2, 158, 160, 5, 16, 9, 8, 159, 109, 3,
	2, 2, 2, 159, 114, 3, 2, 2, 2, 159, 119, 3, 2, 2, 2, 159, 124, 3, 2, 2,
	2, 159, 129, 3, 2, 2, 2, 159, 132, 3, 2, 2, 2, 159, 135, 3, 2, 2, 2, 159,
	138, 3, 2, 2, 2, 159, 141, 3, 2, 2, 2, 159, 144, 3, 2, 2, 2, 159, 147,
	3, 2, 2, 2, 159, 150, 3, 2, 2, 2, 159, 153, 3, 2, 2, 2, 159, 156, 3, 2,
	2, 2, 160, 163, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2,
	162, 17, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 164, 165, 7, 48, 2, 2, 165,
	166, 7, 20, 2, 2, 166, 167, 5, 16, 9, 2, 167, 168, 7, 21, 2, 2, 168, 169,
	7, 16, 2, 2, 169, 170, 8, 10, 1, 2, 170, 178, 3, 2, 2, 2, 171, 172, 7,
	48, 2, 2, 172, 173, 7, 20, 2, 2, 173, 174, 5, 16, 9, 2, 174, 175, 5, 20,
	11, 2, 175, 176, 7, 16, 2, 2, 176, 178, 3, 2, 2, 2, 177, 164, 3, 2, 2,
	2, 177, 171, 3, 2, 2, 2, 178, 19, 3, 2, 2, 2, 179, 180, 8, 11, 1, 2, 180,
	181, 7, 15, 2, 2, 181, 182, 5, 16, 9, 2, 182, 183, 7, 21, 2, 2, 183, 184,
	7, 16, 2, 2, 184, 185, 8, 11, 1, 2, 185, 195, 3, 2, 2, 2, 186, 187, 12,
	4, 2, 2, 187, 188, 7, 15, 2, 2, 188, 189, 5, 16, 9, 2, 189, 190, 7, 21,
	2, 2, 190, 191, 7, 16, 2, 2, 191, 192, 8, 11, 1, 2, 192, 194, 3, 2, 2,
	2, 193, 186, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 195,
	196, 3, 2, 2, 2, 196, 21, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 10, 25, 36,
	64, 107, 159, 161, 177, 195,
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
	"start", "instrucciones", "declaracion", "tipovariable", "igualacion",
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
	gramaticaParserRULE_declaracion     = 2
	gramaticaParserRULE_tipovariable    = 3
	gramaticaParserRULE_igualacion      = 4
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

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(gramaticaParserEOF, 0)
}

func (s *StartContext) AllInstrucciones() []IInstruccionesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem())
	var tst = make([]IInstruccionesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionesContext)
		}
	}

	return tst
}

func (s *StartContext) Instrucciones(i int) IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-18)&-(0x1f+1)) == 0 && ((1<<uint((_la-18)))&((1<<(gramaticaParserTK_par_apertura-18))|(1<<(gramaticaParserTK_menos-18))|(1<<(gramaticaParserTKR_pow-18))|(1<<(gramaticaParserTKR_powf-18))|(1<<(gramaticaParserTKR_let-18))|(1<<(gramaticaParserTKR_println-18))|(1<<(gramaticaParserTKR_true-18))|(1<<(gramaticaParserTKR_false-18)))) != 0) || (((_la-63)&-(0x1f+1)) == 0 && ((1<<uint((_la-63)))&((1<<(gramaticaParserTK_entero-63))|(1<<(gramaticaParserTK_decimal-63))|(1<<(gramaticaParserTK_id-63))|(1<<(gramaticaParserTK_cadena-63))|(1<<(gramaticaParserTK_caracter-63)))) != 0) {
		{
			p.SetState(20)
			p.Instrucciones()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(26)
		p.Match(gramaticaParserEOF)
	}

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	_expresion IExpresionContext
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

func (s *InstruccionesContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *InstruccionesContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *InstruccionesContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *InstruccionesContext) Impresion() IImpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpresionContext)
}

func (s *InstruccionesContext) Declaracion() IDeclaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *InstruccionesContext) Identificadores() IIdentificadoresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentificadoresContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentificadoresContext)
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

	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)

			var _x = p.expresion(0)

			localctx.(*InstruccionesContext)._expresion = _x
		}
		fmt.Println("mensaje en instrucciones: ", localctx.(*InstruccionesContext).Get_expresion().GetValorexpresion())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Impresion()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)
			p.Declaracion()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(33)
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

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *DeclaracionContext) TKR_let() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_let, 0)
}

func (s *DeclaracionContext) TKR_mut() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_mut, 0)
}

func (s *DeclaracionContext) TK_id() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_id, 0)
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

func (s *DeclaracionContext) Igualacion() IIgualacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIgualacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIgualacionContext)
}

func (s *DeclaracionContext) TK_pcoma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_pcoma, 0)
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
	p.EnterRule(localctx, 4, gramaticaParserRULE_declaracion)

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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(37)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(38)
			p.Match(gramaticaParserTK_id)
		}
		{
			p.SetState(39)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(40)
			p.Tipovariable()
		}
		{
			p.SetState(41)
			p.Igualacion()
		}
		{
			p.SetState(42)
			p.Match(gramaticaParserTK_pcoma)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(45)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(46)
			p.Match(gramaticaParserTK_id)
		}
		{
			p.SetState(47)
			p.Igualacion()
		}
		{
			p.SetState(48)
			p.Match(gramaticaParserTK_pcoma)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(51)
			p.Match(gramaticaParserTK_id)
		}
		{
			p.SetState(52)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(53)
			p.Tipovariable()
		}
		{
			p.SetState(54)
			p.Igualacion()
		}
		{
			p.SetState(55)
			p.Match(gramaticaParserTK_pcoma)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(57)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(58)
			p.Match(gramaticaParserTK_id)
		}
		{
			p.SetState(59)
			p.Igualacion()
		}
		{
			p.SetState(60)
			p.Match(gramaticaParserTK_pcoma)
		}

	}

	return localctx
}

// ITipovariableContext is an interface to support dynamic dispatch.
type ITipovariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTipovariableContext differentiates from other interfaces.
	IsTipovariableContext()
}

type TipovariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 6, gramaticaParserRULE_tipovariable)
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
	{
		p.SetState(64)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(gramaticaParserTKR_numericos_enteros-33))|(1<<(gramaticaParserTKR_numericos_flotantes-33))|(1<<(gramaticaParserTKR_bool-33))|(1<<(gramaticaParserTKR_char-33))|(1<<(gramaticaParserTKR_String-33))|(1<<(gramaticaParserTKR_usize-33)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIgualacionContext is an interface to support dynamic dispatch.
type IIgualacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIgualacionContext differentiates from other interfaces.
	IsIgualacionContext()
}

type IgualacionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIgualacionContext() *IgualacionContext {
	var p = new(IgualacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_igualacion
	return p
}

func (*IgualacionContext) IsIgualacionContext() {}

func NewIgualacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IgualacionContext {
	var p = new(IgualacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_igualacion

	return p
}

func (s *IgualacionContext) GetParser() antlr.Parser { return s.parser }

func (s *IgualacionContext) TK_igual() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_igual, 0)
}

func (s *IgualacionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *IgualacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IgualacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IgualacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterIgualacion(s)
	}
}

func (s *IgualacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitIgualacion(s)
	}
}

func (p *gramaticaParser) Igualacion() (localctx IIgualacionContext) {
	this := p
	_ = this

	localctx = NewIgualacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gramaticaParserRULE_igualacion)

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
		p.SetState(66)
		p.Match(gramaticaParserTK_igual)
	}
	{
		p.SetState(67)
		p.expresion(0)
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

func (s *IdentificadoresContext) Igualacion() IIgualacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIgualacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIgualacionContext)
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
		p.SetState(69)
		p.Match(gramaticaParserTK_id)
	}
	{
		p.SetState(70)
		p.Igualacion()
	}
	{
		p.SetState(71)
		p.Match(gramaticaParserTK_pcoma)
	}

	return localctx
}

// IValoresContext is an interface to support dynamic dispatch.
type IValoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValoresContext differentiates from other interfaces.
	IsValoresContext()
}

type ValoresContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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
	{
		p.SetState(73)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(gramaticaParserTKR_true-47))|(1<<(gramaticaParserTKR_false-47))|(1<<(gramaticaParserTK_entero-47))|(1<<(gramaticaParserTK_decimal-47))|(1<<(gramaticaParserTK_id-47))|(1<<(gramaticaParserTK_cadena-47))|(1<<(gramaticaParserTK_caracter-47)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetValorexpresion returns the valorexpresion attribute.
	GetValorexpresion() string

	// SetValorexpresion sets the valorexpresion attribute.
	SetValorexpresion(string)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	valorexpresion string
	e1             IExpresionContext
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

func (s *ExpresionContext) GetE2() IExpresionContext { return s.e2 }

func (s *ExpresionContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *ExpresionContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *ExpresionContext) GetValorexpresion() string { return s.valorexpresion }

func (s *ExpresionContext) SetValorexpresion(v string) { s.valorexpresion = v }

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

func (s *ExpresionContext) TK_suma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_suma, 0)
}

func (s *ExpresionContext) TK_por() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_por, 0)
}

func (s *ExpresionContext) TK_diagonal() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_diagonal, 0)
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(76)
			p.Match(gramaticaParserTK_menos)
		}
		{
			p.SetState(77)
			p.expresion(21)
		}

	case 2:
		{
			p.SetState(78)
			p.Match(gramaticaParserTKR_pow)
		}
		{
			p.SetState(79)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(80)
			p.expresion(0)
		}
		{
			p.SetState(81)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(82)
			p.expresion(0)
		}
		{
			p.SetState(83)
			p.Match(gramaticaParserTK_par_cierre)
		}

	case 3:
		{
			p.SetState(85)
			p.Match(gramaticaParserTKR_powf)
		}
		{
			p.SetState(86)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(87)
			p.expresion(0)
		}
		{
			p.SetState(88)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(89)
			p.expresion(0)
		}
		{
			p.SetState(90)
			p.Match(gramaticaParserTK_par_cierre)
		}

	case 4:
		{
			p.SetState(92)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(93)
			p.expresion(0)
		}
		{
			p.SetState(94)
			p.Match(gramaticaParserTK_par_cierre)
		}

	case 5:
		{
			p.SetState(96)
			p.Valores()
		}
		{
			p.SetState(97)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(98)
			p.Match(gramaticaParserTKR_numericos_enteros)
		}

	case 6:
		{
			p.SetState(100)
			p.Valores()
		}
		{
			p.SetState(101)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(102)
			p.Match(gramaticaParserTKR_numericos_flotantes)
		}

	case 7:
		{
			p.SetState(104)
			p.Valores()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(108)

					var _m = p.Match(gramaticaParserTK_suma)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(109)

					var _x = p.expresion(21)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).valorexpresion = Interfaces.OperacionAritmetica((func() string {
					if localctx.(*ExpresionContext).GetE1() == nil {
						return ""
					} else {
						return p.GetTokenStream().GetTextFromTokens(localctx.(*ExpresionContext).GetE1().GetStart(), localctx.(*ExpresionContext).e1.GetStop())
					}
				}()), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), (func() string {
					if localctx.(*ExpresionContext).GetE2() == nil {
						return ""
					} else {
						return p.GetTokenStream().GetTextFromTokens(localctx.(*ExpresionContext).GetE2().GetStart(), localctx.(*ExpresionContext).e2.GetStop())
					}
				}()))

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(112)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(113)

					var _m = p.Match(gramaticaParserTK_menos)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(114)

					var _x = p.expresion(20)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).valorexpresion = Interfaces.OperacionAritmetica((func() string {
					if localctx.(*ExpresionContext).GetE1() == nil {
						return ""
					} else {
						return p.GetTokenStream().GetTextFromTokens(localctx.(*ExpresionContext).GetE1().GetStart(), localctx.(*ExpresionContext).e1.GetStop())
					}
				}()), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), (func() string {
					if localctx.(*ExpresionContext).GetE2() == nil {
						return ""
					} else {
						return p.GetTokenStream().GetTextFromTokens(localctx.(*ExpresionContext).GetE2().GetStart(), localctx.(*ExpresionContext).e2.GetStop())
					}
				}()))

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(117)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(118)

					var _m = p.Match(gramaticaParserTK_por)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(119)

					var _x = p.expresion(19)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).valorexpresion = Interfaces.OperacionAritmetica((func() string {
					if localctx.(*ExpresionContext).GetE1() == nil {
						return ""
					} else {
						return p.GetTokenStream().GetTextFromTokens(localctx.(*ExpresionContext).GetE1().GetStart(), localctx.(*ExpresionContext).e1.GetStop())
					}
				}()), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), (func() string {
					if localctx.(*ExpresionContext).GetE2() == nil {
						return ""
					} else {
						return p.GetTokenStream().GetTextFromTokens(localctx.(*ExpresionContext).GetE2().GetStart(), localctx.(*ExpresionContext).e2.GetStop())
					}
				}()))

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(123)

					var _m = p.Match(gramaticaParserTK_diagonal)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(124)

					var _x = p.expresion(18)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).valorexpresion = Interfaces.OperacionAritmetica((func() string {
					if localctx.(*ExpresionContext).GetE1() == nil {
						return ""
					} else {
						return p.GetTokenStream().GetTextFromTokens(localctx.(*ExpresionContext).GetE1().GetStart(), localctx.(*ExpresionContext).e1.GetStop())
					}
				}()), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), (func() string {
					if localctx.(*ExpresionContext).GetE2() == nil {
						return ""
					} else {
						return p.GetTokenStream().GetTextFromTokens(localctx.(*ExpresionContext).GetE2().GetStart(), localctx.(*ExpresionContext).e2.GetStop())
					}
				}()))

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(127)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(128)
					p.Match(gramaticaParserTK_porcentaje)
				}
				{
					p.SetState(129)
					p.expresion(15)
				}

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(130)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(131)
					p.Match(gramaticaParserTK_menor)
				}
				{
					p.SetState(132)
					p.expresion(14)
				}

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(133)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(134)
					p.Match(gramaticaParserTK_mayor)
				}
				{
					p.SetState(135)
					p.expresion(13)
				}

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(136)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(137)
					p.Match(gramaticaParserTK_mayor_igual)
				}
				{
					p.SetState(138)
					p.expresion(12)
				}

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(139)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(140)
					p.Match(gramaticaParserTK_menor_igual)
				}
				{
					p.SetState(141)
					p.expresion(11)
				}

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(143)
					p.Match(gramaticaParserTK_igualacion)
				}
				{
					p.SetState(144)
					p.expresion(10)
				}

			case 11:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(145)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(146)
					p.Match(gramaticaParserTK_diferente)
				}
				{
					p.SetState(147)
					p.expresion(9)
				}

			case 12:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(149)
					p.Match(gramaticaParserTK_or)
				}
				{
					p.SetState(150)
					p.expresion(8)
				}

			case 13:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(152)
					p.Match(gramaticaParserTK_and)
				}
				{
					p.SetState(153)
					p.expresion(7)
				}

			case 14:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(154)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(155)
					p.Match(gramaticaParserTK_sig_admiracion)
				}
				{
					p.SetState(156)
					p.expresion(6)
				}

			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
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

	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(163)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(164)
			p.expresion(0)
		}
		{
			p.SetState(165)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(166)
			p.Match(gramaticaParserTK_pcoma)
		}
		fmt.Println("Impresion")

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(170)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(171)
			p.expresion(0)
		}
		{
			p.SetState(172)
			p.impresioncomas(0)
		}
		{
			p.SetState(173)
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
		p.SetState(178)
		p.Match(gramaticaParserTK_coma)
	}
	{
		p.SetState(179)
		p.expresion(0)
	}
	{
		p.SetState(180)
		p.Match(gramaticaParserTK_par_cierre)
	}
	{
		p.SetState(181)
		p.Match(gramaticaParserTK_pcoma)
	}
	fmt.Println("Impresion especial")

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewImpresioncomasContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_impresioncomas)
			p.SetState(184)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(185)
				p.Match(gramaticaParserTK_coma)
			}
			{
				p.SetState(186)
				p.expresion(0)
			}
			{
				p.SetState(187)
				p.Match(gramaticaParserTK_par_cierre)
			}
			{
				p.SetState(188)
				p.Match(gramaticaParserTK_pcoma)
			}
			fmt.Println("Impresion especial")

		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
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
