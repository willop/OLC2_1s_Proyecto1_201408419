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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 78, 562,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 7, 2, 52, 10, 2, 12, 2, 14, 2, 55, 11,
	2, 3, 2, 3, 2, 7, 2, 59, 10, 2, 12, 2, 14, 2, 62, 11, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 85, 10, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 100,
	10, 4, 12, 4, 14, 4, 103, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 7, 6, 115, 10, 6, 12, 6, 14, 6, 118, 11, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 140, 10, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 5, 8, 159, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	5, 10, 251, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 265, 10, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 285, 10, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 297, 10, 13, 12,
	13, 14, 13, 300, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5,
	14, 382, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 7, 14, 457, 10, 14, 12, 14, 14, 14, 460, 11, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 5, 15, 477, 10, 15, 3, 16, 7, 16, 480, 10, 16, 12,
	16, 14, 16, 483, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 520, 10, 19, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 7, 21, 528, 10, 21, 12, 21, 14,
	21, 531, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 5, 22, 543, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 5, 23, 551, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 2, 5, 6, 24, 26, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 4, 3, 2,
	25, 27, 3, 2, 23, 24, 2, 603, 2, 53, 3, 2, 2, 2, 4, 84, 3, 2, 2, 2, 6,
	86, 3, 2, 2, 2, 8, 104, 3, 2, 2, 2, 10, 116, 3, 2, 2, 2, 12, 139, 3, 2,
	2, 2, 14, 158, 3, 2, 2, 2, 16, 160, 3, 2, 2, 2, 18, 250, 3, 2, 2, 2, 20,
	264, 3, 2, 2, 2, 22, 284, 3, 2, 2, 2, 24, 286, 3, 2, 2, 2, 26, 381, 3,
	2, 2, 2, 28, 476, 3, 2, 2, 2, 30, 481, 3, 2, 2, 2, 32, 486, 3, 2, 2, 2,
	34, 490, 3, 2, 2, 2, 36, 519, 3, 2, 2, 2, 38, 521, 3, 2, 2, 2, 40, 529,
	3, 2, 2, 2, 42, 542, 3, 2, 2, 2, 44, 550, 3, 2, 2, 2, 46, 552, 3, 2, 2,
	2, 48, 557, 3, 2, 2, 2, 50, 52, 5, 4, 3, 2, 51, 50, 3, 2, 2, 2, 52, 55,
	3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2,
	55, 53, 3, 2, 2, 2, 56, 60, 5, 8, 5, 2, 57, 59, 5, 4, 3, 2, 58, 57, 3,
	2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61,
	63, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 64, 8, 2, 1, 2, 64, 3, 3, 2, 2,
	2, 65, 66, 7, 52, 2, 2, 66, 67, 7, 75, 2, 2, 67, 68, 7, 20, 2, 2, 68, 69,
	7, 21, 2, 2, 69, 70, 7, 10, 2, 2, 70, 71, 5, 10, 6, 2, 71, 72, 7, 11, 2,
	2, 72, 73, 8, 3, 1, 2, 73, 85, 3, 2, 2, 2, 74, 75, 7, 52, 2, 2, 75, 76,
	7, 75, 2, 2, 76, 77, 7, 20, 2, 2, 77, 78, 5, 6, 4, 2, 78, 79, 7, 21, 2,
	2, 79, 80, 7, 10, 2, 2, 80, 81, 5, 10, 6, 2, 81, 82, 7, 11, 2, 2, 82, 83,
	8, 3, 1, 2, 83, 85, 3, 2, 2, 2, 84, 65, 3, 2, 2, 2, 84, 74, 3, 2, 2, 2,
	85, 5, 3, 2, 2, 2, 86, 87, 8, 4, 1, 2, 87, 88, 7, 75, 2, 2, 88, 89, 7,
	14, 2, 2, 89, 90, 5, 20, 11, 2, 90, 91, 8, 4, 1, 2, 91, 101, 3, 2, 2, 2,
	92, 93, 12, 4, 2, 2, 93, 94, 7, 15, 2, 2, 94, 95, 7, 75, 2, 2, 95, 96,
	7, 14, 2, 2, 96, 97, 5, 20, 11, 2, 97, 98, 8, 4, 1, 2, 98, 100, 3, 2, 2,
	2, 99, 92, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101,
	102, 3, 2, 2, 2, 102, 7, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104, 105, 7,
	52, 2, 2, 105, 106, 7, 66, 2, 2, 106, 107, 7, 20, 2, 2, 107, 108, 7, 21,
	2, 2, 108, 109, 7, 10, 2, 2, 109, 110, 5, 10, 6, 2, 110, 111, 7, 11, 2,
	2, 111, 112, 8, 5, 1, 2, 112, 9, 3, 2, 2, 2, 113, 115, 5, 12, 7, 2, 114,
	113, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117,
	3, 2, 2, 2, 117, 119, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 120, 8, 6,
	1, 2, 120, 11, 3, 2, 2, 2, 121, 122, 5, 28, 15, 2, 122, 123, 8, 7, 1, 2,
	123, 140, 3, 2, 2, 2, 124, 125, 5, 14, 8, 2, 125, 126, 8, 7, 1, 2, 126,
	140, 3, 2, 2, 2, 127, 128, 5, 18, 10, 2, 128, 129, 8, 7, 1, 2, 129, 140,
	3, 2, 2, 2, 130, 131, 5, 34, 18, 2, 131, 132, 8, 7, 1, 2, 132, 140, 3,
	2, 2, 2, 133, 134, 5, 44, 23, 2, 134, 135, 8, 7, 1, 2, 135, 140, 3, 2,
	2, 2, 136, 137, 5, 16, 9, 2, 137, 138, 8, 7, 1, 2, 138, 140, 3, 2, 2, 2,
	139, 121, 3, 2, 2, 2, 139, 124, 3, 2, 2, 2, 139, 127, 3, 2, 2, 2, 139,
	130, 3, 2, 2, 2, 139, 133, 3, 2, 2, 2, 139, 136, 3, 2, 2, 2, 140, 13, 3,
	2, 2, 2, 141, 142, 7, 75, 2, 2, 142, 143, 7, 20, 2, 2, 143, 144, 7, 21,
	2, 2, 144, 145, 7, 16, 2, 2, 145, 159, 8, 8, 1, 2, 146, 147, 7, 75, 2,
	2, 147, 148, 7, 22, 2, 2, 148, 149, 5, 26, 14, 2, 149, 150, 7, 16, 2, 2,
	150, 151, 8, 8, 1, 2, 151, 159, 3, 2, 2, 2, 152, 153, 5, 24, 13, 2, 153,
	154, 7, 22, 2, 2, 154, 155, 5, 26, 14, 2, 155, 156, 7, 16, 2, 2, 156, 157,
	8, 8, 1, 2, 157, 159, 3, 2, 2, 2, 158, 141, 3, 2, 2, 2, 158, 146, 3, 2,
	2, 2, 158, 152, 3, 2, 2, 2, 159, 15, 3, 2, 2, 2, 160, 161, 7, 71, 2, 2,
	161, 162, 7, 16, 2, 2, 162, 163, 8, 9, 1, 2, 163, 17, 3, 2, 2, 2, 164,
	165, 7, 45, 2, 2, 165, 166, 7, 46, 2, 2, 166, 167, 7, 75, 2, 2, 167, 168,
	7, 14, 2, 2, 168, 169, 5, 20, 11, 2, 169, 170, 7, 22, 2, 2, 170, 171, 5,
	26, 14, 2, 171, 172, 7, 16, 2, 2, 172, 173, 8, 10, 1, 2, 173, 251, 3, 2,
	2, 2, 174, 175, 7, 45, 2, 2, 175, 176, 7, 46, 2, 2, 176, 177, 7, 75, 2,
	2, 177, 178, 7, 22, 2, 2, 178, 179, 5, 26, 14, 2, 179, 180, 7, 16, 2, 2,
	180, 181, 8, 10, 1, 2, 181, 251, 3, 2, 2, 2, 182, 183, 7, 45, 2, 2, 183,
	184, 7, 46, 2, 2, 184, 185, 7, 75, 2, 2, 185, 186, 7, 14, 2, 2, 186, 187,
	7, 12, 2, 2, 187, 188, 5, 20, 11, 2, 188, 189, 7, 16, 2, 2, 189, 190, 5,
	22, 12, 2, 190, 191, 7, 13, 2, 2, 191, 192, 7, 22, 2, 2, 192, 193, 5, 26,
	14, 2, 193, 194, 7, 16, 2, 2, 194, 195, 8, 10, 1, 2, 195, 251, 3, 2, 2,
	2, 196, 197, 7, 45, 2, 2, 197, 198, 7, 75, 2, 2, 198, 199, 7, 14, 2, 2,
	199, 200, 7, 12, 2, 2, 200, 201, 5, 20, 11, 2, 201, 202, 7, 16, 2, 2, 202,
	203, 5, 22, 12, 2, 203, 204, 7, 13, 2, 2, 204, 205, 7, 22, 2, 2, 205, 206,
	5, 26, 14, 2, 206, 207, 7, 16, 2, 2, 207, 208, 8, 10, 1, 2, 208, 251, 3,
	2, 2, 2, 209, 210, 7, 45, 2, 2, 210, 211, 7, 46, 2, 2, 211, 212, 7, 75,
	2, 2, 212, 213, 7, 14, 2, 2, 213, 214, 7, 38, 2, 2, 214, 215, 7, 17, 2,
	2, 215, 216, 5, 20, 11, 2, 216, 217, 7, 18, 2, 2, 217, 218, 7, 22, 2, 2,
	218, 219, 5, 26, 14, 2, 219, 220, 7, 16, 2, 2, 220, 221, 8, 10, 1, 2, 221,
	251, 3, 2, 2, 2, 222, 223, 7, 45, 2, 2, 223, 224, 7, 75, 2, 2, 224, 225,
	7, 14, 2, 2, 225, 226, 7, 38, 2, 2, 226, 227, 7, 17, 2, 2, 227, 228, 5,
	20, 11, 2, 228, 229, 7, 18, 2, 2, 229, 230, 7, 22, 2, 2, 230, 231, 5, 26,
	14, 2, 231, 232, 7, 16, 2, 2, 232, 233, 8, 10, 1, 2, 233, 251, 3, 2, 2,
	2, 234, 235, 7, 45, 2, 2, 235, 236, 7, 75, 2, 2, 236, 237, 7, 14, 2, 2,
	237, 238, 5, 20, 11, 2, 238, 239, 7, 22, 2, 2, 239, 240, 5, 26, 14, 2,
	240, 241, 7, 16, 2, 2, 241, 242, 8, 10, 1, 2, 242, 251, 3, 2, 2, 2, 243,
	244, 7, 45, 2, 2, 244, 245, 7, 75, 2, 2, 245, 246, 7, 22, 2, 2, 246, 247,
	5, 26, 14, 2, 247, 248, 7, 16, 2, 2, 248, 249, 8, 10, 1, 2, 249, 251, 3,
	2, 2, 2, 250, 164, 3, 2, 2, 2, 250, 174, 3, 2, 2, 2, 250, 182, 3, 2, 2,
	2, 250, 196, 3, 2, 2, 2, 250, 209, 3, 2, 2, 2, 250, 222, 3, 2, 2, 2, 250,
	234, 3, 2, 2, 2, 250, 243, 3, 2, 2, 2, 251, 19, 3, 2, 2, 2, 252, 253, 7,
	35, 2, 2, 253, 265, 8, 11, 1, 2, 254, 255, 7, 36, 2, 2, 255, 265, 8, 11,
	1, 2, 256, 257, 7, 43, 2, 2, 257, 265, 8, 11, 1, 2, 258, 259, 7, 40, 2,
	2, 259, 265, 8, 11, 1, 2, 260, 261, 7, 41, 2, 2, 261, 265, 8, 11, 1, 2,
	262, 263, 7, 44, 2, 2, 263, 265, 8, 11, 1, 2, 264, 252, 3, 2, 2, 2, 264,
	254, 3, 2, 2, 2, 264, 256, 3, 2, 2, 2, 264, 258, 3, 2, 2, 2, 264, 260,
	3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 265, 21, 3, 2, 2, 2, 266, 267, 7, 73,
	2, 2, 267, 285, 8, 12, 1, 2, 268, 269, 7, 74, 2, 2, 269, 285, 8, 12, 1,
	2, 270, 271, 7, 77, 2, 2, 271, 285, 8, 12, 1, 2, 272, 273, 7, 78, 2, 2,
	273, 285, 8, 12, 1, 2, 274, 275, 7, 29, 2, 2, 275, 276, 7, 42, 2, 2, 276,
	285, 8, 12, 1, 2, 277, 278, 7, 50, 2, 2, 278, 285, 8, 12, 1, 2, 279, 280,
	7, 51, 2, 2, 280, 285, 8, 12, 1, 2, 281, 282, 5, 24, 13, 2, 282, 283, 8,
	12, 1, 2, 283, 285, 3, 2, 2, 2, 284, 266, 3, 2, 2, 2, 284, 268, 3, 2, 2,
	2, 284, 270, 3, 2, 2, 2, 284, 272, 3, 2, 2, 2, 284, 274, 3, 2, 2, 2, 284,
	277, 3, 2, 2, 2, 284, 279, 3, 2, 2, 2, 284, 281, 3, 2, 2, 2, 285, 23, 3,
	2, 2, 2, 286, 287, 8, 13, 1, 2, 287, 288, 7, 75, 2, 2, 288, 289, 8, 13,
	1, 2, 289, 298, 3, 2, 2, 2, 290, 291, 12, 4, 2, 2, 291, 292, 7, 12, 2,
	2, 292, 293, 5, 26, 14, 2, 293, 294, 7, 13, 2, 2, 294, 295, 8, 13, 1, 2,
	295, 297, 3, 2, 2, 2, 296, 290, 3, 2, 2, 2, 297, 300, 3, 2, 2, 2, 298,
	296, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 25, 3, 2, 2, 2, 300, 298, 3,
	2, 2, 2, 301, 302, 8, 14, 1, 2, 302, 303, 7, 24, 2, 2, 303, 304, 5, 26,
	14, 29, 304, 305, 8, 14, 1, 2, 305, 382, 3, 2, 2, 2, 306, 307, 7, 37, 2,
	2, 307, 308, 7, 20, 2, 2, 308, 309, 5, 26, 14, 2, 309, 310, 7, 15, 2, 2,
	310, 311, 5, 26, 14, 2, 311, 312, 7, 21, 2, 2, 312, 313, 8, 14, 1, 2, 313,
	382, 3, 2, 2, 2, 314, 315, 7, 39, 2, 2, 315, 316, 7, 20, 2, 2, 316, 317,
	5, 26, 14, 2, 317, 318, 7, 15, 2, 2, 318, 319, 5, 26, 14, 2, 319, 320,
	7, 21, 2, 2, 320, 321, 8, 14, 1, 2, 321, 382, 3, 2, 2, 2, 322, 323, 7,
	30, 2, 2, 323, 324, 5, 26, 14, 16, 324, 325, 8, 14, 1, 2, 325, 382, 3,
	2, 2, 2, 326, 327, 7, 20, 2, 2, 327, 328, 5, 26, 14, 2, 328, 329, 7, 21,
	2, 2, 329, 330, 8, 14, 1, 2, 330, 382, 3, 2, 2, 2, 331, 332, 5, 22, 12,
	2, 332, 333, 7, 48, 2, 2, 333, 334, 7, 35, 2, 2, 334, 335, 8, 14, 1, 2,
	335, 382, 3, 2, 2, 2, 336, 337, 5, 22, 12, 2, 337, 338, 7, 48, 2, 2, 338,
	339, 7, 36, 2, 2, 339, 340, 8, 14, 1, 2, 340, 382, 3, 2, 2, 2, 341, 342,
	7, 12, 2, 2, 342, 343, 5, 26, 14, 2, 343, 344, 7, 16, 2, 2, 344, 345, 5,
	26, 14, 2, 345, 346, 7, 13, 2, 2, 346, 347, 8, 14, 1, 2, 347, 382, 3, 2,
	2, 2, 348, 349, 7, 12, 2, 2, 349, 350, 5, 26, 14, 2, 350, 351, 5, 30, 16,
	2, 351, 352, 7, 13, 2, 2, 352, 353, 8, 14, 1, 2, 353, 382, 3, 2, 2, 2,
	354, 355, 7, 38, 2, 2, 355, 356, 7, 30, 2, 2, 356, 357, 7, 12, 2, 2, 357,
	358, 5, 26, 14, 2, 358, 359, 7, 16, 2, 2, 359, 360, 5, 26, 14, 2, 360,
	361, 7, 13, 2, 2, 361, 362, 8, 14, 1, 2, 362, 382, 3, 2, 2, 2, 363, 364,
	7, 38, 2, 2, 364, 365, 7, 30, 2, 2, 365, 366, 7, 12, 2, 2, 366, 367, 5,
	26, 14, 2, 367, 368, 5, 30, 16, 2, 368, 369, 7, 13, 2, 2, 369, 370, 8,
	14, 1, 2, 370, 382, 3, 2, 2, 2, 371, 372, 7, 38, 2, 2, 372, 373, 7, 14,
	2, 2, 373, 374, 7, 14, 2, 2, 374, 375, 7, 58, 2, 2, 375, 376, 7, 20, 2,
	2, 376, 377, 7, 21, 2, 2, 377, 382, 8, 14, 1, 2, 378, 379, 5, 22, 12, 2,
	379, 380, 8, 14, 1, 2, 380, 382, 3, 2, 2, 2, 381, 301, 3, 2, 2, 2, 381,
	306, 3, 2, 2, 2, 381, 314, 3, 2, 2, 2, 381, 322, 3, 2, 2, 2, 381, 326,
	3, 2, 2, 2, 381, 331, 3, 2, 2, 2, 381, 336, 3, 2, 2, 2, 381, 341, 3, 2,
	2, 2, 381, 348, 3, 2, 2, 2, 381, 354, 3, 2, 2, 2, 381, 363, 3, 2, 2, 2,
	381, 371, 3, 2, 2, 2, 381, 378, 3, 2, 2, 2, 382, 458, 3, 2, 2, 2, 383,
	384, 12, 28, 2, 2, 384, 385, 9, 2, 2, 2, 385, 386, 5, 26, 14, 29, 386,
	387, 8, 14, 1, 2, 387, 457, 3, 2, 2, 2, 388, 389, 12, 27, 2, 2, 389, 390,
	9, 3, 2, 2, 390, 391, 5, 26, 14, 28, 391, 392, 8, 14, 1, 2, 392, 457, 3,
	2, 2, 2, 393, 394, 12, 24, 2, 2, 394, 395, 7, 17, 2, 2, 395, 396, 5, 26,
	14, 25, 396, 397, 8, 14, 1, 2, 397, 457, 3, 2, 2, 2, 398, 399, 12, 23,
	2, 2, 399, 400, 7, 18, 2, 2, 400, 401, 5, 26, 14, 24, 401, 402, 8, 14,
	1, 2, 402, 457, 3, 2, 2, 2, 403, 404, 12, 22, 2, 2, 404, 405, 7, 8, 2,
	2, 405, 406, 5, 26, 14, 23, 406, 407, 8, 14, 1, 2, 407, 457, 3, 2, 2, 2,
	408, 409, 12, 21, 2, 2, 409, 410, 7, 9, 2, 2, 410, 411, 5, 26, 14, 22,
	411, 412, 8, 14, 1, 2, 412, 457, 3, 2, 2, 2, 413, 414, 12, 20, 2, 2, 414,
	415, 7, 6, 2, 2, 415, 416, 5, 26, 14, 21, 416, 417, 8, 14, 1, 2, 417, 457,
	3, 2, 2, 2, 418, 419, 12, 19, 2, 2, 419, 420, 7, 7, 2, 2, 420, 421, 5,
	26, 14, 20, 421, 422, 8, 14, 1, 2, 422, 457, 3, 2, 2, 2, 423, 424, 12,
	18, 2, 2, 424, 425, 7, 4, 2, 2, 425, 426, 5, 26, 14, 19, 426, 427, 8, 14,
	1, 2, 427, 457, 3, 2, 2, 2, 428, 429, 12, 17, 2, 2, 429, 430, 7, 5, 2,
	2, 430, 431, 5, 26, 14, 18, 431, 432, 8, 14, 1, 2, 432, 457, 3, 2, 2, 2,
	433, 434, 12, 12, 2, 2, 434, 435, 7, 19, 2, 2, 435, 436, 7, 54, 2, 2, 436,
	437, 7, 20, 2, 2, 437, 438, 7, 21, 2, 2, 438, 457, 8, 14, 1, 2, 439, 440,
	12, 11, 2, 2, 440, 441, 7, 19, 2, 2, 441, 442, 7, 55, 2, 2, 442, 443, 7,
	20, 2, 2, 443, 444, 7, 21, 2, 2, 444, 457, 8, 14, 1, 2, 445, 446, 12, 10,
	2, 2, 446, 447, 7, 19, 2, 2, 447, 448, 7, 56, 2, 2, 448, 449, 7, 20, 2,
	2, 449, 450, 7, 21, 2, 2, 450, 457, 8, 14, 1, 2, 451, 452, 12, 9, 2, 2,
	452, 453, 7, 19, 2, 2, 453, 454, 7, 57, 2, 2, 454, 455, 7, 20, 2, 2, 455,
	457, 7, 21, 2, 2, 456, 383, 3, 2, 2, 2, 456, 388, 3, 2, 2, 2, 456, 393,
	3, 2, 2, 2, 456, 398, 3, 2, 2, 2, 456, 403, 3, 2, 2, 2, 456, 408, 3, 2,
	2, 2, 456, 413, 3, 2, 2, 2, 456, 418, 3, 2, 2, 2, 456, 423, 3, 2, 2, 2,
	456, 428, 3, 2, 2, 2, 456, 433, 3, 2, 2, 2, 456, 439, 3, 2, 2, 2, 456,
	445, 3, 2, 2, 2, 456, 451, 3, 2, 2, 2, 457, 460, 3, 2, 2, 2, 458, 456,
	3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 27, 3, 2, 2, 2, 460, 458, 3, 2,
	2, 2, 461, 462, 7, 49, 2, 2, 462, 463, 7, 20, 2, 2, 463, 464, 5, 26, 14,
	2, 464, 465, 7, 21, 2, 2, 465, 466, 7, 16, 2, 2, 466, 467, 8, 15, 1, 2,
	467, 477, 3, 2, 2, 2, 468, 469, 7, 49, 2, 2, 469, 470, 7, 20, 2, 2, 470,
	471, 5, 26, 14, 2, 471, 472, 5, 30, 16, 2, 472, 473, 7, 21, 2, 2, 473,
	474, 7, 16, 2, 2, 474, 475, 8, 15, 1, 2, 475, 477, 3, 2, 2, 2, 476, 461,
	3, 2, 2, 2, 476, 468, 3, 2, 2, 2, 477, 29, 3, 2, 2, 2, 478, 480, 5, 32,
	17, 2, 479, 478, 3, 2, 2, 2, 480, 483, 3, 2, 2, 2, 481, 479, 3, 2, 2, 2,
	481, 482, 3, 2, 2, 2, 482, 484, 3, 2, 2, 2, 483, 481, 3, 2, 2, 2, 484,
	485, 8, 16, 1, 2, 485, 31, 3, 2, 2, 2, 486, 487, 7, 15, 2, 2, 487, 488,
	5, 26, 14, 2, 488, 489, 8, 17, 1, 2, 489, 33, 3, 2, 2, 2, 490, 491, 5,
	36, 19, 2, 491, 492, 8, 18, 1, 2, 492, 35, 3, 2, 2, 2, 493, 494, 7, 67,
	2, 2, 494, 495, 5, 26, 14, 2, 495, 496, 5, 42, 22, 2, 496, 497, 8, 19,
	1, 2, 497, 520, 3, 2, 2, 2, 498, 499, 7, 67, 2, 2, 499, 500, 5, 26, 14,
	2, 500, 501, 5, 42, 22, 2, 501, 502, 7, 69, 2, 2, 502, 503, 5, 42, 22,
	2, 503, 504, 8, 19, 1, 2, 504, 520, 3, 2, 2, 2, 505, 506, 7, 67, 2, 2,
	506, 507, 5, 26, 14, 2, 507, 508, 5, 42, 22, 2, 508, 509, 5, 40, 21, 2,
	509, 510, 7, 69, 2, 2, 510, 511, 5, 42, 22, 2, 511, 512, 8, 19, 1, 2, 512,
	520, 3, 2, 2, 2, 513, 514, 7, 67, 2, 2, 514, 515, 5, 26, 14, 2, 515, 516,
	5, 42, 22, 2, 516, 517, 5, 40, 21, 2, 517, 518, 8, 19, 1, 2, 518, 520,
	3, 2, 2, 2, 519, 493, 3, 2, 2, 2, 519, 498, 3, 2, 2, 2, 519, 505, 3, 2,
	2, 2, 519, 513, 3, 2, 2, 2, 520, 37, 3, 2, 2, 2, 521, 522, 7, 68, 2, 2,
	522, 523, 5, 26, 14, 2, 523, 524, 5, 42, 22, 2, 524, 525, 8, 20, 1, 2,
	525, 39, 3, 2, 2, 2, 526, 528, 5, 38, 20, 2, 527, 526, 3, 2, 2, 2, 528,
	531, 3, 2, 2, 2, 529, 527, 3, 2, 2, 2, 529, 530, 3, 2, 2, 2, 530, 532,
	3, 2, 2, 2, 531, 529, 3, 2, 2, 2, 532, 533, 8, 21, 1, 2, 533, 41, 3, 2,
	2, 2, 534, 535, 7, 10, 2, 2, 535, 536, 5, 10, 6, 2, 536, 537, 7, 11, 2,
	2, 537, 538, 8, 22, 1, 2, 538, 543, 3, 2, 2, 2, 539, 540, 7, 10, 2, 2,
	540, 541, 7, 11, 2, 2, 541, 543, 8, 22, 1, 2, 542, 534, 3, 2, 2, 2, 542,
	539, 3, 2, 2, 2, 543, 43, 3, 2, 2, 2, 544, 545, 5, 46, 24, 2, 545, 546,
	8, 23, 1, 2, 546, 551, 3, 2, 2, 2, 547, 548, 5, 48, 25, 2, 548, 549, 8,
	23, 1, 2, 549, 551, 3, 2, 2, 2, 550, 544, 3, 2, 2, 2, 550, 547, 3, 2, 2,
	2, 551, 45, 3, 2, 2, 2, 552, 553, 7, 70, 2, 2, 553, 554, 5, 26, 14, 2,
	554, 555, 5, 42, 22, 2, 555, 556, 8, 24, 1, 2, 556, 47, 3, 2, 2, 2, 557,
	558, 7, 72, 2, 2, 558, 559, 5, 42, 22, 2, 559, 560, 8, 25, 1, 2, 560, 49,
	3, 2, 2, 2, 22, 53, 60, 84, 101, 116, 139, 158, 250, 264, 284, 298, 381,
	456, 458, 476, 481, 519, 529, 542, 550,
}
var literalNames = []string{
	"", "'->'", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'{'", "'}'",
	"'['", "']'", "':'", "','", "';'", "'<'", "'>'", "'.'", "'('", "')'", "'='",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'", "'!'", "'?'", "", "",
	"", "'i64'", "'f64'", "'pow'", "'Vec'", "'powf'", "'bool'", "'char'", "'str'",
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
	"asignacion", "control", "declaracion", "tipovariable", "valores", "arreglo",
	"expresion", "impresion", "impresionexpresion", "expcoma", "condicionales",
	"funcionif", "funcionelseif", "listaelseif", "bloque", "bucles", "fwhile",
	"floop",
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
	gramaticaParserRULE_valores            = 10
	gramaticaParserRULE_arreglo            = 11
	gramaticaParserRULE_expresion          = 12
	gramaticaParserRULE_impresion          = 13
	gramaticaParserRULE_impresionexpresion = 14
	gramaticaParserRULE_expcoma            = 15
	gramaticaParserRULE_condicionales      = 16
	gramaticaParserRULE_funcionif          = 17
	gramaticaParserRULE_funcionelseif      = 18
	gramaticaParserRULE_listaelseif        = 19
	gramaticaParserRULE_bloque             = 20
	gramaticaParserRULE_bucles             = 21
	gramaticaParserRULE_fwhile             = 22
	gramaticaParserRULE_floop              = 23
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
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(48)

				var _x = p.Funciones()

				localctx.(*StartContext)._funciones = _x
			}
			localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._funciones)

		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(54)

		var _x = p.Main()

		localctx.(*StartContext).listtt = _x
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserTKR_fn {
		{
			p.SetState(55)

			var _x = p.Funciones()

			localctx.(*StartContext)._funciones = _x
		}
		localctx.(*StartContext).r = append(localctx.(*StartContext).r, localctx.(*StartContext)._funciones)

		p.SetState(60)
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

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Match(gramaticaParserTKR_fn)
		}
		{
			p.SetState(64)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*FuncionesContext).idd = _m
		}
		{
			p.SetState(65)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(66)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(67)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(68)

			var _x = p.Instrucciones()

			localctx.(*FuncionesContext)._instrucciones = _x
		}
		{
			p.SetState(69)
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
			p.SetState(72)
			p.Match(gramaticaParserTKR_fn)
		}
		{
			p.SetState(73)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*FuncionesContext).idd = _m
		}
		{
			p.SetState(74)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(75)

			var _x = p.declararParametros(0)

			localctx.(*FuncionesContext)._declararParametros = _x
		}
		{
			p.SetState(76)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(77)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(78)

			var _x = p.Instrucciones()

			localctx.(*FuncionesContext)._instrucciones = _x
		}
		{
			p.SetState(79)
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
		p.SetState(85)

		var _m = p.Match(gramaticaParserTK_id)

		localctx.(*DeclararParametrosContext).e1 = _m
	}
	{
		p.SetState(86)
		p.Match(gramaticaParserTK_dosPuntos)
	}
	{
		p.SetState(87)

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
	p.SetState(99)
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
			p.SetState(90)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(91)
				p.Match(gramaticaParserTK_coma)
			}
			{
				p.SetState(92)

				var _m = p.Match(gramaticaParserTK_id)

				localctx.(*DeclararParametrosContext).e1 = _m
			}
			{
				p.SetState(93)
				p.Match(gramaticaParserTK_dosPuntos)
			}
			{
				p.SetState(94)

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
		p.SetState(101)
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
		p.SetState(102)
		p.Match(gramaticaParserTKR_fn)
	}
	{
		p.SetState(103)
		p.Match(gramaticaParserTKR_main)
	}
	{
		p.SetState(104)
		p.Match(gramaticaParserTK_par_apertura)
	}
	{
		p.SetState(105)
		p.Match(gramaticaParserTK_par_cierre)
	}
	{
		p.SetState(106)
		p.Match(gramaticaParserTK_corchete_apertura)
	}
	{
		p.SetState(107)

		var _x = p.Instrucciones()

		localctx.(*MainContext)._instrucciones = _x
	}
	{
		p.SetState(108)
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
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(gramaticaParserTKR_let-43))|(1<<(gramaticaParserTKR_println-43))|(1<<(gramaticaParserTKR_if-43))|(1<<(gramaticaParserTKR_while-43))|(1<<(gramaticaParserTKR_break-43))|(1<<(gramaticaParserTKR_loop-43))|(1<<(gramaticaParserTK_id-43)))) != 0 {
		{
			p.SetState(111)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(116)
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

	// Get_impresion returns the _impresion rule contexts.
	Get_impresion() IImpresionContext

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Get_condicionales returns the _condicionales rule contexts.
	Get_condicionales() ICondicionalesContext

	// Get_bucles returns the _bucles rule contexts.
	Get_bucles() IBuclesContext

	// Get_control returns the _control rule contexts.
	Get_control() IControlContext

	// Set_impresion sets the _impresion rule contexts.
	Set_impresion(IImpresionContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_condicionales sets the _condicionales rule contexts.
	Set_condicionales(ICondicionalesContext)

	// Set_bucles sets the _bucles rule contexts.
	Set_bucles(IBuclesContext)

	// Set_control sets the _control rule contexts.
	Set_control(IControlContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	inst           Interfaces.Instruccion
	_impresion     IImpresionContext
	_asignacion    IAsignacionContext
	_declaracion   IDeclaracionContext
	_condicionales ICondicionalesContext
	_bucles        IBuclesContext
	_control       IControlContext
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

func (s *InstruccionContext) Get_impresion() IImpresionContext { return s._impresion }

func (s *InstruccionContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *InstruccionContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *InstruccionContext) Get_condicionales() ICondicionalesContext { return s._condicionales }

func (s *InstruccionContext) Get_bucles() IBuclesContext { return s._bucles }

func (s *InstruccionContext) Get_control() IControlContext { return s._control }

func (s *InstruccionContext) Set_impresion(v IImpresionContext) { s._impresion = v }

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_condicionales(v ICondicionalesContext) { s._condicionales = v }

func (s *InstruccionContext) Set_bucles(v IBuclesContext) { s._bucles = v }

func (s *InstruccionContext) Set_control(v IControlContext) { s._control = v }

func (s *InstruccionContext) GetInst() Interfaces.Instruccion { return s.inst }

func (s *InstruccionContext) SetInst(v Interfaces.Instruccion) { s.inst = v }

func (s *InstruccionContext) Impresion() IImpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpresionContext)
}

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *InstruccionContext) Declaracion() IDeclaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTKR_println:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)

			var _x = p.Impresion()

			localctx.(*InstruccionContext)._impresion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_impresion().GetInst()

	case gramaticaParserTK_id:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_asignacion().GetInst()

	case gramaticaParserTKR_let:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_declaracion().GetInst()

	case gramaticaParserTKR_if:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(128)

			var _x = p.Condicionales()

			localctx.(*InstruccionContext)._condicionales = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_condicionales().GetInst()

	case gramaticaParserTKR_while, gramaticaParserTKR_loop:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(131)

			var _x = p.Bucles()

			localctx.(*InstruccionContext)._bucles = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_bucles().GetInst()

	case gramaticaParserTKR_break:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(134)

			var _x = p.Control()

			localctx.(*InstruccionContext)._control = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_control().GetInst()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	// GetE1 returns the e1 token.
	GetE1() antlr.Token

	// Set_TK_id sets the _TK_id token.
	Set_TK_id(antlr.Token)

	// SetE1 sets the e1 token.
	SetE1(antlr.Token)

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// Get_arreglo returns the _arreglo rule contexts.
	Get_arreglo() IArregloContext

	// GetEe returns the ee rule contexts.
	GetEe() IExpresionContext

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// Set_arreglo sets the _arreglo rule contexts.
	Set_arreglo(IArregloContext)

	// SetEe sets the ee rule contexts.
	SetEe(IExpresionContext)

	// GetInst returns the inst attribute.
	GetInst() Interfaces.Instruccion

	// SetInst sets the inst attribute.
	SetInst(Interfaces.Instruccion)

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	inst     Interfaces.Instruccion
	_TK_id   antlr.Token
	e1       antlr.Token
	e2       IExpresionContext
	_arreglo IArregloContext
	ee       IExpresionContext
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

func (s *AsignacionContext) GetE1() antlr.Token { return s.e1 }

func (s *AsignacionContext) Set_TK_id(v antlr.Token) { s._TK_id = v }

func (s *AsignacionContext) SetE1(v antlr.Token) { s.e1 = v }

func (s *AsignacionContext) GetE2() IExpresionContext { return s.e2 }

func (s *AsignacionContext) Get_arreglo() IArregloContext { return s._arreglo }

func (s *AsignacionContext) GetEe() IExpresionContext { return s.ee }

func (s *AsignacionContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *AsignacionContext) Set_arreglo(v IArregloContext) { s._arreglo = v }

func (s *AsignacionContext) SetEe(v IExpresionContext) { s.ee = v }

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

func (s *AsignacionContext) TK_igual() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_igual, 0)
}

func (s *AsignacionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionContext) Arreglo() IArregloContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArregloContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArregloContext)
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

	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*AsignacionContext)._TK_id = _m
		}
		{
			p.SetState(140)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(141)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(142)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*AsignacionContext).inst = Funcion.NewLlamadafuncion((func() string {
			if localctx.(*AsignacionContext).Get_TK_id() == nil {
				return ""
			} else {
				return localctx.(*AsignacionContext).Get_TK_id().GetText()
			}
		}()), true, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*AsignacionContext).e1 = _m
		}
		{
			p.SetState(145)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(146)

			var _x = p.expresion(0)

			localctx.(*AsignacionContext).e2 = _x
		}
		{
			p.SetState(147)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*AsignacionContext).inst = Instruccion.NewAsignacion((func() string {
			if localctx.(*AsignacionContext).GetE1() == nil {
				return ""
			} else {
				return localctx.(*AsignacionContext).GetE1().GetText()
			}
		}()), localctx.(*AsignacionContext).GetE2().GetExp())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(150)

			var _x = p.arreglo(0)

			localctx.(*AsignacionContext)._arreglo = _x
		}
		{
			p.SetState(151)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(152)

			var _x = p.expresion(0)

			localctx.(*AsignacionContext).ee = _x
		}
		{
			p.SetState(153)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*AsignacionContext).inst = Instruccion.NewAsignacionArray(localctx.(*AsignacionContext).Get_arreglo().GetExp(), localctx.(*AsignacionContext).GetEe().GetExp())

	}

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
		p.SetState(158)
		p.Match(gramaticaParserTKR_break)
	}
	{
		p.SetState(159)
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

func (s *DeclaracionContext) TKR_vec() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_vec, 0)
}

func (s *DeclaracionContext) TK_menor() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_menor, 0)
}

func (s *DeclaracionContext) TK_mayor() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_mayor, 0)
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

	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(163)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(164)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(165)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(166)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(167)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(168)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(169)
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
			p.SetState(172)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(173)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(174)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(175)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(176)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(177)
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
			p.SetState(180)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(181)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(182)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(183)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(184)
			p.Match(gramaticaParserTK_llave_apertura)
		}
		{
			p.SetState(185)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(186)
			p.Match(gramaticaParserTK_pcoma)
		}
		{
			p.SetState(187)

			var _x = p.Valores()

			localctx.(*DeclaracionContext).cant = _x
		}
		{
			p.SetState(188)
			p.Match(gramaticaParserTK_llave_cierre)
		}
		{
			p.SetState(189)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(190)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(191)
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
			p.SetState(194)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(195)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(196)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(197)
			p.Match(gramaticaParserTK_llave_apertura)
		}
		{
			p.SetState(198)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(199)
			p.Match(gramaticaParserTK_pcoma)
		}
		{
			p.SetState(200)

			var _x = p.Valores()

			localctx.(*DeclaracionContext).cant = _x
		}
		{
			p.SetState(201)
			p.Match(gramaticaParserTK_llave_cierre)
		}
		{
			p.SetState(202)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(203)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(204)
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
			p.SetState(207)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(208)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(209)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(210)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(211)
			p.Match(gramaticaParserTKR_vec)
		}
		{
			p.SetState(212)
			p.Match(gramaticaParserTK_menor)
		}
		{
			p.SetState(213)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(214)
			p.Match(gramaticaParserTK_mayor)
		}
		{
			p.SetState(215)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(216)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(217)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracionVector((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipovariable().GetTipovar(), localctx.(*DeclaracionContext).Get_expresion().GetExp(), true)

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(220)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(221)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(222)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(223)
			p.Match(gramaticaParserTKR_vec)
		}
		{
			p.SetState(224)
			p.Match(gramaticaParserTK_menor)
		}
		{
			p.SetState(225)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(226)
			p.Match(gramaticaParserTK_mayor)
		}
		{
			p.SetState(227)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(228)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(229)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracionVector((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipovariable().GetTipovar(), localctx.(*DeclaracionContext).Get_expresion().GetExp(), false)

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(232)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(233)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(234)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(235)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(236)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(237)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(238)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracion((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipovariable().GetTipovar(), localctx.(*DeclaracionContext).Get_expresion().GetExp(), false, false, false)

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(241)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(242)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(243)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(244)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(245)
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

	p.SetState(262)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTKR_numericos_enteros:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Match(gramaticaParserTKR_numericos_enteros)
		}
		localctx.(*TipovariableContext).tipovar = Enum.INTEGER

	case gramaticaParserTKR_numericos_flotantes:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(252)
			p.Match(gramaticaParserTKR_numericos_flotantes)
		}
		localctx.(*TipovariableContext).tipovar = Enum.FLOAT

	case gramaticaParserTKR_String:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(254)
			p.Match(gramaticaParserTKR_String)
		}
		localctx.(*TipovariableContext).tipovar = Enum.STRING

	case gramaticaParserTKR_bool:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(256)
			p.Match(gramaticaParserTKR_bool)
		}
		localctx.(*TipovariableContext).tipovar = Enum.BOOLEAN

	case gramaticaParserTKR_char:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(258)
			p.Match(gramaticaParserTKR_char)
		}
		localctx.(*TipovariableContext).tipovar = Enum.CHAR

	case gramaticaParserTKR_usize:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(260)
			p.Match(gramaticaParserTKR_usize)
		}
		localctx.(*TipovariableContext).tipovar = Enum.USIZE

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 20, gramaticaParserRULE_valores)

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

	p.SetState(282)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTK_entero:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)

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
			p.SetState(266)

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
			p.SetState(268)

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
			p.SetState(270)

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
			p.SetState(272)
			p.Match(gramaticaParserTK_amp)
		}
		{
			p.SetState(273)

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
			p.SetState(275)
			p.Match(gramaticaParserTKR_true)
		}
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(true, Enum.BOOLEAN)

	case gramaticaParserTKR_false:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(277)
			p.Match(gramaticaParserTKR_false)
		}
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(false, Enum.BOOLEAN)

	case gramaticaParserTK_id:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(279)

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
	_startState := 22
	p.EnterRecursionRule(localctx, 22, gramaticaParserRULE_arreglo, _p)

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
		p.SetState(285)

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
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewArregloContext(p, _parentctx, _parentState)
			localctx.(*ArregloContext).ar = _prevctx
			p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_arreglo)
			p.SetState(288)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(289)
				p.Match(gramaticaParserTK_llave_apertura)
			}
			{
				p.SetState(290)

				var _x = p.expresion(0)

				localctx.(*ArregloContext)._expresion = _x
			}
			{
				p.SetState(291)
				p.Match(gramaticaParserTK_llave_cierre)
			}
			localctx.(*ArregloContext).exp = Expresiones.NewAccesoArray(localctx.(*ArregloContext).GetAr().GetExp(), localctx.(*ArregloContext).Get_expresion().GetExp())

		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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

func (s *ExpresionContext) TKR_vec() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_vec, 0)
}

func (s *ExpresionContext) AllTK_dosPuntos() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserTK_dosPuntos)
}

func (s *ExpresionContext) TK_dosPuntos(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_dosPuntos, i)
}

func (s *ExpresionContext) TKR_new() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_new, 0)
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
	_startState := 24
	p.EnterRecursionRule(localctx, 24, gramaticaParserRULE_expresion, _p)
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
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(300)
			p.Match(gramaticaParserTK_menos)
		}
		{
			p.SetState(301)

			var _x = p.expresion(27)

			localctx.(*ExpresionContext).e1 = _x
		}
		numero := -1
		numm := Expresion.NuevoPrimitivo(numero, Enum.INTEGER)
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), numm, Enum.MULTIPLICACION)

	case 2:
		{
			p.SetState(304)
			p.Match(gramaticaParserTKR_pow)
		}
		{
			p.SetState(305)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(306)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(307)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(308)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e2 = _x
		}
		{
			p.SetState(309)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.POW)

	case 3:
		{
			p.SetState(312)
			p.Match(gramaticaParserTKR_powf)
		}
		{
			p.SetState(313)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(314)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(315)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(316)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e2 = _x
		}
		{
			p.SetState(317)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.POWF)

	case 4:
		{
			p.SetState(320)
			p.Match(gramaticaParserTK_sig_admiracion)
		}
		{
			p.SetState(321)

			var _x = p.expresion(14)

			localctx.(*ExpresionContext).e1 = _x
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE1().GetExp(), Enum.NOT)

	case 5:
		{
			p.SetState(324)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(325)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).va = _x
		}
		{
			p.SetState(326)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = localctx.(*ExpresionContext).GetVa().GetExp()

	case 6:
		{
			p.SetState(329)

			var _x = p.Valores()

			localctx.(*ExpresionContext).val = _x
		}
		{
			p.SetState(330)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(331)
			p.Match(gramaticaParserTKR_numericos_enteros)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NewAsi64(localctx.(*ExpresionContext).GetVal().GetExp())

	case 7:
		{
			p.SetState(334)

			var _x = p.Valores()

			localctx.(*ExpresionContext).val = _x
		}
		{
			p.SetState(335)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(336)
			p.Match(gramaticaParserTKR_numericos_flotantes)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NewAsf64(localctx.(*ExpresionContext).GetVal().GetExp())

	case 8:
		{
			p.SetState(339)
			p.Match(gramaticaParserTK_llave_apertura)
		}
		{
			p.SetState(340)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(341)
			p.Match(gramaticaParserTK_pcoma)
		}
		{
			p.SetState(342)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e2 = _x
		}
		{
			p.SetState(343)
			p.Match(gramaticaParserTK_llave_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresion.NewArray(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), nil, Enum.MULTIPLE)

	case 9:
		{
			p.SetState(346)
			p.Match(gramaticaParserTK_llave_apertura)
		}
		{
			p.SetState(347)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(348)

			var _x = p.Impresionexpresion()

			localctx.(*ExpresionContext).l1 = _x
		}
		{
			p.SetState(349)
			p.Match(gramaticaParserTK_llave_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresion.NewArray(localctx.(*ExpresionContext).GetE1().GetExp(), nil, localctx.(*ExpresionContext).GetL1().GetLista(), Enum.NORMAL)

	case 10:
		{
			p.SetState(352)
			p.Match(gramaticaParserTKR_vec)
		}
		{
			p.SetState(353)
			p.Match(gramaticaParserTK_sig_admiracion)
		}
		{
			p.SetState(354)
			p.Match(gramaticaParserTK_llave_apertura)
		}
		{
			p.SetState(355)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(356)
			p.Match(gramaticaParserTK_pcoma)
		}
		{
			p.SetState(357)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e2 = _x
		}
		{
			p.SetState(358)
			p.Match(gramaticaParserTK_llave_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresion.NewVector(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), nil, Enum.MULTIPLE)

	case 11:
		{
			p.SetState(361)
			p.Match(gramaticaParserTKR_vec)
		}
		{
			p.SetState(362)
			p.Match(gramaticaParserTK_sig_admiracion)
		}
		{
			p.SetState(363)
			p.Match(gramaticaParserTK_llave_apertura)
		}
		{
			p.SetState(364)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(365)

			var _x = p.Impresionexpresion()

			localctx.(*ExpresionContext).l1 = _x
		}
		{
			p.SetState(366)
			p.Match(gramaticaParserTK_llave_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresion.NewVector(localctx.(*ExpresionContext).GetE1().GetExp(), nil, localctx.(*ExpresionContext).GetL1().GetLista(), Enum.NORMAL)

	case 12:
		{
			p.SetState(369)
			p.Match(gramaticaParserTKR_vec)
		}
		{
			p.SetState(370)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(371)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(372)
			p.Match(gramaticaParserTKR_new)
		}
		{
			p.SetState(373)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(374)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresion.NewVector(localctx.(*ExpresionContext).GetE1().GetExp(), nil, localctx.(*ExpresionContext).GetL1().GetLista(), Enum.NORMAL)

	case 13:
		{
			p.SetState(376)

			var _x = p.Valores()

			localctx.(*ExpresionContext).vall = _x
		}
		localctx.(*ExpresionContext).exp = localctx.(*ExpresionContext).GetVall().GetExp()
		fmt.Println(localctx.(*ExpresionContext).exp)

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(454)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(381)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
				}
				{
					p.SetState(382)

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
					p.SetState(383)

					var _x = p.expresion(27)

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
				p.SetState(386)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
				}
				{
					p.SetState(387)

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
					p.SetState(388)

					var _x = p.expresion(26)

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
				p.SetState(391)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(392)
					p.Match(gramaticaParserTK_menor)
				}
				{
					p.SetState(393)

					var _x = p.expresion(23)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.MENOR_QUE)

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(396)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(397)
					p.Match(gramaticaParserTK_mayor)
				}
				{
					p.SetState(398)

					var _x = p.expresion(22)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.MAYOR_QUE)

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(401)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(402)
					p.Match(gramaticaParserTK_mayor_igual)
				}
				{
					p.SetState(403)

					var _x = p.expresion(21)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.MAYOR_IGUAL)

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(406)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(407)
					p.Match(gramaticaParserTK_menor_igual)
				}
				{
					p.SetState(408)

					var _x = p.expresion(20)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.MENOR_IGUAL)

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(411)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(412)
					p.Match(gramaticaParserTK_igualacion)
				}
				{
					p.SetState(413)

					var _x = p.expresion(19)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.IGUALDAD)

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(416)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(417)
					p.Match(gramaticaParserTK_diferente)
				}
				{
					p.SetState(418)

					var _x = p.expresion(18)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.DESIGUAL)

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(421)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(422)
					p.Match(gramaticaParserTK_or)
				}
				{
					p.SetState(423)

					var _x = p.expresion(17)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.OR)

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(426)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(427)
					p.Match(gramaticaParserTK_and)
				}
				{
					p.SetState(428)

					var _x = p.expresion(16)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Enum.AND)

			case 11:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).vll = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(431)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(432)
					p.Match(gramaticaParserTK_punto)
				}
				{
					p.SetState(433)
					p.Match(gramaticaParserTKR_abs)
				}
				{
					p.SetState(434)
					p.Match(gramaticaParserTK_par_apertura)
				}
				{
					p.SetState(435)
					p.Match(gramaticaParserTK_par_cierre)
				}
				localctx.(*ExpresionContext).exp = Expresiones.Newabs(localctx.(*ExpresionContext).GetVll().GetExp())

			case 12:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(437)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(438)
					p.Match(gramaticaParserTK_punto)
				}
				{
					p.SetState(439)
					p.Match(gramaticaParserTKR_sqrt)
				}
				{
					p.SetState(440)
					p.Match(gramaticaParserTK_par_apertura)
				}
				{
					p.SetState(441)
					p.Match(gramaticaParserTK_par_cierre)
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE1().GetExp(), Enum.MULTIPLICACION)

			case 13:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).vll = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(443)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(444)
					p.Match(gramaticaParserTK_punto)
				}
				{
					p.SetState(445)
					p.Match(gramaticaParserTKR_to_string)
				}
				{
					p.SetState(446)
					p.Match(gramaticaParserTK_par_apertura)
				}
				{
					p.SetState(447)
					p.Match(gramaticaParserTK_par_cierre)
				}
				localctx.(*ExpresionContext).exp = Expresiones.NewFto_string(localctx.(*ExpresionContext).GetVll().GetExp())

			case 14:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).vll = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(449)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(450)
					p.Match(gramaticaParserTK_punto)
				}
				{
					p.SetState(451)
					p.Match(gramaticaParserTKR_clone)
				}
				{
					p.SetState(452)
					p.Match(gramaticaParserTK_par_apertura)
				}
				{
					p.SetState(453)
					p.Match(gramaticaParserTK_par_cierre)
				}

			}

		}
		p.SetState(458)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 26, gramaticaParserRULE_impresion)

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

	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(459)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(460)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(461)

			var _x = p.expresion(0)

			localctx.(*ImpresionContext).e1 = _x
		}
		{
			p.SetState(462)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(463)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*ImpresionContext).inst = Instruccion.NuevoPrint(localctx.(*ImpresionContext).GetE1().GetExp(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(466)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(467)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(468)

			var _x = p.expresion(0)

			localctx.(*ImpresionContext).e2 = _x
		}
		{
			p.SetState(469)

			var _x = p.Impresionexpresion()

			localctx.(*ImpresionContext).li = _x
		}
		{
			p.SetState(470)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(471)
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
	p.EnterRule(localctx, 28, gramaticaParserRULE_impresionexpresion)
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
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserTK_coma {
		{
			p.SetState(476)

			var _x = p.Expcoma()

			localctx.(*ImpresionexpresionContext)._expcoma = _x
		}
		localctx.(*ImpresionexpresionContext).list = append(localctx.(*ImpresionexpresionContext).list, localctx.(*ImpresionexpresionContext)._expcoma)

		p.SetState(481)
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
	p.EnterRule(localctx, 30, gramaticaParserRULE_expcoma)

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
		p.SetState(484)
		p.Match(gramaticaParserTK_coma)
	}
	{
		p.SetState(485)

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
	p.EnterRule(localctx, 32, gramaticaParserRULE_condicionales)

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
		p.SetState(488)

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
	p.EnterRule(localctx, 34, gramaticaParserRULE_funcionif)

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

	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(491)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(492)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(493)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).ee = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetEe().GetLista(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(496)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(497)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(498)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).e5 = _x
		}
		{
			p.SetState(499)
			p.Match(gramaticaParserTKR_else)
		}
		{
			p.SetState(500)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b2 = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetE5().GetLista(), nil, localctx.(*FuncionifContext).GetB2().GetLista())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(503)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(504)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(505)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b1 = _x
		}
		{
			p.SetState(506)

			var _x = p.Listaelseif()

			localctx.(*FuncionifContext)._listaelseif = _x
		}
		{
			p.SetState(507)
			p.Match(gramaticaParserTKR_else)
		}
		{
			p.SetState(508)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b2 = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetB1().GetLista(), localctx.(*FuncionifContext).Get_listaelseif().GetLista(), localctx.(*FuncionifContext).GetB2().GetLista())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(511)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(512)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(513)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b1 = _x
		}
		{
			p.SetState(514)

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
	p.EnterRule(localctx, 36, gramaticaParserRULE_funcionelseif)

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
		p.SetState(519)
		p.Match(gramaticaParserTKR_elseif)
	}
	{
		p.SetState(520)

		var _x = p.expresion(0)

		localctx.(*FuncionelseifContext).e1 = _x
	}
	{
		p.SetState(521)

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
	p.EnterRule(localctx, 38, gramaticaParserRULE_listaelseif)
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
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserTKR_elseif {
		{
			p.SetState(524)

			var _x = p.Funcionelseif()

			localctx.(*ListaelseifContext)._funcionelseif = _x
		}
		localctx.(*ListaelseifContext).list = append(localctx.(*ListaelseifContext).list, localctx.(*ListaelseifContext)._funcionelseif)

		p.SetState(529)
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
	p.EnterRule(localctx, 40, gramaticaParserRULE_bloque)

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

	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(532)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(533)

			var _x = p.Instrucciones()

			localctx.(*BloqueContext)._instrucciones = _x
		}
		{
			p.SetState(534)
			p.Match(gramaticaParserTK_corchete_cierre)
		}
		localctx.(*BloqueContext).lista = localctx.(*BloqueContext).Get_instrucciones().GetLista()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(537)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(538)
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
	p.EnterRule(localctx, 42, gramaticaParserRULE_bucles)

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

	p.SetState(548)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTKR_while:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(542)

			var _x = p.Fwhile()

			localctx.(*BuclesContext)._fwhile = _x
		}
		localctx.(*BuclesContext).inst = localctx.(*BuclesContext).Get_fwhile().GetInst()

	case gramaticaParserTKR_loop:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(545)

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
	p.EnterRule(localctx, 44, gramaticaParserRULE_fwhile)

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
		p.SetState(550)
		p.Match(gramaticaParserTKR_while)
	}
	{
		p.SetState(551)

		var _x = p.expresion(0)

		localctx.(*FwhileContext).e1 = _x
	}
	{
		p.SetState(552)

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
	p.EnterRule(localctx, 46, gramaticaParserRULE_floop)

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
		p.SetState(555)
		p.Match(gramaticaParserTKR_loop)
	}
	{
		p.SetState(556)

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

	case 11:
		var t *ArregloContext = nil
		if localctx != nil {
			t = localctx.(*ArregloContext)
		}
		return p.Arreglo_Sempred(t, predIndex)

	case 12:
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
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
