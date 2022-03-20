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
import "proyecto1/Expresiones"
import "proyecto1/Enum"
import "proyecto1/Funcion"

//import "proyecto1/Simbolo"
import arrayList "github.com/colegno/arraylist"

//import "proyecto1/Operaciones"
//import "proyecto1/Expresion"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 78, 514,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 7, 2, 54, 10, 2, 12, 2,
	14, 2, 57, 11, 2, 3, 2, 3, 2, 7, 2, 61, 10, 2, 12, 2, 14, 2, 64, 11, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 87, 10, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 7, 4, 102, 10, 4, 12, 4, 14, 4, 105, 11, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 7, 6, 117, 10, 6, 12, 6, 14, 6,
	120, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 148, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 5, 10, 221, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 235, 10, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 5, 13, 261, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 273, 10, 14, 12, 14, 14, 14,
	276, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 334, 10, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 409, 10,
	15, 12, 15, 14, 15, 412, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16,
	429, 10, 16, 3, 17, 7, 17, 432, 10, 17, 12, 17, 14, 17, 435, 11, 17, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 5, 20, 472, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 22, 7, 22, 480, 10, 22, 12, 22, 14, 22, 483, 11, 22, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 495, 10,
	23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 503, 10, 24, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 2, 5, 6,
	26, 28, 27, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 2, 4, 3, 2, 25, 27, 3, 2, 23, 24, 2,
	549, 2, 55, 3, 2, 2, 2, 4, 86, 3, 2, 2, 2, 6, 88, 3, 2, 2, 2, 8, 106, 3,
	2, 2, 2, 10, 118, 3, 2, 2, 2, 12, 147, 3, 2, 2, 2, 14, 149, 3, 2, 2, 2,
	16, 155, 3, 2, 2, 2, 18, 220, 3, 2, 2, 2, 20, 234, 3, 2, 2, 2, 22, 236,
	3, 2, 2, 2, 24, 260, 3, 2, 2, 2, 26, 262, 3, 2, 2, 2, 28, 333, 3, 2, 2,
	2, 30, 428, 3, 2, 2, 2, 32, 433, 3, 2, 2, 2, 34, 438, 3, 2, 2, 2, 36, 442,
	3, 2, 2, 2, 38, 471, 3, 2, 2, 2, 40, 473, 3, 2, 2, 2, 42, 481, 3, 2, 2,
	2, 44, 494, 3, 2, 2, 2, 46, 502, 3, 2, 2, 2, 48, 504, 3, 2, 2, 2, 50, 509,
	3, 2, 2, 2, 52, 54, 5, 4, 3, 2, 53, 52, 3, 2, 2, 2, 54, 57, 3, 2, 2, 2,
	55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 58, 3, 2, 2, 2, 57, 55, 3,
	2, 2, 2, 58, 62, 5, 8, 5, 2, 59, 61, 5, 4, 3, 2, 60, 59, 3, 2, 2, 2, 61,
	64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2,
	2, 64, 62, 3, 2, 2, 2, 65, 66, 8, 2, 1, 2, 66, 3, 3, 2, 2, 2, 67, 68, 7,
	52, 2, 2, 68, 69, 7, 75, 2, 2, 69, 70, 7, 20, 2, 2, 70, 71, 7, 21, 2, 2,
	71, 72, 7, 10, 2, 2, 72, 73, 5, 10, 6, 2, 73, 74, 7, 11, 2, 2, 74, 75,
	8, 3, 1, 2, 75, 87, 3, 2, 2, 2, 76, 77, 7, 52, 2, 2, 77, 78, 7, 75, 2,
	2, 78, 79, 7, 20, 2, 2, 79, 80, 5, 6, 4, 2, 80, 81, 7, 21, 2, 2, 81, 82,
	7, 10, 2, 2, 82, 83, 5, 10, 6, 2, 83, 84, 7, 11, 2, 2, 84, 85, 8, 3, 1,
	2, 85, 87, 3, 2, 2, 2, 86, 67, 3, 2, 2, 2, 86, 76, 3, 2, 2, 2, 87, 5, 3,
	2, 2, 2, 88, 89, 8, 4, 1, 2, 89, 90, 7, 75, 2, 2, 90, 91, 7, 14, 2, 2,
	91, 92, 5, 20, 11, 2, 92, 93, 8, 4, 1, 2, 93, 103, 3, 2, 2, 2, 94, 95,
	12, 4, 2, 2, 95, 96, 7, 15, 2, 2, 96, 97, 7, 75, 2, 2, 97, 98, 7, 14, 2,
	2, 98, 99, 5, 20, 11, 2, 99, 100, 8, 4, 1, 2, 100, 102, 3, 2, 2, 2, 101,
	94, 3, 2, 2, 2, 102, 105, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3,
	2, 2, 2, 104, 7, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 106, 107, 7, 52, 2,
	2, 107, 108, 7, 66, 2, 2, 108, 109, 7, 20, 2, 2, 109, 110, 7, 21, 2, 2,
	110, 111, 7, 10, 2, 2, 111, 112, 5, 10, 6, 2, 112, 113, 7, 11, 2, 2, 113,
	114, 8, 5, 1, 2, 114, 9, 3, 2, 2, 2, 115, 117, 5, 12, 7, 2, 116, 115, 3,
	2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2,
	2, 119, 121, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 122, 8, 6, 1, 2, 122,
	11, 3, 2, 2, 2, 123, 124, 5, 28, 15, 2, 124, 125, 8, 7, 1, 2, 125, 148,
	3, 2, 2, 2, 126, 127, 5, 30, 16, 2, 127, 128, 8, 7, 1, 2, 128, 148, 3,
	2, 2, 2, 129, 130, 5, 18, 10, 2, 130, 131, 8, 7, 1, 2, 131, 148, 3, 2,
	2, 2, 132, 133, 5, 22, 12, 2, 133, 134, 8, 7, 1, 2, 134, 148, 3, 2, 2,
	2, 135, 136, 5, 36, 19, 2, 136, 137, 8, 7, 1, 2, 137, 148, 3, 2, 2, 2,
	138, 139, 5, 46, 24, 2, 139, 140, 8, 7, 1, 2, 140, 148, 3, 2, 2, 2, 141,
	142, 5, 16, 9, 2, 142, 143, 8, 7, 1, 2, 143, 148, 3, 2, 2, 2, 144, 145,
	5, 14, 8, 2, 145, 146, 8, 7, 1, 2, 146, 148, 3, 2, 2, 2, 147, 123, 3, 2,
	2, 2, 147, 126, 3, 2, 2, 2, 147, 129, 3, 2, 2, 2, 147, 132, 3, 2, 2, 2,
	147, 135, 3, 2, 2, 2, 147, 138, 3, 2, 2, 2, 147, 141, 3, 2, 2, 2, 147,
	144, 3, 2, 2, 2, 148, 13, 3, 2, 2, 2, 149, 150, 7, 75, 2, 2, 150, 151,
	7, 20, 2, 2, 151, 152, 7, 21, 2, 2, 152, 153, 7, 16, 2, 2, 153, 154, 8,
	8, 1, 2, 154, 15, 3, 2, 2, 2, 155, 156, 7, 71, 2, 2, 156, 157, 7, 16, 2,
	2, 157, 158, 8, 9, 1, 2, 158, 17, 3, 2, 2, 2, 159, 160, 7, 45, 2, 2, 160,
	161, 7, 46, 2, 2, 161, 162, 7, 75, 2, 2, 162, 163, 7, 14, 2, 2, 163, 164,
	5, 20, 11, 2, 164, 165, 7, 22, 2, 2, 165, 166, 5, 28, 15, 2, 166, 167,
	7, 16, 2, 2, 167, 168, 8, 10, 1, 2, 168, 221, 3, 2, 2, 2, 169, 170, 7,
	45, 2, 2, 170, 171, 7, 46, 2, 2, 171, 172, 7, 75, 2, 2, 172, 173, 7, 22,
	2, 2, 173, 174, 5, 28, 15, 2, 174, 175, 7, 16, 2, 2, 175, 176, 8, 10, 1,
	2, 176, 221, 3, 2, 2, 2, 177, 178, 7, 45, 2, 2, 178, 179, 7, 46, 2, 2,
	179, 180, 7, 75, 2, 2, 180, 181, 7, 14, 2, 2, 181, 182, 7, 12, 2, 2, 182,
	183, 5, 20, 11, 2, 183, 184, 7, 16, 2, 2, 184, 185, 5, 24, 13, 2, 185,
	186, 7, 13, 2, 2, 186, 187, 7, 22, 2, 2, 187, 188, 5, 28, 15, 2, 188, 189,
	7, 16, 2, 2, 189, 190, 8, 10, 1, 2, 190, 221, 3, 2, 2, 2, 191, 192, 7,
	45, 2, 2, 192, 193, 7, 75, 2, 2, 193, 194, 7, 14, 2, 2, 194, 195, 7, 12,
	2, 2, 195, 196, 5, 20, 11, 2, 196, 197, 7, 16, 2, 2, 197, 198, 5, 24, 13,
	2, 198, 199, 7, 13, 2, 2, 199, 200, 7, 22, 2, 2, 200, 201, 5, 28, 15, 2,
	201, 202, 7, 16, 2, 2, 202, 203, 8, 10, 1, 2, 203, 221, 3, 2, 2, 2, 204,
	205, 7, 45, 2, 2, 205, 206, 7, 75, 2, 2, 206, 207, 7, 14, 2, 2, 207, 208,
	5, 20, 11, 2, 208, 209, 7, 22, 2, 2, 209, 210, 5, 28, 15, 2, 210, 211,
	7, 16, 2, 2, 211, 212, 8, 10, 1, 2, 212, 221, 3, 2, 2, 2, 213, 214, 7,
	45, 2, 2, 214, 215, 7, 75, 2, 2, 215, 216, 7, 22, 2, 2, 216, 217, 5, 28,
	15, 2, 217, 218, 7, 16, 2, 2, 218, 219, 8, 10, 1, 2, 219, 221, 3, 2, 2,
	2, 220, 159, 3, 2, 2, 2, 220, 169, 3, 2, 2, 2, 220, 177, 3, 2, 2, 2, 220,
	191, 3, 2, 2, 2, 220, 204, 3, 2, 2, 2, 220, 213, 3, 2, 2, 2, 221, 19, 3,
	2, 2, 2, 222, 223, 7, 35, 2, 2, 223, 235, 8, 11, 1, 2, 224, 225, 7, 36,
	2, 2, 225, 235, 8, 11, 1, 2, 226, 227, 7, 43, 2, 2, 227, 235, 8, 11, 1,
	2, 228, 229, 7, 40, 2, 2, 229, 235, 8, 11, 1, 2, 230, 231, 7, 41, 2, 2,
	231, 235, 8, 11, 1, 2, 232, 233, 7, 44, 2, 2, 233, 235, 8, 11, 1, 2, 234,
	222, 3, 2, 2, 2, 234, 224, 3, 2, 2, 2, 234, 226, 3, 2, 2, 2, 234, 228,
	3, 2, 2, 2, 234, 230, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 235, 21, 3, 2,
	2, 2, 236, 237, 7, 75, 2, 2, 237, 238, 7, 22, 2, 2, 238, 239, 5, 28, 15,
	2, 239, 240, 7, 16, 2, 2, 240, 241, 8, 12, 1, 2, 241, 23, 3, 2, 2, 2, 242,
	243, 7, 73, 2, 2, 243, 261, 8, 13, 1, 2, 244, 245, 7, 74, 2, 2, 245, 261,
	8, 13, 1, 2, 246, 247, 7, 77, 2, 2, 247, 261, 8, 13, 1, 2, 248, 249, 7,
	78, 2, 2, 249, 261, 8, 13, 1, 2, 250, 251, 7, 29, 2, 2, 251, 252, 7, 42,
	2, 2, 252, 261, 8, 13, 1, 2, 253, 254, 7, 50, 2, 2, 254, 261, 8, 13, 1,
	2, 255, 256, 7, 51, 2, 2, 256, 261, 8, 13, 1, 2, 257, 258, 5, 26, 14, 2,
	258, 259, 8, 13, 1, 2, 259, 261, 3, 2, 2, 2, 260, 242, 3, 2, 2, 2, 260,
	244, 3, 2, 2, 2, 260, 246, 3, 2, 2, 2, 260, 248, 3, 2, 2, 2, 260, 250,
	3, 2, 2, 2, 260, 253, 3, 2, 2, 2, 260, 255, 3, 2, 2, 2, 260, 257, 3, 2,
	2, 2, 261, 25, 3, 2, 2, 2, 262, 263, 8, 14, 1, 2, 263, 264, 7, 75, 2, 2,
	264, 265, 8, 14, 1, 2, 265, 274, 3, 2, 2, 2, 266, 267, 12, 4, 2, 2, 267,
	268, 7, 12, 2, 2, 268, 269, 5, 28, 15, 2, 269, 270, 7, 13, 2, 2, 270, 271,
	8, 14, 1, 2, 271, 273, 3, 2, 2, 2, 272, 266, 3, 2, 2, 2, 273, 276, 3, 2,
	2, 2, 274, 272, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 27, 3, 2, 2, 2,
	276, 274, 3, 2, 2, 2, 277, 278, 8, 15, 1, 2, 278, 279, 7, 24, 2, 2, 279,
	280, 5, 28, 15, 26, 280, 281, 8, 15, 1, 2, 281, 334, 3, 2, 2, 2, 282, 283,
	7, 37, 2, 2, 283, 284, 7, 20, 2, 2, 284, 285, 5, 28, 15, 2, 285, 286, 7,
	15, 2, 2, 286, 287, 5, 28, 15, 2, 287, 288, 7, 21, 2, 2, 288, 289, 8, 15,
	1, 2, 289, 334, 3, 2, 2, 2, 290, 291, 7, 39, 2, 2, 291, 292, 7, 20, 2,
	2, 292, 293, 5, 28, 15, 2, 293, 294, 7, 15, 2, 2, 294, 295, 5, 28, 15,
	2, 295, 296, 7, 21, 2, 2, 296, 297, 8, 15, 1, 2, 297, 334, 3, 2, 2, 2,
	298, 299, 7, 30, 2, 2, 299, 300, 5, 28, 15, 13, 300, 301, 8, 15, 1, 2,
	301, 334, 3, 2, 2, 2, 302, 303, 7, 20, 2, 2, 303, 304, 5, 28, 15, 2, 304,
	305, 7, 21, 2, 2, 305, 306, 8, 15, 1, 2, 306, 334, 3, 2, 2, 2, 307, 308,
	5, 24, 13, 2, 308, 309, 7, 48, 2, 2, 309, 310, 7, 35, 2, 2, 310, 311, 8,
	15, 1, 2, 311, 334, 3, 2, 2, 2, 312, 313, 5, 24, 13, 2, 313, 314, 7, 48,
	2, 2, 314, 315, 7, 36, 2, 2, 315, 316, 8, 15, 1, 2, 316, 334, 3, 2, 2,
	2, 317, 318, 7, 12, 2, 2, 318, 319, 5, 28, 15, 2, 319, 320, 7, 16, 2, 2,
	320, 321, 5, 28, 15, 2, 321, 322, 7, 13, 2, 2, 322, 323, 8, 15, 1, 2, 323,
	334, 3, 2, 2, 2, 324, 325, 7, 12, 2, 2, 325, 326, 5, 28, 15, 2, 326, 327,
	5, 32, 17, 2, 327, 328, 7, 13, 2, 2, 328, 329, 8, 15, 1, 2, 329, 334, 3,
	2, 2, 2, 330, 331, 5, 24, 13, 2, 331, 332, 8, 15, 1, 2, 332, 334, 3, 2,
	2, 2, 333, 277, 3, 2, 2, 2, 333, 282, 3, 2, 2, 2, 333, 290, 3, 2, 2, 2,
	333, 298, 3, 2, 2, 2, 333, 302, 3, 2, 2, 2, 333, 307, 3, 2, 2, 2, 333,
	312, 3, 2, 2, 2, 333, 317, 3, 2, 2, 2, 333, 324, 3, 2, 2, 2, 333, 330,
	3, 2, 2, 2, 334, 410, 3, 2, 2, 2, 335, 336, 12, 25, 2, 2, 336, 337, 9,
	2, 2, 2, 337, 338, 5, 28, 15, 26, 338, 339, 8, 15, 1, 2, 339, 409, 3, 2,
	2, 2, 340, 341, 12, 24, 2, 2, 341, 342, 9, 3, 2, 2, 342, 343, 5, 28, 15,
	25, 343, 344, 8, 15, 1, 2, 344, 409, 3, 2, 2, 2, 345, 346, 12, 21, 2, 2,
	346, 347, 7, 17, 2, 2, 347, 348, 5, 28, 15, 22, 348, 349, 8, 15, 1, 2,
	349, 409, 3, 2, 2, 2, 350, 351, 12, 20, 2, 2, 351, 352, 7, 18, 2, 2, 352,
	353, 5, 28, 15, 21, 353, 354, 8, 15, 1, 2, 354, 409, 3, 2, 2, 2, 355, 356,
	12, 19, 2, 2, 356, 357, 7, 8, 2, 2, 357, 358, 5, 28, 15, 20, 358, 359,
	8, 15, 1, 2, 359, 409, 3, 2, 2, 2, 360, 361, 12, 18, 2, 2, 361, 362, 7,
	9, 2, 2, 362, 363, 5, 28, 15, 19, 363, 364, 8, 15, 1, 2, 364, 409, 3, 2,
	2, 2, 365, 366, 12, 17, 2, 2, 366, 367, 7, 6, 2, 2, 367, 368, 5, 28, 15,
	18, 368, 369, 8, 15, 1, 2, 369, 409, 3, 2, 2, 2, 370, 371, 12, 16, 2, 2,
	371, 372, 7, 7, 2, 2, 372, 373, 5, 28, 15, 17, 373, 374, 8, 15, 1, 2, 374,
	409, 3, 2, 2, 2, 375, 376, 12, 15, 2, 2, 376, 377, 7, 4, 2, 2, 377, 378,
	5, 28, 15, 16, 378, 379, 8, 15, 1, 2, 379, 409, 3, 2, 2, 2, 380, 381, 12,
	14, 2, 2, 381, 382, 7, 5, 2, 2, 382, 383, 5, 28, 15, 15, 383, 384, 8, 15,
	1, 2, 384, 409, 3, 2, 2, 2, 385, 386, 12, 9, 2, 2, 386, 387, 7, 19, 2,
	2, 387, 388, 7, 54, 2, 2, 388, 389, 7, 20, 2, 2, 389, 390, 7, 21, 2, 2,
	390, 409, 8, 15, 1, 2, 391, 392, 12, 8, 2, 2, 392, 393, 7, 19, 2, 2, 393,
	394, 7, 55, 2, 2, 394, 395, 7, 20, 2, 2, 395, 396, 7, 21, 2, 2, 396, 409,
	8, 15, 1, 2, 397, 398, 12, 7, 2, 2, 398, 399, 7, 19, 2, 2, 399, 400, 7,
	56, 2, 2, 400, 401, 7, 20, 2, 2, 401, 402, 7, 21, 2, 2, 402, 409, 8, 15,
	1, 2, 403, 404, 12, 6, 2, 2, 404, 405, 7, 19, 2, 2, 405, 406, 7, 57, 2,
	2, 406, 407, 7, 20, 2, 2, 407, 409, 7, 21, 2, 2, 408, 335, 3, 2, 2, 2,
	408, 340, 3, 2, 2, 2, 408, 345, 3, 2, 2, 2, 408, 350, 3, 2, 2, 2, 408,
	355, 3, 2, 2, 2, 408, 360, 3, 2, 2, 2, 408, 365, 3, 2, 2, 2, 408, 370,
	3, 2, 2, 2, 408, 375, 3, 2, 2, 2, 408, 380, 3, 2, 2, 2, 408, 385, 3, 2,
	2, 2, 408, 391, 3, 2, 2, 2, 408, 397, 3, 2, 2, 2, 408, 403, 3, 2, 2, 2,
	409, 412, 3, 2, 2, 2, 410, 408, 3, 2, 2, 2, 410, 411, 3, 2, 2, 2, 411,
	29, 3, 2, 2, 2, 412, 410, 3, 2, 2, 2, 413, 414, 7, 49, 2, 2, 414, 415,
	7, 20, 2, 2, 415, 416, 5, 28, 15, 2, 416, 417, 7, 21, 2, 2, 417, 418, 7,
	16, 2, 2, 418, 419, 8, 16, 1, 2, 419, 429, 3, 2, 2, 2, 420, 421, 7, 49,
	2, 2, 421, 422, 7, 20, 2, 2, 422, 423, 5, 28, 15, 2, 423, 424, 5, 32, 17,
	2, 424, 425, 7, 21, 2, 2, 425, 426, 7, 16, 2, 2, 426, 427, 8, 16, 1, 2,
	427, 429, 3, 2, 2, 2, 428, 413, 3, 2, 2, 2, 428, 420, 3, 2, 2, 2, 429,
	31, 3, 2, 2, 2, 430, 432, 5, 34, 18, 2, 431, 430, 3, 2, 2, 2, 432, 435,
	3, 2, 2, 2, 433, 431, 3, 2, 2, 2, 433, 434, 3, 2, 2, 2, 434, 436, 3, 2,
	2, 2, 435, 433, 3, 2, 2, 2, 436, 437, 8, 17, 1, 2, 437, 33, 3, 2, 2, 2,
	438, 439, 7, 15, 2, 2, 439, 440, 5, 28, 15, 2, 440, 441, 8, 18, 1, 2, 441,
	35, 3, 2, 2, 2, 442, 443, 5, 38, 20, 2, 443, 444, 8, 19, 1, 2, 444, 37,
	3, 2, 2, 2, 445, 446, 7, 67, 2, 2, 446, 447, 5, 28, 15, 2, 447, 448, 5,
	44, 23, 2, 448, 449, 8, 20, 1, 2, 449, 472, 3, 2, 2, 2, 450, 451, 7, 67,
	2, 2, 451, 452, 5, 28, 15, 2, 452, 453, 5, 44, 23, 2, 453, 454, 7, 69,
	2, 2, 454, 455, 5, 44, 23, 2, 455, 456, 8, 20, 1, 2, 456, 472, 3, 2, 2,
	2, 457, 458, 7, 67, 2, 2, 458, 459, 5, 28, 15, 2, 459, 460, 5, 44, 23,
	2, 460, 461, 5, 42, 22, 2, 461, 462, 7, 69, 2, 2, 462, 463, 5, 44, 23,
	2, 463, 464, 8, 20, 1, 2, 464, 472, 3, 2, 2, 2, 465, 466, 7, 67, 2, 2,
	466, 467, 5, 28, 15, 2, 467, 468, 5, 44, 23, 2, 468, 469, 5, 42, 22, 2,
	469, 470, 8, 20, 1, 2, 470, 472, 3, 2, 2, 2, 471, 445, 3, 2, 2, 2, 471,
	450, 3, 2, 2, 2, 471, 457, 3, 2, 2, 2, 471, 465, 3, 2, 2, 2, 472, 39, 3,
	2, 2, 2, 473, 474, 7, 68, 2, 2, 474, 475, 5, 28, 15, 2, 475, 476, 5, 44,
	23, 2, 476, 477, 8, 21, 1, 2, 477, 41, 3, 2, 2, 2, 478, 480, 5, 40, 21,
	2, 479, 478, 3, 2, 2, 2, 480, 483, 3, 2, 2, 2, 481, 479, 3, 2, 2, 2, 481,
	482, 3, 2, 2, 2, 482, 484, 3, 2, 2, 2, 483, 481, 3, 2, 2, 2, 484, 485,
	8, 22, 1, 2, 485, 43, 3, 2, 2, 2, 486, 487, 7, 10, 2, 2, 487, 488, 5, 10,
	6, 2, 488, 489, 7, 11, 2, 2, 489, 490, 8, 23, 1, 2, 490, 495, 3, 2, 2,
	2, 491, 492, 7, 10, 2, 2, 492, 493, 7, 11, 2, 2, 493, 495, 8, 23, 1, 2,
	494, 486, 3, 2, 2, 2, 494, 491, 3, 2, 2, 2, 495, 45, 3, 2, 2, 2, 496, 497,
	5, 48, 25, 2, 497, 498, 8, 24, 1, 2, 498, 503, 3, 2, 2, 2, 499, 500, 5,
	50, 26, 2, 500, 501, 8, 24, 1, 2, 501, 503, 3, 2, 2, 2, 502, 496, 3, 2,
	2, 2, 502, 499, 3, 2, 2, 2, 503, 47, 3, 2, 2, 2, 504, 505, 7, 70, 2, 2,
	505, 506, 5, 28, 15, 2, 506, 507, 5, 44, 23, 2, 507, 508, 8, 25, 1, 2,
	508, 49, 3, 2, 2, 2, 509, 510, 7, 72, 2, 2, 510, 511, 5, 44, 23, 2, 511,
	512, 8, 26, 1, 2, 512, 51, 3, 2, 2, 2, 21, 55, 62, 86, 103, 118, 147, 220,
	234, 260, 274, 333, 408, 410, 428, 433, 471, 481, 494, 502,
}
var literalNames = []string{
	"", "'->'", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'{'", "'}'",
	"'['", "']'", "':'", "','", "';'", "'<'", "'>'", "'.'", "'('", "')'", "'='",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'", "'!'", "'?'", "", "",
	"", "'i64'", "'f64'", "'pow'", "'vec'", "'powf'", "'bool'", "'char'", "'str'",
	"'String'", "'usize'", "'let'", "'mut'", "'struct'", "'as'", "'println!'",
	"'true'", "'false'", "'fn'", "'return'", "'abs'", "'sqrt'", "'to_string'",
	"'clone'", "'new'", "'len'", "'push'", "'remove'", "'contains'", "'insert'",
	"'capacity'", "'witch_capacity'", "'main'", "'if'", "'else if'", "'else'",
	"'while'", "'break'", "'loop'",
}
var symbolicNames = []string{
	"", "TK_flecha", "TK_or", "TK_and", "TK_igualacion", "TK_diferente", "TK_mayor_igual",
	"TK_menor_igual", "TK_corchete_apertura", "TK_corchete_cierre", "TK_llave_apertura",
	"TK_llave_cierre", "TK_dosPuntos", "TK_coma", "TK_pcoma", "TK_menor", "TK_mayor",
	"TK_punto", "TK_par_apertura", "TK_par_cierre", "TK_igual", "TK_suma",
	"TK_menos", "TK_por", "TK_diagonal", "TK_porcentaje", "TK_linea", "TK_amp",
	"TK_sig_admiracion", "TK_sig_interrogacion", "WHITESPACE", "TK_comentario_multi",
	"TK_comentario_lineal", "TKR_numericos_enteros", "TKR_numericos_flotantes",
	"TKR_pow", "TKR_vec", "TKR_powf", "TKR_bool", "TKR_char", "TKR_Str", "TKR_String",
	"TKR_usize", "TKR_let", "TKR_mut", "TKR_struct", "TKR_as", "TKR_println",
	"TKR_true", "TKR_false", "TKR_fn", "TKR_return", "TKR_abs", "TKR_sqrt",
	"TKR_to_string", "TKR_clone", "TKR_new", "TKR_len", "TKR_push", "TKR_remove",
	"TKR_contains", "TKR_insert", "TKR_capacity", "TKR_with_capacity", "TKR_main",
	"TKR_if", "TKR_elseif", "TKR_else", "TKR_while", "TKR_break", "TKR_loop",
	"TK_entero", "TK_decimal", "TK_id", "Letra", "TK_cadena", "TK_caracter",
}

var ruleNames = []string{
	"start", "funciones", "declararParametros", "main", "instrucciones", "instruccion",
	"asignacion", "control", "declaracion", "tipovariable", "identificadores",
	"valores", "arreglo", "expresion", "impresion", "impresionexpresion", "expcoma",
	"condicionales", "funcionif", "funcionelseif", "listaelseif", "bloque",
	"bucles", "fwhile", "floop",
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

//Temporalgramatica := "Esta en temporal";

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
	gramaticaParserTKR_Str                 = 40
	gramaticaParserTKR_String              = 41
	gramaticaParserTKR_usize               = 42
	gramaticaParserTKR_let                 = 43
	gramaticaParserTKR_mut                 = 44
	gramaticaParserTKR_struct              = 45
	gramaticaParserTKR_as                  = 46
	gramaticaParserTKR_println             = 47
	gramaticaParserTKR_true                = 48
	gramaticaParserTKR_false               = 49
	gramaticaParserTKR_fn                  = 50
	gramaticaParserTKR_return              = 51
	gramaticaParserTKR_abs                 = 52
	gramaticaParserTKR_sqrt                = 53
	gramaticaParserTKR_to_string           = 54
	gramaticaParserTKR_clone               = 55
	gramaticaParserTKR_new                 = 56
	gramaticaParserTKR_len                 = 57
	gramaticaParserTKR_push                = 58
	gramaticaParserTKR_remove              = 59
	gramaticaParserTKR_contains            = 60
	gramaticaParserTKR_insert              = 61
	gramaticaParserTKR_capacity            = 62
	gramaticaParserTKR_with_capacity       = 63
	gramaticaParserTKR_main                = 64
	gramaticaParserTKR_if                  = 65
	gramaticaParserTKR_elseif              = 66
	gramaticaParserTKR_else                = 67
	gramaticaParserTKR_while               = 68
	gramaticaParserTKR_break               = 69
	gramaticaParserTKR_loop                = 70
	gramaticaParserTK_entero               = 71
	gramaticaParserTK_decimal              = 72
	gramaticaParserTK_id                   = 73
	gramaticaParserLetra                   = 74
	gramaticaParserTK_cadena               = 75
	gramaticaParserTK_caracter             = 76
)

// gramaticaParser rules.
const (
	gramaticaParserRULE_start              = 0
	gramaticaParserRULE_funciones          = 1
	gramaticaParserRULE_declararParametros = 2
	gramaticaParserRULE_main               = 3
	gramaticaParserRULE_instrucciones      = 4
	gramaticaParserRULE_instruccion        = 5
	gramaticaParserRULE_asignacion         = 6
	gramaticaParserRULE_control            = 7
	gramaticaParserRULE_declaracion        = 8
	gramaticaParserRULE_tipovariable       = 9
	gramaticaParserRULE_identificadores    = 10
	gramaticaParserRULE_valores            = 11
	gramaticaParserRULE_arreglo            = 12
	gramaticaParserRULE_expresion          = 13
	gramaticaParserRULE_impresion          = 14
	gramaticaParserRULE_impresionexpresion = 15
	gramaticaParserRULE_expcoma            = 16
	gramaticaParserRULE_condicionales      = 17
	gramaticaParserRULE_funcionif          = 18
	gramaticaParserRULE_funcionelseif      = 19
	gramaticaParserRULE_listaelseif        = 20
	gramaticaParserRULE_bloque             = 21
	gramaticaParserRULE_bucles             = 22
	gramaticaParserRULE_fwhile             = 23
	gramaticaParserRULE_floop              = 24
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_funciones returns the _funciones rule contexts.
	Get_funciones() IFuncionesContext

	// GetListtt returns the listtt rule contexts.
	GetListtt() IMainContext

	// Set_funciones sets the _funciones rule contexts.
	Set_funciones(IFuncionesContext)

	// SetListtt sets the listtt rule contexts.
	SetListtt(IMainContext)

	// GetE returns the e rule context list.
	GetE() []IFuncionesContext

	// GetR returns the r rule context list.
	GetR() []IFuncionesContext

	// SetE sets the e rule context list.
	SetE([]IFuncionesContext)

	// SetR sets the r rule context list.
	SetR([]IFuncionesContext)

	// GetListainstrucciones returns the listainstrucciones attribute.
	GetListainstrucciones() *arrayList.List

	// SetListainstrucciones sets the listainstrucciones attribute.
	SetListainstrucciones(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	listainstrucciones *arrayList.List
	_funciones         IFuncionesContext
	e                  []IFuncionesContext
	listtt             IMainContext
	r                  []IFuncionesContext
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

func (s *StartContext) Get_funciones() IFuncionesContext { return s._funciones }

func (s *StartContext) GetListtt() IMainContext { return s.listtt }

func (s *StartContext) Set_funciones(v IFuncionesContext) { s._funciones = v }

func (s *StartContext) SetListtt(v IMainContext) { s.listtt = v }

func (s *StartContext) GetE() []IFuncionesContext { return s.e }

func (s *StartContext) GetR() []IFuncionesContext { return s.r }

func (s *StartContext) SetE(v []IFuncionesContext) { s.e = v }

func (s *StartContext) SetR(v []IFuncionesContext) { s.r = v }

func (s *StartContext) GetListainstrucciones() *arrayList.List { return s.listainstrucciones }

func (s *StartContext) SetListainstrucciones(v *arrayList.List) { s.listainstrucciones = v }

func (s *StartContext) Main() IMainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMainContext)
}

func (s *StartContext) AllFunciones() []IFuncionesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncionesContext)(nil)).Elem())
	var tst = make([]IFuncionesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncionesContext)
		}
	}

	return tst
}

func (s *StartContext) Funciones(i int) IFuncionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncionesContext)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(50)

				var _x = p.Funciones()

				localctx.(*StartContext)._funciones = _x
			}
			localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._funciones)

		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(56)

		var _x = p.Main()

		localctx.(*StartContext).listtt = _x
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserTKR_fn {
		{
			p.SetState(57)

			var _x = p.Funciones()

			localctx.(*StartContext)._funciones = _x
		}
		localctx.(*StartContext).r = append(localctx.(*StartContext).r, localctx.(*StartContext)._funciones)

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	localctx.(*StartContext).listainstrucciones = arrayList.New()
	listaInst := localctx.(*StartContext).GetE()
	for _, e := range listaInst {
		localctx.(*StartContext).listainstrucciones.Add(e.GetInst())
	}
	//fmt.Println("La lista es: --------",localctx.(*StartContext).listainstrucciones)
	listaInst2 := localctx.(*StartContext).GetR()
	for _, r := range listaInst2 {
		localctx.(*StartContext).listainstrucciones.Add(r.GetInst())
	}
	//fmt.Println("La lista es: --------",localctx.(*StartContext).listainstrucciones)

	//fmt.Println("\n\n\nLa lista es: ",localctx.(*StartContext).GetListtt().GetListainstrucciones())
	listatemp := localctx.(*StartContext).GetListtt().GetListainstrucciones()
	for _, s := range listatemp.ToArray() {
		localctx.(*StartContext).listainstrucciones.Add(s)
	}
	fmt.Println("Fin del for")
	fmt.Println("Lista final: ", localctx.(*StartContext).listainstrucciones)

	return localctx
}

// IFuncionesContext is an interface to support dynamic dispatch.
type IFuncionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdd returns the idd token.
	GetIdd() antlr.Token

	// SetIdd sets the idd token.
	SetIdd(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_declararParametros returns the _declararParametros rule contexts.
	Get_declararParametros() IDeclararParametrosContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_declararParametros sets the _declararParametros rule contexts.
	Set_declararParametros(IDeclararParametrosContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsFuncionesContext differentiates from other interfaces.
	IsFuncionesContext()
}

type FuncionesContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	inst                Interfaces.Instruccion
	idd                 antlr.Token
	_instrucciones      IInstruccionesContext
	_declararParametros IDeclararParametrosContext
}

func NewEmptyFuncionesContext() *FuncionesContext {
	var p = new(FuncionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_funciones
	return p
}

func (*FuncionesContext) IsFuncionesContext() {}

func NewFuncionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionesContext {
	var p = new(FuncionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_funciones

	return p
}

func (s *FuncionesContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionesContext) GetIdd() antlr.Token { return s.idd }

func (s *FuncionesContext) SetIdd(v antlr.Token) { s.idd = v }

func (s *FuncionesContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *FuncionesContext) Get_declararParametros() IDeclararParametrosContext {
	return s._declararParametros
}

func (s *FuncionesContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *FuncionesContext) Set_declararParametros(v IDeclararParametrosContext) {
	s._declararParametros = v
}

func (s *FuncionesContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *FuncionesContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *FuncionesContext) TKR_fn() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_fn, 0)
}

func (s *FuncionesContext) TK_par_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_apertura, 0)
}

func (s *FuncionesContext) TK_par_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_cierre, 0)
}

func (s *FuncionesContext) TK_corchete_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_corchete_apertura, 0)
}

func (s *FuncionesContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *FuncionesContext) TK_corchete_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_corchete_cierre, 0)
}

func (s *FuncionesContext) TK_id() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_id, 0)
}

func (s *FuncionesContext) DeclararParametros() IDeclararParametrosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclararParametrosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclararParametrosContext)
}

func (s *FuncionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFunciones(s)
	}
}

func (s *FuncionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFunciones(s)
	}
}

func (p *gramaticaParser) Funciones() (localctx IFuncionesContext) {
	this := p
	_ = this

	localctx = NewFuncionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gramaticaParserRULE_funciones)

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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Match(gramaticaParserTKR_fn)
		}
		{
			p.SetState(66)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*FuncionesContext).idd = _m
		}
		{
			p.SetState(67)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(68)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(69)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(70)

			var _x = p.Instrucciones()

			localctx.(*FuncionesContext)._instrucciones = _x
		}
		{
			p.SetState(71)
			p.Match(gramaticaParserTK_corchete_cierre)
		}
		localctx.(*FuncionesContext).inst = Funcion.NewFuncion((func() string {
			if localctx.(*FuncionesContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*FuncionesContext).GetIdd().GetText()
			}
		}()), Enum.STRING, nil, localctx.(*FuncionesContext).Get_instrucciones().GetLista())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.Match(gramaticaParserTKR_fn)
		}
		{
			p.SetState(75)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*FuncionesContext).idd = _m
		}
		{
			p.SetState(76)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(77)

			var _x = p.declararParametros(0)

			localctx.(*FuncionesContext)._declararParametros = _x
		}
		{
			p.SetState(78)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(79)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(80)

			var _x = p.Instrucciones()

			localctx.(*FuncionesContext)._instrucciones = _x
		}
		{
			p.SetState(81)
			p.Match(gramaticaParserTK_corchete_cierre)
		}
		localctx.(*FuncionesContext).inst = Funcion.NewFuncion((func() string {
			if localctx.(*FuncionesContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*FuncionesContext).GetIdd().GetText()
			}
		}()), Enum.STRING, localctx.(*FuncionesContext).Get_declararParametros().GetLista(), localctx.(*FuncionesContext).Get_instrucciones().GetLista())

	}

	return localctx
}

// IDeclararParametrosContext is an interface to support dynamic dispatch.
type IDeclararParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 token.
	GetE1() antlr.Token

	// SetE1 sets the e1 token.
	SetE1(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IDeclararParametrosContext

	// GetE2 returns the e2 rule contexts.
	GetE2() ITipovariableContext

	// SetList sets the list rule contexts.
	SetList(IDeclararParametrosContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(ITipovariableContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsDeclararParametrosContext differentiates from other interfaces.
	IsDeclararParametrosContext()
}

type DeclararParametrosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lista  *arrayList.List
	list   IDeclararParametrosContext
	e1     antlr.Token
	e2     ITipovariableContext
}

func NewEmptyDeclararParametrosContext() *DeclararParametrosContext {
	var p = new(DeclararParametrosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_declararParametros
	return p
}

func (*DeclararParametrosContext) IsDeclararParametrosContext() {}

func NewDeclararParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclararParametrosContext {
	var p = new(DeclararParametrosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_declararParametros

	return p
}

func (s *DeclararParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclararParametrosContext) GetE1() antlr.Token { return s.e1 }

func (s *DeclararParametrosContext) SetE1(v antlr.Token) { s.e1 = v }

func (s *DeclararParametrosContext) GetList() IDeclararParametrosContext { return s.list }

func (s *DeclararParametrosContext) GetE2() ITipovariableContext { return s.e2 }

func (s *DeclararParametrosContext) SetList(v IDeclararParametrosContext) { s.list = v }

func (s *DeclararParametrosContext) SetE2(v ITipovariableContext) { s.e2 = v }

func (s *DeclararParametrosContext) GetLista() *arrayList.List { return s.lista }

func (s *DeclararParametrosContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *DeclararParametrosContext) TK_dosPuntos() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_dosPuntos, 0)
}

func (s *DeclararParametrosContext) TK_id() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_id, 0)
}

func (s *DeclararParametrosContext) Tipovariable() ITipovariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipovariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipovariableContext)
}

func (s *DeclararParametrosContext) TK_coma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_coma, 0)
}

func (s *DeclararParametrosContext) DeclararParametros() IDeclararParametrosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclararParametrosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclararParametrosContext)
}

func (s *DeclararParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclararParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDeclararParametros(s)
	}
}

func (s *DeclararParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDeclararParametros(s)
	}
}

func (p *gramaticaParser) DeclararParametros() (localctx IDeclararParametrosContext) {
	return p.declararParametros(0)
}

func (p *gramaticaParser) declararParametros(_p int) (localctx IDeclararParametrosContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDeclararParametrosContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDeclararParametrosContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, gramaticaParserRULE_declararParametros, _p)

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
		p.SetState(87)

		var _m = p.Match(gramaticaParserTK_id)

		localctx.(*DeclararParametrosContext).e1 = _m
	}
	{
		p.SetState(88)
		p.Match(gramaticaParserTK_dosPuntos)
	}
	{
		p.SetState(89)

		var _x = p.Tipovariable()

		localctx.(*DeclararParametrosContext).e2 = _x
	}
	localctx.(*DeclararParametrosContext).lista = arrayList.New()
	localctx.(*DeclararParametrosContext).lista.Add(Funcion.NewParam((func() string {
		if localctx.(*DeclararParametrosContext).GetE1() == nil {
			return ""
		} else {
			return localctx.(*DeclararParametrosContext).GetE1().GetText()
		}
	}()), localctx.(*DeclararParametrosContext).GetE2().GetTipovar()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDeclararParametrosContext(p, _parentctx, _parentState)
			localctx.(*DeclararParametrosContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_declararParametros)
			p.SetState(92)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(93)
				p.Match(gramaticaParserTK_coma)
			}
			{
				p.SetState(94)

				var _m = p.Match(gramaticaParserTK_id)

				localctx.(*DeclararParametrosContext).e1 = _m
			}
			{
				p.SetState(95)
				p.Match(gramaticaParserTK_dosPuntos)
			}
			{
				p.SetState(96)

				var _x = p.Tipovariable()

				localctx.(*DeclararParametrosContext).e2 = _x
			}
			localctx.(*DeclararParametrosContext).GetList().GetLista().Add(Funcion.NewParam((func() string {
				if localctx.(*DeclararParametrosContext).GetE1() == nil {
					return ""
				} else {
					return localctx.(*DeclararParametrosContext).GetE1().GetText()
				}
			}()), localctx.(*DeclararParametrosContext).GetE2().GetTipovar()))
			localctx.(*DeclararParametrosContext).lista = localctx.(*DeclararParametrosContext).GetList().GetLista()

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetListainstrucciones returns the listainstrucciones attribute.
	GetListainstrucciones() *arrayList.List

	// SetListainstrucciones sets the listainstrucciones attribute.
	SetListainstrucciones(*arrayList.List)

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	listainstrucciones *arrayList.List
	_instrucciones     IInstruccionesContext
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *MainContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *MainContext) GetListainstrucciones() *arrayList.List { return s.listainstrucciones }

func (s *MainContext) SetListainstrucciones(v *arrayList.List) { s.listainstrucciones = v }

func (s *MainContext) TKR_fn() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_fn, 0)
}

func (s *MainContext) TKR_main() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_main, 0)
}

func (s *MainContext) TK_par_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_apertura, 0)
}

func (s *MainContext) TK_par_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_cierre, 0)
}

func (s *MainContext) TK_corchete_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_corchete_apertura, 0)
}

func (s *MainContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *MainContext) TK_corchete_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_corchete_cierre, 0)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitMain(s)
	}
}

func (p *gramaticaParser) Main() (localctx IMainContext) {
	this := p
	_ = this

	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gramaticaParserRULE_main)

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
		p.SetState(104)
		p.Match(gramaticaParserTKR_fn)
	}
	{
		p.SetState(105)
		p.Match(gramaticaParserTKR_main)
	}
	{
		p.SetState(106)
		p.Match(gramaticaParserTK_par_apertura)
	}
	{
		p.SetState(107)
		p.Match(gramaticaParserTK_par_cierre)
	}
	{
		p.SetState(108)
		p.Match(gramaticaParserTK_corchete_apertura)
	}
	{
		p.SetState(109)

		var _x = p.Instrucciones()

		localctx.(*MainContext)._instrucciones = _x
	}
	{
		p.SetState(110)
		p.Match(gramaticaParserTK_corchete_cierre)
	}
	localctx.(*MainContext).listainstrucciones = localctx.(*MainContext).Get_instrucciones().GetLista()

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

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	lista        *arrayList.List
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

func (s *InstruccionesContext) GetLista() *arrayList.List { return s.lista }

func (s *InstruccionesContext) SetLista(v *arrayList.List) { s.lista = v }

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
	p.EnterRule(localctx, 8, gramaticaParserRULE_instrucciones)

	localctx.(*InstruccionesContext).lista = arrayList.New()

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
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gramaticaParserTK_llave_apertura)|(1<<gramaticaParserTK_par_apertura)|(1<<gramaticaParserTK_menos)|(1<<gramaticaParserTK_amp)|(1<<gramaticaParserTK_sig_admiracion))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(gramaticaParserTKR_pow-35))|(1<<(gramaticaParserTKR_powf-35))|(1<<(gramaticaParserTKR_let-35))|(1<<(gramaticaParserTKR_println-35))|(1<<(gramaticaParserTKR_true-35))|(1<<(gramaticaParserTKR_false-35))|(1<<(gramaticaParserTKR_if-35)))) != 0) || (((_la-68)&-(0x1f+1)) == 0 && ((1<<uint((_la-68)))&((1<<(gramaticaParserTKR_while-68))|(1<<(gramaticaParserTKR_break-68))|(1<<(gramaticaParserTKR_loop-68))|(1<<(gramaticaParserTK_entero-68))|(1<<(gramaticaParserTK_decimal-68))|(1<<(gramaticaParserTK_id-68))|(1<<(gramaticaParserTK_cadena-68))|(1<<(gramaticaParserTK_caracter-68)))) != 0) {
		{
			p.SetState(113)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	listaInst := localctx.(*InstruccionesContext).GetE()
	for _, e := range listaInst {
		localctx.(*InstruccionesContext).lista.Add(e.GetInst())
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

	// Get_impresion returns the _impresion rule contexts.
	Get_impresion() IImpresionContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Get_identificadores returns the _identificadores rule contexts.
	Get_identificadores() IIdentificadoresContext

	// Get_condicionales returns the _condicionales rule contexts.
	Get_condicionales() ICondicionalesContext

	// Get_bucles returns the _bucles rule contexts.
	Get_bucles() IBuclesContext

	// Get_control returns the _control rule contexts.
	Get_control() IControlContext

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_impresion sets the _impresion rule contexts.
	Set_impresion(IImpresionContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_identificadores sets the _identificadores rule contexts.
	Set_identificadores(IIdentificadoresContext)

	// Set_condicionales sets the _condicionales rule contexts.
	Set_condicionales(ICondicionalesContext)

	// Set_bucles sets the _bucles rule contexts.
	Set_bucles(IBuclesContext)

	// Set_control sets the _control rule contexts.
	Set_control(IControlContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	inst             Interfaces.Instruccion
	_expresion       IExpresionContext
	_impresion       IImpresionContext
	_declaracion     IDeclaracionContext
	_identificadores IIdentificadoresContext
	_condicionales   ICondicionalesContext
	_bucles          IBuclesContext
	_control         IControlContext
	_asignacion      IAsignacionContext
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

func (s *InstruccionContext) Get_impresion() IImpresionContext { return s._impresion }

func (s *InstruccionContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *InstruccionContext) Get_identificadores() IIdentificadoresContext { return s._identificadores }

func (s *InstruccionContext) Get_condicionales() ICondicionalesContext { return s._condicionales }

func (s *InstruccionContext) Get_bucles() IBuclesContext { return s._bucles }

func (s *InstruccionContext) Get_control() IControlContext { return s._control }

func (s *InstruccionContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *InstruccionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *InstruccionContext) Set_impresion(v IImpresionContext) { s._impresion = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_identificadores(v IIdentificadoresContext) { s._identificadores = v }

func (s *InstruccionContext) Set_condicionales(v ICondicionalesContext) { s._condicionales = v }

func (s *InstruccionContext) Set_bucles(v IBuclesContext) { s._bucles = v }

func (s *InstruccionContext) Set_control(v IControlContext) { s._control = v }

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *InstruccionContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

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

func (s *InstruccionContext) Condicionales() ICondicionalesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondicionalesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondicionalesContext)
}

func (s *InstruccionContext) Bucles() IBuclesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuclesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuclesContext)
}

func (s *InstruccionContext) Control() IControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControlContext)
}

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
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
	p.EnterRule(localctx, 10, gramaticaParserRULE_instruccion)

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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)

			var _x = p.expresion(0)

			localctx.(*InstruccionContext)._expresion = _x
		}
		fmt.Println("mensaje en instrucciones: ", localctx.(*InstruccionContext).Get_expresion().GetExp())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)

			var _x = p.Impresion()

			localctx.(*InstruccionContext)._impresion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_impresion().GetInst()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_declaracion().GetInst()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)

			var _x = p.Identificadores()

			localctx.(*InstruccionContext)._identificadores = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_identificadores().GetInst()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(133)

			var _x = p.Condicionales()

			localctx.(*InstruccionContext)._condicionales = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_condicionales().GetInst()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(136)

			var _x = p.Bucles()

			localctx.(*InstruccionContext)._bucles = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_bucles().GetInst()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(139)

			var _x = p.Control()

			localctx.(*InstruccionContext)._control = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_control().GetInst()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(142)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_asignacion().GetInst()

	}

	return localctx
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_id returns the _TK_id token.
	Get_TK_id() antlr.Token

	// Set_TK_id sets the _TK_id token.
	Set_TK_id(antlr.Token)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	inst   Interfaces.Instruccion
	_TK_id antlr.Token
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacion
	return p
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) Get_TK_id() antlr.Token { return s._TK_id }

func (s *AsignacionContext) Set_TK_id(v antlr.Token) { s._TK_id = v }

func (s *AsignacionContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *AsignacionContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *AsignacionContext) TK_id() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_id, 0)
}

func (s *AsignacionContext) TK_par_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_apertura, 0)
}

func (s *AsignacionContext) TK_par_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_cierre, 0)
}

func (s *AsignacionContext) TK_pcoma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_pcoma, 0)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (p *gramaticaParser) Asignacion() (localctx IAsignacionContext) {
	this := p
	_ = this

	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gramaticaParserRULE_asignacion)

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
		p.SetState(147)

		var _m = p.Match(gramaticaParserTK_id)

		localctx.(*AsignacionContext)._TK_id = _m
	}
	{
		p.SetState(148)
		p.Match(gramaticaParserTK_par_apertura)
	}
	{
		p.SetState(149)
		p.Match(gramaticaParserTK_par_cierre)
	}
	{
		p.SetState(150)
		p.Match(gramaticaParserTK_pcoma)
	}
	localctx.(*AsignacionContext).inst = Funcion.NewLlamadafuncion((func() string {
		if localctx.(*AsignacionContext).Get_TK_id() == nil {
			return ""
		} else {
			return localctx.(*AsignacionContext).Get_TK_id().GetText()
		}
	}()), true, nil)

	return localctx
}

// IControlContext is an interface to support dynamic dispatch.
type IControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsControlContext differentiates from other interfaces.
	IsControlContext()
}

type ControlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	inst   Interfaces.Instruccion
}

func NewEmptyControlContext() *ControlContext {
	var p = new(ControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_control
	return p
}

func (*ControlContext) IsControlContext() {}

func NewControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlContext {
	var p = new(ControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_control

	return p
}

func (s *ControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *ControlContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *ControlContext) TKR_break() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_break, 0)
}

func (s *ControlContext) TK_pcoma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_pcoma, 0)
}

func (s *ControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterControl(s)
	}
}

func (s *ControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitControl(s)
	}
}

func (p *gramaticaParser) Control() (localctx IControlContext) {
	this := p
	_ = this

	localctx = NewControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gramaticaParserRULE_control)

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
		p.SetState(153)
		p.Match(gramaticaParserTKR_break)
	}
	{
		p.SetState(154)
		p.Match(gramaticaParserTK_pcoma)
	}
	localctx.(*ControlContext).inst = Instruccion.NewBreak()

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

	// GetCant returns the cant rule contexts.
	GetCant() IValoresContext

	// Set_tipovariable sets the _tipovariable rule contexts.
	Set_tipovariable(ITipovariableContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// SetCant sets the cant rule contexts.
	SetCant(IValoresContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	inst          Interfaces.Instruccion
	idd           antlr.Token
	_tipovariable ITipovariableContext
	_expresion    IExpresionContext
	cant          IValoresContext
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

func (s *DeclaracionContext) GetCant() IValoresContext { return s.cant }

func (s *DeclaracionContext) Set_tipovariable(v ITipovariableContext) { s._tipovariable = v }

func (s *DeclaracionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *DeclaracionContext) SetCant(v IValoresContext) { s.cant = v }

func (s *DeclaracionContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *DeclaracionContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

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

func (s *DeclaracionContext) AllTK_pcoma() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserTK_pcoma)
}

func (s *DeclaracionContext) TK_pcoma(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_pcoma, i)
}

func (s *DeclaracionContext) TK_id() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_id, 0)
}

func (s *DeclaracionContext) TK_llave_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_llave_apertura, 0)
}

func (s *DeclaracionContext) TK_llave_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_llave_cierre, 0)
}

func (s *DeclaracionContext) Valores() IValoresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValoresContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValoresContext)
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
	p.EnterRule(localctx, 16, gramaticaParserRULE_declaracion)

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

	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(158)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(159)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(160)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(161)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(162)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(163)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(164)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracion((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipovariable().GetTipovar(), localctx.(*DeclaracionContext).Get_expresion().GetExp(), false, false, true)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(168)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(169)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(170)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(171)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(172)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracion((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), Enum.SINTIPO, localctx.(*DeclaracionContext).Get_expresion().GetExp(), false, false, true)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(176)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(177)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(178)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(179)
			p.Match(gramaticaParserTK_llave_apertura)
		}
		{
			p.SetState(180)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(181)
			p.Match(gramaticaParserTK_pcoma)
		}
		{
			p.SetState(182)

			var _x = p.Valores()

			localctx.(*DeclaracionContext).cant = _x
		}
		{
			p.SetState(183)
			p.Match(gramaticaParserTK_llave_cierre)
		}
		{
			p.SetState(184)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(185)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(186)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracionArray((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipovariable().GetTipovar(), localctx.(*DeclaracionContext).GetCant().GetExp(), localctx.(*DeclaracionContext).Get_expresion().GetExp(), true)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(189)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(190)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(191)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(192)
			p.Match(gramaticaParserTK_llave_apertura)
		}
		{
			p.SetState(193)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(194)
			p.Match(gramaticaParserTK_pcoma)
		}
		{
			p.SetState(195)

			var _x = p.Valores()

			localctx.(*DeclaracionContext).cant = _x
		}
		{
			p.SetState(196)
			p.Match(gramaticaParserTK_llave_cierre)
		}
		{
			p.SetState(197)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(198)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(199)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracionArray((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipovariable().GetTipovar(), localctx.(*DeclaracionContext).GetCant().GetExp(), localctx.(*DeclaracionContext).Get_expresion().GetExp(), false)

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(202)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(203)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(204)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(205)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(206)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(207)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(208)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracion((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipovariable().GetTipovar(), localctx.(*DeclaracionContext).Get_expresion().GetExp(), false, false, false)

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(211)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(212)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(213)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(214)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(215)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracion((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), Enum.SINTIPO, localctx.(*DeclaracionContext).Get_expresion().GetExp(), false, false, false)

	}

	return localctx
}

// ITipovariableContext is an interface to support dynamic dispatch.
type ITipovariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipovar returns the tipovar attribute.
	GetTipovar() Enum.Tipoexpresion

	// SetTipovar sets the tipovar attribute.
	SetTipovar(Enum.Tipoexpresion)

	// IsTipovariableContext differentiates from other interfaces.
	IsTipovariableContext()
}

type TipovariableContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	tipovar Enum.Tipoexpresion
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

func (s *TipovariableContext) GetTipovar() Enum.Tipoexpresion { return s.tipovar }

func (s *TipovariableContext) SetTipovar(v Enum.Tipoexpresion) { s.tipovar = v }

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
	p.EnterRule(localctx, 18, gramaticaParserRULE_tipovariable)

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

	p.SetState(232)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTKR_numericos_enteros:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Match(gramaticaParserTKR_numericos_enteros)
		}
		localctx.(*TipovariableContext).tipovar = Enum.INTEGER

	case gramaticaParserTKR_numericos_flotantes:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.Match(gramaticaParserTKR_numericos_flotantes)
		}
		localctx.(*TipovariableContext).tipovar = Enum.FLOAT

	case gramaticaParserTKR_String:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(224)
			p.Match(gramaticaParserTKR_String)
		}
		localctx.(*TipovariableContext).tipovar = Enum.STRING

	case gramaticaParserTKR_bool:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(226)
			p.Match(gramaticaParserTKR_bool)
		}
		localctx.(*TipovariableContext).tipovar = Enum.BOOLEAN

	case gramaticaParserTKR_char:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(228)
			p.Match(gramaticaParserTKR_char)
		}
		localctx.(*TipovariableContext).tipovar = Enum.CHAR

	case gramaticaParserTKR_usize:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(230)
			p.Match(gramaticaParserTKR_usize)
		}
		localctx.(*TipovariableContext).tipovar = Enum.USIZE

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

	// GetE1 returns the e1 token.
	GetE1() antlr.Token

	// SetE1 sets the e1 token.
	SetE1(antlr.Token)

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsIdentificadoresContext differentiates from other interfaces.
	IsIdentificadoresContext()
}

type IdentificadoresContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	inst   Interfaces.Instruccion
	e1     antlr.Token
	e2     IExpresionContext
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

func (s *IdentificadoresContext) GetE1() antlr.Token { return s.e1 }

func (s *IdentificadoresContext) SetE1(v antlr.Token) { s.e1 = v }

func (s *IdentificadoresContext) GetE2() IExpresionContext { return s.e2 }

func (s *IdentificadoresContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *IdentificadoresContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *IdentificadoresContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *IdentificadoresContext) TK_igual() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_igual, 0)
}

func (s *IdentificadoresContext) TK_pcoma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_pcoma, 0)
}

func (s *IdentificadoresContext) TK_id() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_id, 0)
}

func (s *IdentificadoresContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
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
	p.EnterRule(localctx, 20, gramaticaParserRULE_identificadores)

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
		p.SetState(234)

		var _m = p.Match(gramaticaParserTK_id)

		localctx.(*IdentificadoresContext).e1 = _m
	}
	{
		p.SetState(235)
		p.Match(gramaticaParserTK_igual)
	}
	{
		p.SetState(236)

		var _x = p.expresion(0)

		localctx.(*IdentificadoresContext).e2 = _x
	}
	{
		p.SetState(237)
		p.Match(gramaticaParserTK_pcoma)
	}
	localctx.(*IdentificadoresContext).inst = Instruccion.NewAsignacion((func() string {
		if localctx.(*IdentificadoresContext).GetE1() == nil {
			return ""
		} else {
			return localctx.(*IdentificadoresContext).GetE1().GetText()
		}
	}()), localctx.(*IdentificadoresContext).GetE2().GetExp())

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

	// Get_TKR_Str returns the _TKR_Str token.
	Get_TKR_Str() antlr.Token

	// Set_TK_entero sets the _TK_entero token.
	Set_TK_entero(antlr.Token)

	// Set_TK_decimal sets the _TK_decimal token.
	Set_TK_decimal(antlr.Token)

	// Set_TK_cadena sets the _TK_cadena token.
	Set_TK_cadena(antlr.Token)

	// Set_TK_caracter sets the _TK_caracter token.
	Set_TK_caracter(antlr.Token)

	// Set_TKR_Str sets the _TKR_Str token.
	Set_TKR_Str(antlr.Token)

	// GetAr returns the ar rule contexts.
	GetAr() IArregloContext

	// SetAr sets the ar rule contexts.
	SetAr(IArregloContext)

	// GetExp returns the exp attribute.
	GetExp() Interfaces.Expresion

	// SetExp sets the exp attribute.
	SetExp(Interfaces.Expresion)

	// IsValoresContext differentiates from other interfaces.
	IsValoresContext()
}

type ValoresContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	exp          Interfaces.Expresion
	_TK_entero   antlr.Token
	_TK_decimal  antlr.Token
	_TK_cadena   antlr.Token
	_TK_caracter antlr.Token
	_TKR_Str     antlr.Token
	ar           IArregloContext
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

func (s *ValoresContext) Get_TKR_Str() antlr.Token { return s._TKR_Str }

func (s *ValoresContext) Set_TK_entero(v antlr.Token) { s._TK_entero = v }

func (s *ValoresContext) Set_TK_decimal(v antlr.Token) { s._TK_decimal = v }

func (s *ValoresContext) Set_TK_cadena(v antlr.Token) { s._TK_cadena = v }

func (s *ValoresContext) Set_TK_caracter(v antlr.Token) { s._TK_caracter = v }

func (s *ValoresContext) Set_TKR_Str(v antlr.Token) { s._TKR_Str = v }

func (s *ValoresContext) GetAr() IArregloContext { return s.ar }

func (s *ValoresContext) SetAr(v IArregloContext) { s.ar = v }

func (s *ValoresContext) GetExp() Interfaces.Expresion { return s.exp }

func (s *ValoresContext) SetExp(v Interfaces.Expresion) { s.exp = v }

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

func (s *ValoresContext) TK_amp() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_amp, 0)
}

func (s *ValoresContext) TKR_Str() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_Str, 0)
}

func (s *ValoresContext) TKR_true() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_true, 0)
}

func (s *ValoresContext) TKR_false() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_false, 0)
}

func (s *ValoresContext) Arreglo() IArregloContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArregloContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArregloContext)
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
	p.EnterRule(localctx, 22, gramaticaParserRULE_valores)

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

	p.SetState(258)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTK_entero:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)

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
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(numero, Enum.INTEGER)

	case gramaticaParserTK_decimal:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)

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
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(decimal, Enum.FLOAT)

	case gramaticaParserTK_cadena:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(244)

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
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(str, Enum.STRING)

	case gramaticaParserTK_caracter:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(246)

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
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(str, Enum.CHAR)

	case gramaticaParserTK_amp:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(248)
			p.Match(gramaticaParserTK_amp)
		}
		{
			p.SetState(249)

			var _m = p.Match(gramaticaParserTKR_Str)

			localctx.(*ValoresContext)._TKR_Str = _m
		}
		str := (func() string {
			if localctx.(*ValoresContext).Get_TKR_Str() == nil {
				return ""
			} else {
				return localctx.(*ValoresContext).Get_TKR_Str().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ValoresContext).Get_TKR_Str() == nil {
				return ""
			} else {
				return localctx.(*ValoresContext).Get_TKR_Str().GetText()
			}
		}()))-1]
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(str, Enum.STR)

	case gramaticaParserTKR_true:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(251)
			p.Match(gramaticaParserTKR_true)
		}
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(true, Enum.BOOLEAN)

	case gramaticaParserTKR_false:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(253)
			p.Match(gramaticaParserTKR_false)
		}
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(false, Enum.BOOLEAN)

	case gramaticaParserTK_id:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(255)

			var _x = p.arreglo(0)

			localctx.(*ValoresContext).ar = _x
		}
		localctx.(*ValoresContext).exp = localctx.(*ValoresContext).GetAr().GetExp()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArregloContext is an interface to support dynamic dispatch.
type IArregloContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_id returns the _TK_id token.
	Get_TK_id() antlr.Token

	// Set_TK_id sets the _TK_id token.
	Set_TK_id(antlr.Token)

	// GetAr returns the ar rule contexts.
	GetAr() IArregloContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// SetAr sets the ar rule contexts.
	SetAr(IArregloContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetExp returns the exp attribute.
	GetExp() Interfaces.Expresion

	// SetExp sets the exp attribute.
	SetExp(Interfaces.Expresion)

	// IsArregloContext differentiates from other interfaces.
	IsArregloContext()
}

type ArregloContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	exp        Interfaces.Expresion
	ar         IArregloContext
	_TK_id     antlr.Token
	_expresion IExpresionContext
}

func NewEmptyArregloContext() *ArregloContext {
	var p = new(ArregloContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_arreglo
	return p
}

func (*ArregloContext) IsArregloContext() {}

func NewArregloContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArregloContext {
	var p = new(ArregloContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_arreglo

	return p
}

func (s *ArregloContext) GetParser() antlr.Parser { return s.parser }

func (s *ArregloContext) Get_TK_id() antlr.Token { return s._TK_id }

func (s *ArregloContext) Set_TK_id(v antlr.Token) { s._TK_id = v }

func (s *ArregloContext) GetAr() IArregloContext { return s.ar }

func (s *ArregloContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ArregloContext) SetAr(v IArregloContext) { s.ar = v }

func (s *ArregloContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ArregloContext) GetExp() Interfaces.Expresion { return s.exp }

func (s *ArregloContext) SetExp(v Interfaces.Expresion) { s.exp = v }

func (s *ArregloContext) TK_id() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_id, 0)
}

func (s *ArregloContext) TK_llave_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_llave_apertura, 0)
}

func (s *ArregloContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ArregloContext) TK_llave_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_llave_cierre, 0)
}

func (s *ArregloContext) Arreglo() IArregloContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArregloContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArregloContext)
}

func (s *ArregloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArregloContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArregloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterArreglo(s)
	}
}

func (s *ArregloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitArreglo(s)
	}
}

func (p *gramaticaParser) Arreglo() (localctx IArregloContext) {
	return p.arreglo(0)
}

func (p *gramaticaParser) arreglo(_p int) (localctx IArregloContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewArregloContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArregloContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, gramaticaParserRULE_arreglo, _p)

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
		p.SetState(261)

		var _m = p.Match(gramaticaParserTK_id)

		localctx.(*ArregloContext)._TK_id = _m
	}
	localctx.(*ArregloContext).exp = Expresiones.NewLlamarvariable((func() string {
		if localctx.(*ArregloContext).Get_TK_id() == nil {
			return ""
		} else {
			return localctx.(*ArregloContext).Get_TK_id().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewArregloContext(p, _parentctx, _parentState)
			localctx.(*ArregloContext).ar = _prevctx
			p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_arreglo)
			p.SetState(264)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(265)
				p.Match(gramaticaParserTK_llave_apertura)
			}
			{
				p.SetState(266)

				var _x = p.expresion(0)

				localctx.(*ArregloContext)._expresion = _x
			}
			{
				p.SetState(267)
				p.Match(gramaticaParserTK_llave_cierre)
			}
			localctx.(*ArregloContext).exp = Expresiones.NewAccesoArray(localctx.(*ArregloContext).GetAr().GetExp(), localctx.(*ArregloContext).Get_expresion().GetExp())

		}
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
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

	// GetVll returns the vll rule contexts.
	GetVll() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// GetVa returns the va rule contexts.
	GetVa() IExpresionContext

	// GetVal returns the val rule contexts.
	GetVal() IValoresContext

	// GetL1 returns the l1 rule contexts.
	GetL1() IImpresionexpresionContext

	// GetVall returns the vall rule contexts.
	GetVall() IValoresContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetVll sets the vll rule contexts.
	SetVll(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// SetVa sets the va rule contexts.
	SetVa(IExpresionContext)

	// SetVal sets the val rule contexts.
	SetVal(IValoresContext)

	// SetL1 sets the l1 rule contexts.
	SetL1(IImpresionexpresionContext)

	// SetVall sets the vall rule contexts.
	SetVall(IValoresContext)

	// GetExp returns the exp attribute.
	GetExp() Interfaces.Expresion

	// SetExp sets the exp attribute.
	SetExp(Interfaces.Expresion)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	exp    Interfaces.Expresion
	e1     IExpresionContext
	vll    IExpresionContext
	e2     IExpresionContext
	va     IExpresionContext
	val    IValoresContext
	l1     IImpresionexpresionContext
	vall   IValoresContext
	op     antlr.Token
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

func (s *ExpresionContext) GetVll() IExpresionContext { return s.vll }

func (s *ExpresionContext) GetE2() IExpresionContext { return s.e2 }

func (s *ExpresionContext) GetVa() IExpresionContext { return s.va }

func (s *ExpresionContext) GetVal() IValoresContext { return s.val }

func (s *ExpresionContext) GetL1() IImpresionexpresionContext { return s.l1 }

func (s *ExpresionContext) GetVall() IValoresContext { return s.vall }

func (s *ExpresionContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *ExpresionContext) SetVll(v IExpresionContext) { s.vll = v }

func (s *ExpresionContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *ExpresionContext) SetVa(v IExpresionContext) { s.va = v }

func (s *ExpresionContext) SetVal(v IValoresContext) { s.val = v }

func (s *ExpresionContext) SetL1(v IImpresionexpresionContext) { s.l1 = v }

func (s *ExpresionContext) SetVall(v IValoresContext) { s.vall = v }

func (s *ExpresionContext) GetExp() Interfaces.Expresion { return s.exp }

func (s *ExpresionContext) SetExp(v Interfaces.Expresion) { s.exp = v }

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

func (s *ExpresionContext) TK_sig_admiracion() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_sig_admiracion, 0)
}

func (s *ExpresionContext) TKR_as() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_as, 0)
}

func (s *ExpresionContext) TKR_numericos_enteros() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_numericos_enteros, 0)
}

func (s *ExpresionContext) Valores() IValoresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValoresContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValoresContext)
}

func (s *ExpresionContext) TKR_numericos_flotantes() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_numericos_flotantes, 0)
}

func (s *ExpresionContext) TK_llave_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_llave_apertura, 0)
}

func (s *ExpresionContext) TK_pcoma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_pcoma, 0)
}

func (s *ExpresionContext) TK_llave_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_llave_cierre, 0)
}

func (s *ExpresionContext) Impresionexpresion() IImpresionexpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpresionexpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpresionexpresionContext)
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

func (s *ExpresionContext) TK_suma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_suma, 0)
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

func (s *ExpresionContext) TK_punto() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_punto, 0)
}

func (s *ExpresionContext) TKR_abs() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_abs, 0)
}

func (s *ExpresionContext) TKR_sqrt() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_sqrt, 0)
}

func (s *ExpresionContext) TKR_to_string() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_to_string, 0)
}

func (s *ExpresionContext) TKR_clone() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_clone, 0)
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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, gramaticaParserRULE_expresion, _p)
	var _la int

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
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(276)
			p.Match(gramaticaParserTK_menos)
		}
		{
			p.SetState(277)

			var _x = p.expresion(24)

			localctx.(*ExpresionContext).e1 = _x
		}
		numero := -1
		numm := Expresion.NuevoPrimitivo(numero, Enum.INTEGER)
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), numm, Enum.MULTIPLICACION)

	case 2:
		{
			p.SetState(280)
			p.Match(gramaticaParserTKR_pow)
		}
		{
			p.SetState(281)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(282)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(283)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(284)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e2 = _x
		}
		{
			p.SetState(285)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.POW)

	case 3:
		{
			p.SetState(288)
			p.Match(gramaticaParserTKR_powf)
		}
		{
			p.SetState(289)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(290)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(291)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(292)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e2 = _x
		}
		{
			p.SetState(293)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.POWF)

	case 4:
		{
			p.SetState(296)
			p.Match(gramaticaParserTK_sig_admiracion)
		}
		{
			p.SetState(297)

			var _x = p.expresion(11)

			localctx.(*ExpresionContext).e1 = _x
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE1().GetExp(), Enum.NOT)

	case 5:
		{
			p.SetState(300)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(301)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).va = _x
		}
		{
			p.SetState(302)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = localctx.(*ExpresionContext).GetVa().GetExp()

	case 6:
		{
			p.SetState(305)

			var _x = p.Valores()

			localctx.(*ExpresionContext).val = _x
		}
		{
			p.SetState(306)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(307)
			p.Match(gramaticaParserTKR_numericos_enteros)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NewAsi64(localctx.(*ExpresionContext).GetVal().GetExp())

	case 7:
		{
			p.SetState(310)

			var _x = p.Valores()

			localctx.(*ExpresionContext).val = _x
		}
		{
			p.SetState(311)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(312)
			p.Match(gramaticaParserTKR_numericos_flotantes)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NewAsf64(localctx.(*ExpresionContext).GetVal().GetExp())

	case 8:
		{
			p.SetState(315)
			p.Match(gramaticaParserTK_llave_apertura)
		}
		{
			p.SetState(316)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(317)
			p.Match(gramaticaParserTK_pcoma)
		}
		{
			p.SetState(318)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e2 = _x
		}
		{
			p.SetState(319)
			p.Match(gramaticaParserTK_llave_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresion.NewArray(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), nil, Enum.MULTIPLE)

	case 9:
		{
			p.SetState(322)
			p.Match(gramaticaParserTK_llave_apertura)
		}
		{
			p.SetState(323)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(324)

			var _x = p.Impresionexpresion()

			localctx.(*ExpresionContext).l1 = _x
		}
		{
			p.SetState(325)
			p.Match(gramaticaParserTK_llave_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresion.NewArray(localctx.(*ExpresionContext).GetE1().GetExp(), nil, localctx.(*ExpresionContext).GetL1().GetLista(), Enum.NORMAL)

	case 10:
		{
			p.SetState(328)

			var _x = p.Valores()

			localctx.(*ExpresionContext).vall = _x
		}
		localctx.(*ExpresionContext).exp = localctx.(*ExpresionContext).GetVall().GetExp()
		fmt.Println(localctx.(*ExpresionContext).exp)

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(406)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(333)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
				}
				{
					p.SetState(334)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gramaticaParserTK_por)|(1<<gramaticaParserTK_diagonal)|(1<<gramaticaParserTK_porcentaje))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(335)

					var _x = p.expresion(24)

					localctx.(*ExpresionContext).e2 = _x
				}

				if (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()) == "*" {
					localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.MULTIPLICACION)
				} else if (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()) == "%" {
					localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.MODULO)
				} else {
					localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.DIVISION)
				}

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(338)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(339)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == gramaticaParserTK_suma || _la == gramaticaParserTK_menos) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(340)

					var _x = p.expresion(23)

					localctx.(*ExpresionContext).e2 = _x
				}

				if (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()) == "+" {
					localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.SUMA)
				} else {
					localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.RESTA)
				}

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(343)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(344)
					p.Match(gramaticaParserTK_menor)
				}
				{
					p.SetState(345)

					var _x = p.expresion(20)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.MENOR_QUE)

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(348)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(349)
					p.Match(gramaticaParserTK_mayor)
				}
				{
					p.SetState(350)

					var _x = p.expresion(19)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.MAYOR_QUE)

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(353)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(354)
					p.Match(gramaticaParserTK_mayor_igual)
				}
				{
					p.SetState(355)

					var _x = p.expresion(18)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.MAYOR_IGUAL)

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(358)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(359)
					p.Match(gramaticaParserTK_menor_igual)
				}
				{
					p.SetState(360)

					var _x = p.expresion(17)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.MENOR_IGUAL)

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(363)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(364)
					p.Match(gramaticaParserTK_igualacion)
				}
				{
					p.SetState(365)

					var _x = p.expresion(16)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.IGUALDAD)

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(368)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(369)
					p.Match(gramaticaParserTK_diferente)
				}
				{
					p.SetState(370)

					var _x = p.expresion(15)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.DESIGUAL)

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(373)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(374)
					p.Match(gramaticaParserTK_or)
				}
				{
					p.SetState(375)

					var _x = p.expresion(14)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.OR)

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(378)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(379)
					p.Match(gramaticaParserTK_and)
				}
				{
					p.SetState(380)

					var _x = p.expresion(13)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.AND)

			case 11:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).vll = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(383)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(384)
					p.Match(gramaticaParserTK_punto)
				}
				{
					p.SetState(385)
					p.Match(gramaticaParserTKR_abs)
				}
				{
					p.SetState(386)
					p.Match(gramaticaParserTK_par_apertura)
				}
				{
					p.SetState(387)
					p.Match(gramaticaParserTK_par_cierre)
				}
				localctx.(*ExpresionContext).exp = Expresiones.Newabs(localctx.(*ExpresionContext).GetVll().GetExp())

			case 12:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(389)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(390)
					p.Match(gramaticaParserTK_punto)
				}
				{
					p.SetState(391)
					p.Match(gramaticaParserTKR_sqrt)
				}
				{
					p.SetState(392)
					p.Match(gramaticaParserTK_par_apertura)
				}
				{
					p.SetState(393)
					p.Match(gramaticaParserTK_par_cierre)
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE1().GetExp(), Enum.MULTIPLICACION)

			case 13:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).vll = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(395)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(396)
					p.Match(gramaticaParserTK_punto)
				}
				{
					p.SetState(397)
					p.Match(gramaticaParserTKR_to_string)
				}
				{
					p.SetState(398)
					p.Match(gramaticaParserTK_par_apertura)
				}
				{
					p.SetState(399)
					p.Match(gramaticaParserTK_par_cierre)
				}
				localctx.(*ExpresionContext).exp = Expresiones.NewFto_string(localctx.(*ExpresionContext).GetVll().GetExp())

			case 14:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).vll = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(401)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(402)
					p.Match(gramaticaParserTK_punto)
				}
				{
					p.SetState(403)
					p.Match(gramaticaParserTKR_clone)
				}
				{
					p.SetState(404)
					p.Match(gramaticaParserTK_par_apertura)
				}
				{
					p.SetState(405)
					p.Match(gramaticaParserTK_par_cierre)
				}

			}

		}
		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IImpresionContext is an interface to support dynamic dispatch.
type IImpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// GetLi returns the li rule contexts.
	GetLi() IImpresionexpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// SetLi sets the li rule contexts.
	SetLi(IImpresionexpresionContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsImpresionContext differentiates from other interfaces.
	IsImpresionContext()
}

type ImpresionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	inst   Interfaces.Instruccion
	e1     IExpresionContext
	e2     IExpresionContext
	li     IImpresionexpresionContext
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

func (s *ImpresionContext) GetE1() IExpresionContext { return s.e1 }

func (s *ImpresionContext) GetE2() IExpresionContext { return s.e2 }

func (s *ImpresionContext) GetLi() IImpresionexpresionContext { return s.li }

func (s *ImpresionContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *ImpresionContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *ImpresionContext) SetLi(v IImpresionexpresionContext) { s.li = v }

func (s *ImpresionContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *ImpresionContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *ImpresionContext) TKR_println() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_println, 0)
}

func (s *ImpresionContext) TK_par_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_apertura, 0)
}

func (s *ImpresionContext) TK_par_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_cierre, 0)
}

func (s *ImpresionContext) TK_pcoma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_pcoma, 0)
}

func (s *ImpresionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ImpresionContext) Impresionexpresion() IImpresionexpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpresionexpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpresionexpresionContext)
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
	p.EnterRule(localctx, 28, gramaticaParserRULE_impresion)

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

	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(411)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(412)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(413)

			var _x = p.expresion(0)

			localctx.(*ImpresionContext).e1 = _x
		}
		{
			p.SetState(414)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(415)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*ImpresionContext).inst = Instruccion.NuevoPrint(localctx.(*ImpresionContext).GetE1().GetExp(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(418)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(419)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(420)

			var _x = p.expresion(0)

			localctx.(*ImpresionContext).e2 = _x
		}
		{
			p.SetState(421)

			var _x = p.Impresionexpresion()

			localctx.(*ImpresionContext).li = _x
		}
		{
			p.SetState(422)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(423)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*ImpresionContext).inst = Instruccion.NuevoPrint(localctx.(*ImpresionContext).GetE2().GetExp(), localctx.(*ImpresionContext).GetLi().GetLista())

	}

	return localctx
}

// IImpresionexpresionContext is an interface to support dynamic dispatch.
type IImpresionexpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expcoma returns the _expcoma rule contexts.
	Get_expcoma() IExpcomaContext

	// Set_expcoma sets the _expcoma rule contexts.
	Set_expcoma(IExpcomaContext)

	// GetList returns the list rule context list.
	GetList() []IExpcomaContext

	// SetList sets the list rule context list.
	SetList([]IExpcomaContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsImpresionexpresionContext differentiates from other interfaces.
	IsImpresionexpresionContext()
}

type ImpresionexpresionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	lista    *arrayList.List
	_expcoma IExpcomaContext
	list     []IExpcomaContext
}

func NewEmptyImpresionexpresionContext() *ImpresionexpresionContext {
	var p = new(ImpresionexpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_impresionexpresion
	return p
}

func (*ImpresionexpresionContext) IsImpresionexpresionContext() {}

func NewImpresionexpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpresionexpresionContext {
	var p = new(ImpresionexpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_impresionexpresion

	return p
}

func (s *ImpresionexpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpresionexpresionContext) Get_expcoma() IExpcomaContext { return s._expcoma }

func (s *ImpresionexpresionContext) Set_expcoma(v IExpcomaContext) { s._expcoma = v }

func (s *ImpresionexpresionContext) GetList() []IExpcomaContext { return s.list }

func (s *ImpresionexpresionContext) SetList(v []IExpcomaContext) { s.list = v }

func (s *ImpresionexpresionContext) GetLista() *arrayList.List { return s.lista }

func (s *ImpresionexpresionContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ImpresionexpresionContext) AllExpcoma() []IExpcomaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpcomaContext)(nil)).Elem())
	var tst = make([]IExpcomaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpcomaContext)
		}
	}

	return tst
}

func (s *ImpresionexpresionContext) Expcoma(i int) IExpcomaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpcomaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpcomaContext)
}

func (s *ImpresionexpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpresionexpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImpresionexpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterImpresionexpresion(s)
	}
}

func (s *ImpresionexpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitImpresionexpresion(s)
	}
}

func (p *gramaticaParser) Impresionexpresion() (localctx IImpresionexpresionContext) {
	this := p
	_ = this

	localctx = NewImpresionexpresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gramaticaParserRULE_impresionexpresion)
	localctx.(*ImpresionexpresionContext).lista = arrayList.New()
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
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserTK_coma {
		{
			p.SetState(428)

			var _x = p.Expcoma()

			localctx.(*ImpresionexpresionContext)._expcoma = _x
		}
		localctx.(*ImpresionexpresionContext).list = append(localctx.(*ImpresionexpresionContext).list, localctx.(*ImpresionexpresionContext)._expcoma)

		p.SetState(433)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*ImpresionexpresionContext).GetList()
	for _, e := range listInt {
		localctx.(*ImpresionexpresionContext).lista.Add(e.GetExp())
	}

	return localctx
}

// IExpcomaContext is an interface to support dynamic dispatch.
type IExpcomaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// GetExp returns the exp attribute.
	GetExp() Interfaces.Expresion

	// SetExp sets the exp attribute.
	SetExp(Interfaces.Expresion)

	// IsExpcomaContext differentiates from other interfaces.
	IsExpcomaContext()
}

type ExpcomaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	exp    Interfaces.Expresion
	e      IExpresionContext
}

func NewEmptyExpcomaContext() *ExpcomaContext {
	var p = new(ExpcomaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_expcoma
	return p
}

func (*ExpcomaContext) IsExpcomaContext() {}

func NewExpcomaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpcomaContext {
	var p = new(ExpcomaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_expcoma

	return p
}

func (s *ExpcomaContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpcomaContext) GetE() IExpresionContext { return s.e }

func (s *ExpcomaContext) SetE(v IExpresionContext) { s.e = v }

func (s *ExpcomaContext) GetExp() Interfaces.Expresion { return s.exp }

func (s *ExpcomaContext) SetExp(v Interfaces.Expresion) { s.exp = v }

func (s *ExpcomaContext) TK_coma() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_coma, 0)
}

func (s *ExpcomaContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpcomaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpcomaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpcomaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExpcoma(s)
	}
}

func (s *ExpcomaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExpcoma(s)
	}
}

func (p *gramaticaParser) Expcoma() (localctx IExpcomaContext) {
	this := p
	_ = this

	localctx = NewExpcomaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, gramaticaParserRULE_expcoma)

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
		p.SetState(436)
		p.Match(gramaticaParserTK_coma)
	}
	{
		p.SetState(437)

		var _x = p.expresion(0)

		localctx.(*ExpcomaContext).e = _x
	}
	localctx.(*ExpcomaContext).exp = localctx.(*ExpcomaContext).GetE().GetExp()

	return localctx
}

// ICondicionalesContext is an interface to support dynamic dispatch.
type ICondicionalesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_funcionif returns the _funcionif rule contexts.
	Get_funcionif() IFuncionifContext

	// Set_funcionif sets the _funcionif rule contexts.
	Set_funcionif(IFuncionifContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsCondicionalesContext differentiates from other interfaces.
	IsCondicionalesContext()
}

type CondicionalesContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	inst       Interfaces.Instruccion
	_funcionif IFuncionifContext
}

func NewEmptyCondicionalesContext() *CondicionalesContext {
	var p = new(CondicionalesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_condicionales
	return p
}

func (*CondicionalesContext) IsCondicionalesContext() {}

func NewCondicionalesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondicionalesContext {
	var p = new(CondicionalesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_condicionales

	return p
}

func (s *CondicionalesContext) GetParser() antlr.Parser { return s.parser }

func (s *CondicionalesContext) Get_funcionif() IFuncionifContext { return s._funcionif }

func (s *CondicionalesContext) Set_funcionif(v IFuncionifContext) { s._funcionif = v }

func (s *CondicionalesContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *CondicionalesContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *CondicionalesContext) Funcionif() IFuncionifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncionifContext)
}

func (s *CondicionalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondicionalesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondicionalesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCondicionales(s)
	}
}

func (s *CondicionalesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCondicionales(s)
	}
}

func (p *gramaticaParser) Condicionales() (localctx ICondicionalesContext) {
	this := p
	_ = this

	localctx = NewCondicionalesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, gramaticaParserRULE_condicionales)

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
		p.SetState(440)

		var _x = p.Funcionif()

		localctx.(*CondicionalesContext)._funcionif = _x
	}
	localctx.(*CondicionalesContext).inst = localctx.(*CondicionalesContext).Get_funcionif().GetInst()

	return localctx
}

// IFuncionifContext is an interface to support dynamic dispatch.
type IFuncionifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetEe returns the ee rule contexts.
	GetEe() IBloqueContext

	// GetE5 returns the e5 rule contexts.
	GetE5() IBloqueContext

	// GetB2 returns the b2 rule contexts.
	GetB2() IBloqueContext

	// GetB1 returns the b1 rule contexts.
	GetB1() IBloqueContext

	// Get_listaelseif returns the _listaelseif rule contexts.
	Get_listaelseif() IListaelseifContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetEe sets the ee rule contexts.
	SetEe(IBloqueContext)

	// SetE5 sets the e5 rule contexts.
	SetE5(IBloqueContext)

	// SetB2 sets the b2 rule contexts.
	SetB2(IBloqueContext)

	// SetB1 sets the b1 rule contexts.
	SetB1(IBloqueContext)

	// Set_listaelseif sets the _listaelseif rule contexts.
	Set_listaelseif(IListaelseifContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsFuncionifContext differentiates from other interfaces.
	IsFuncionifContext()
}

type FuncionifContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	inst         Interfaces.Instruccion
	e1           IExpresionContext
	ee           IBloqueContext
	e5           IBloqueContext
	b2           IBloqueContext
	b1           IBloqueContext
	_listaelseif IListaelseifContext
}

func NewEmptyFuncionifContext() *FuncionifContext {
	var p = new(FuncionifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcionif
	return p
}

func (*FuncionifContext) IsFuncionifContext() {}

func NewFuncionifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionifContext {
	var p = new(FuncionifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_funcionif

	return p
}

func (s *FuncionifContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionifContext) GetE1() IExpresionContext { return s.e1 }

func (s *FuncionifContext) GetEe() IBloqueContext { return s.ee }

func (s *FuncionifContext) GetE5() IBloqueContext { return s.e5 }

func (s *FuncionifContext) GetB2() IBloqueContext { return s.b2 }

func (s *FuncionifContext) GetB1() IBloqueContext { return s.b1 }

func (s *FuncionifContext) Get_listaelseif() IListaelseifContext { return s._listaelseif }

func (s *FuncionifContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *FuncionifContext) SetEe(v IBloqueContext) { s.ee = v }

func (s *FuncionifContext) SetE5(v IBloqueContext) { s.e5 = v }

func (s *FuncionifContext) SetB2(v IBloqueContext) { s.b2 = v }

func (s *FuncionifContext) SetB1(v IBloqueContext) { s.b1 = v }

func (s *FuncionifContext) Set_listaelseif(v IListaelseifContext) { s._listaelseif = v }

func (s *FuncionifContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *FuncionifContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *FuncionifContext) TKR_if() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_if, 0)
}

func (s *FuncionifContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *FuncionifContext) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *FuncionifContext) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *FuncionifContext) TKR_else() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_else, 0)
}

func (s *FuncionifContext) Listaelseif() IListaelseifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaelseifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaelseifContext)
}

func (s *FuncionifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFuncionif(s)
	}
}

func (s *FuncionifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFuncionif(s)
	}
}

func (p *gramaticaParser) Funcionif() (localctx IFuncionifContext) {
	this := p
	_ = this

	localctx = NewFuncionifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, gramaticaParserRULE_funcionif)

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

	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(443)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(444)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(445)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).ee = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetEe().GetLista(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(448)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(449)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(450)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).e5 = _x
		}
		{
			p.SetState(451)
			p.Match(gramaticaParserTKR_else)
		}
		{
			p.SetState(452)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b2 = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetE5().GetLista(), nil, localctx.(*FuncionifContext).GetB2().GetLista())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(455)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(456)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(457)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b1 = _x
		}
		{
			p.SetState(458)

			var _x = p.Listaelseif()

			localctx.(*FuncionifContext)._listaelseif = _x
		}
		{
			p.SetState(459)
			p.Match(gramaticaParserTKR_else)
		}
		{
			p.SetState(460)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b2 = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetB1().GetLista(), localctx.(*FuncionifContext).Get_listaelseif().GetLista(), localctx.(*FuncionifContext).GetB2().GetLista())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(463)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(464)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(465)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b1 = _x
		}
		{
			p.SetState(466)

			var _x = p.Listaelseif()

			localctx.(*FuncionifContext)._listaelseif = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetB1().GetLista(), localctx.(*FuncionifContext).Get_listaelseif().GetLista(), nil)

	}

	return localctx
}

// IFuncionelseifContext is an interface to support dynamic dispatch.
type IFuncionelseifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsFuncionelseifContext differentiates from other interfaces.
	IsFuncionelseifContext()
}

type FuncionelseifContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	inst    Interfaces.Instruccion
	e1      IExpresionContext
	_bloque IBloqueContext
}

func NewEmptyFuncionelseifContext() *FuncionelseifContext {
	var p = new(FuncionelseifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcionelseif
	return p
}

func (*FuncionelseifContext) IsFuncionelseifContext() {}

func NewFuncionelseifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionelseifContext {
	var p = new(FuncionelseifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_funcionelseif

	return p
}

func (s *FuncionelseifContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionelseifContext) GetE1() IExpresionContext { return s.e1 }

func (s *FuncionelseifContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *FuncionelseifContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *FuncionelseifContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *FuncionelseifContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *FuncionelseifContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *FuncionelseifContext) TKR_elseif() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_elseif, 0)
}

func (s *FuncionelseifContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *FuncionelseifContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *FuncionelseifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionelseifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionelseifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFuncionelseif(s)
	}
}

func (s *FuncionelseifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFuncionelseif(s)
	}
}

func (p *gramaticaParser) Funcionelseif() (localctx IFuncionelseifContext) {
	this := p
	_ = this

	localctx = NewFuncionelseifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, gramaticaParserRULE_funcionelseif)

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
		p.SetState(471)
		p.Match(gramaticaParserTKR_elseif)
	}
	{
		p.SetState(472)

		var _x = p.expresion(0)

		localctx.(*FuncionelseifContext).e1 = _x
	}
	{
		p.SetState(473)

		var _x = p.Bloque()

		localctx.(*FuncionelseifContext)._bloque = _x
	}
	localctx.(*FuncionelseifContext).inst = Instruccion.NewIf(localctx.(*FuncionelseifContext).GetE1().GetExp(), localctx.(*FuncionelseifContext).Get_bloque().GetLista(), nil, nil)

	return localctx
}

// IListaelseifContext is an interface to support dynamic dispatch.
type IListaelseifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_funcionelseif returns the _funcionelseif rule contexts.
	Get_funcionelseif() IFuncionelseifContext

	// Set_funcionelseif sets the _funcionelseif rule contexts.
	Set_funcionelseif(IFuncionelseifContext)

	// GetList returns the list rule context list.
	GetList() []IFuncionelseifContext

	// SetList sets the list rule context list.
	SetList([]IFuncionelseifContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsListaelseifContext differentiates from other interfaces.
	IsListaelseifContext()
}

type ListaelseifContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_funcionelseif IFuncionelseifContext
	list           []IFuncionelseifContext
}

func NewEmptyListaelseifContext() *ListaelseifContext {
	var p = new(ListaelseifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaelseif
	return p
}

func (*ListaelseifContext) IsListaelseifContext() {}

func NewListaelseifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaelseifContext {
	var p = new(ListaelseifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_listaelseif

	return p
}

func (s *ListaelseifContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaelseifContext) Get_funcionelseif() IFuncionelseifContext { return s._funcionelseif }

func (s *ListaelseifContext) Set_funcionelseif(v IFuncionelseifContext) { s._funcionelseif = v }

func (s *ListaelseifContext) GetList() []IFuncionelseifContext { return s.list }

func (s *ListaelseifContext) SetList(v []IFuncionelseifContext) { s.list = v }

func (s *ListaelseifContext) GetLista() *arrayList.List { return s.lista }

func (s *ListaelseifContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ListaelseifContext) AllFuncionelseif() []IFuncionelseifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncionelseifContext)(nil)).Elem())
	var tst = make([]IFuncionelseifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncionelseifContext)
		}
	}

	return tst
}

func (s *ListaelseifContext) Funcionelseif(i int) IFuncionelseifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionelseifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncionelseifContext)
}

func (s *ListaelseifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaelseifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaelseifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterListaelseif(s)
	}
}

func (s *ListaelseifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitListaelseif(s)
	}
}

func (p *gramaticaParser) Listaelseif() (localctx IListaelseifContext) {
	this := p
	_ = this

	localctx = NewListaelseifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, gramaticaParserRULE_listaelseif)
	localctx.(*ListaelseifContext).lista = arrayList.New()
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
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserTKR_elseif {
		{
			p.SetState(476)

			var _x = p.Funcionelseif()

			localctx.(*ListaelseifContext)._funcionelseif = _x
		}
		localctx.(*ListaelseifContext).list = append(localctx.(*ListaelseifContext).list, localctx.(*ListaelseifContext)._funcionelseif)

		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*ListaelseifContext).GetList()
	for _, e := range listInt {
		localctx.(*ListaelseifContext).lista.Add(e.GetInst())
	}

	return localctx
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_instrucciones IInstruccionesContext
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_bloque
	return p
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *BloqueContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *BloqueContext) GetLista() *arrayList.List { return s.lista }

func (s *BloqueContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *BloqueContext) TK_corchete_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_corchete_apertura, 0)
}

func (s *BloqueContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BloqueContext) TK_corchete_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_corchete_cierre, 0)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (p *gramaticaParser) Bloque() (localctx IBloqueContext) {
	this := p
	_ = this

	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, gramaticaParserRULE_bloque)

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

	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(484)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(485)

			var _x = p.Instrucciones()

			localctx.(*BloqueContext)._instrucciones = _x
		}
		{
			p.SetState(486)
			p.Match(gramaticaParserTK_corchete_cierre)
		}
		localctx.(*BloqueContext).lista = localctx.(*BloqueContext).Get_instrucciones().GetLista()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(489)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(490)
			p.Match(gramaticaParserTK_corchete_cierre)
		}
		localctx.(*BloqueContext).lista = arrayList.New()

	}

	return localctx
}

// IBuclesContext is an interface to support dynamic dispatch.
type IBuclesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_fwhile returns the _fwhile rule contexts.
	Get_fwhile() IFwhileContext

	// Get_floop returns the _floop rule contexts.
	Get_floop() IFloopContext

	// Set_fwhile sets the _fwhile rule contexts.
	Set_fwhile(IFwhileContext)

	// Set_floop sets the _floop rule contexts.
	Set_floop(IFloopContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsBuclesContext differentiates from other interfaces.
	IsBuclesContext()
}

type BuclesContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	inst    Interfaces.Instruccion
	_fwhile IFwhileContext
	_floop  IFloopContext
}

func NewEmptyBuclesContext() *BuclesContext {
	var p = new(BuclesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_bucles
	return p
}

func (*BuclesContext) IsBuclesContext() {}

func NewBuclesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuclesContext {
	var p = new(BuclesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_bucles

	return p
}

func (s *BuclesContext) GetParser() antlr.Parser { return s.parser }

func (s *BuclesContext) Get_fwhile() IFwhileContext { return s._fwhile }

func (s *BuclesContext) Get_floop() IFloopContext { return s._floop }

func (s *BuclesContext) Set_fwhile(v IFwhileContext) { s._fwhile = v }

func (s *BuclesContext) Set_floop(v IFloopContext) { s._floop = v }

func (s *BuclesContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *BuclesContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *BuclesContext) Fwhile() IFwhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFwhileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFwhileContext)
}

func (s *BuclesContext) Floop() IFloopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloopContext)
}

func (s *BuclesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuclesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuclesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBucles(s)
	}
}

func (s *BuclesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBucles(s)
	}
}

func (p *gramaticaParser) Bucles() (localctx IBuclesContext) {
	this := p
	_ = this

	localctx = NewBuclesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, gramaticaParserRULE_bucles)

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

	p.SetState(500)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTKR_while:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(494)

			var _x = p.Fwhile()

			localctx.(*BuclesContext)._fwhile = _x
		}
		localctx.(*BuclesContext).inst = localctx.(*BuclesContext).Get_fwhile().GetInst()

	case gramaticaParserTKR_loop:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(497)

			var _x = p.Floop()

			localctx.(*BuclesContext)._floop = _x
		}
		localctx.(*BuclesContext).inst = localctx.(*BuclesContext).Get_floop().GetInst()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFwhileContext is an interface to support dynamic dispatch.
type IFwhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetBl returns the bl rule contexts.
	GetBl() IBloqueContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetBl sets the bl rule contexts.
	SetBl(IBloqueContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsFwhileContext differentiates from other interfaces.
	IsFwhileContext()
}

type FwhileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	inst   Interfaces.Instruccion
	e1     IExpresionContext
	bl     IBloqueContext
}

func NewEmptyFwhileContext() *FwhileContext {
	var p = new(FwhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_fwhile
	return p
}

func (*FwhileContext) IsFwhileContext() {}

func NewFwhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FwhileContext {
	var p = new(FwhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_fwhile

	return p
}

func (s *FwhileContext) GetParser() antlr.Parser { return s.parser }

func (s *FwhileContext) GetE1() IExpresionContext { return s.e1 }

func (s *FwhileContext) GetBl() IBloqueContext { return s.bl }

func (s *FwhileContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *FwhileContext) SetBl(v IBloqueContext) { s.bl = v }

func (s *FwhileContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *FwhileContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *FwhileContext) TKR_while() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_while, 0)
}

func (s *FwhileContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *FwhileContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *FwhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FwhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FwhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFwhile(s)
	}
}

func (s *FwhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFwhile(s)
	}
}

func (p *gramaticaParser) Fwhile() (localctx IFwhileContext) {
	this := p
	_ = this

	localctx = NewFwhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, gramaticaParserRULE_fwhile)

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
		p.SetState(502)
		p.Match(gramaticaParserTKR_while)
	}
	{
		p.SetState(503)

		var _x = p.expresion(0)

		localctx.(*FwhileContext).e1 = _x
	}
	{
		p.SetState(504)

		var _x = p.Bloque()

		localctx.(*FwhileContext).bl = _x
	}
	localctx.(*FwhileContext).inst = Instruccion.NewWhile(localctx.(*FwhileContext).GetE1().GetExp(), localctx.(*FwhileContext).GetBl().GetLista())

	return localctx
}

// IFloopContext is an interface to support dynamic dispatch.
type IFloopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBb returns the bb rule contexts.
	GetBb() IBloqueContext

	// SetBb sets the bb rule contexts.
	SetBb(IBloqueContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsFloopContext differentiates from other interfaces.
	IsFloopContext()
}

type FloopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	inst   Interfaces.Instruccion
	bb     IBloqueContext
}

func NewEmptyFloopContext() *FloopContext {
	var p = new(FloopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gramaticaParserRULE_floop
	return p
}

func (*FloopContext) IsFloopContext() {}

func NewFloopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloopContext {
	var p = new(FloopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_floop

	return p
}

func (s *FloopContext) GetParser() antlr.Parser { return s.parser }

func (s *FloopContext) GetBb() IBloqueContext { return s.bb }

func (s *FloopContext) SetBb(v IBloqueContext) { s.bb = v }

func (s *FloopContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *FloopContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *FloopContext) TKR_loop() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_loop, 0)
}

func (s *FloopContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *FloopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFloop(s)
	}
}

func (s *FloopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFloop(s)
	}
}

func (p *gramaticaParser) Floop() (localctx IFloopContext) {
	this := p
	_ = this

	localctx = NewFloopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, gramaticaParserRULE_floop)

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
		p.SetState(507)
		p.Match(gramaticaParserTKR_loop)
	}
	{
		p.SetState(508)

		var _x = p.Bloque()

		localctx.(*FloopContext).bb = _x
	}
	localctx.(*FloopContext).inst = Instruccion.NewLoop(localctx.(*FloopContext).GetBb().GetLista())

	return localctx
}

func (p *gramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *DeclararParametrosContext = nil
		if localctx != nil {
			t = localctx.(*DeclararParametrosContext)
		}
		return p.DeclararParametros_Sempred(t, predIndex)

	case 12:
		var t *ArregloContext = nil
		if localctx != nil {
			t = localctx.(*ArregloContext)
		}
		return p.Arreglo_Sempred(t, predIndex)

	case 13:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *gramaticaParser) DeclararParametros_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *gramaticaParser) Arreglo_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *gramaticaParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
