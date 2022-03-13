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
import arrayList "github.com/colegno/arraylist"

//import "proyecto1/Operaciones"
//import "proyecto1/Expresion"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 75, 315,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 43, 10, 3, 12, 3, 14, 3, 46,
	11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 65, 10, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 101, 10, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 115,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8,
	139, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 181, 10, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9,
	233, 10, 9, 12, 9, 14, 9, 236, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 251, 10,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 267, 10, 11, 12, 11, 14, 11, 270, 11,
	11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 5, 13, 295, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 5, 16, 313, 10, 16, 3, 16, 2, 4, 16, 20, 17, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 2, 4, 3, 2, 25, 27, 3, 2, 23, 24, 2, 341,
	2, 32, 3, 2, 2, 2, 4, 44, 3, 2, 2, 2, 6, 64, 3, 2, 2, 2, 8, 100, 3, 2,
	2, 2, 10, 114, 3, 2, 2, 2, 12, 116, 3, 2, 2, 2, 14, 138, 3, 2, 2, 2, 16,
	180, 3, 2, 2, 2, 18, 250, 3, 2, 2, 2, 20, 252, 3, 2, 2, 2, 22, 271, 3,
	2, 2, 2, 24, 294, 3, 2, 2, 2, 26, 296, 3, 2, 2, 2, 28, 301, 3, 2, 2, 2,
	30, 312, 3, 2, 2, 2, 32, 33, 7, 52, 2, 2, 33, 34, 7, 66, 2, 2, 34, 35,
	7, 20, 2, 2, 35, 36, 7, 21, 2, 2, 36, 37, 7, 10, 2, 2, 37, 38, 5, 4, 3,
	2, 38, 39, 7, 11, 2, 2, 39, 40, 8, 2, 1, 2, 40, 3, 3, 2, 2, 2, 41, 43,
	5, 6, 4, 2, 42, 41, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2,
	44, 45, 3, 2, 2, 2, 45, 47, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 48, 8,
	3, 1, 2, 48, 5, 3, 2, 2, 2, 49, 50, 5, 16, 9, 2, 50, 51, 8, 4, 1, 2, 51,
	65, 3, 2, 2, 2, 52, 53, 5, 18, 10, 2, 53, 54, 8, 4, 1, 2, 54, 65, 3, 2,
	2, 2, 55, 56, 5, 8, 5, 2, 56, 57, 8, 4, 1, 2, 57, 65, 3, 2, 2, 2, 58, 59,
	5, 12, 7, 2, 59, 60, 8, 4, 1, 2, 60, 65, 3, 2, 2, 2, 61, 62, 5, 22, 12,
	2, 62, 63, 8, 4, 1, 2, 63, 65, 3, 2, 2, 2, 64, 49, 3, 2, 2, 2, 64, 52,
	3, 2, 2, 2, 64, 55, 3, 2, 2, 2, 64, 58, 3, 2, 2, 2, 64, 61, 3, 2, 2, 2,
	65, 7, 3, 2, 2, 2, 66, 67, 7, 45, 2, 2, 67, 68, 7, 46, 2, 2, 68, 69, 7,
	72, 2, 2, 69, 70, 7, 14, 2, 2, 70, 71, 5, 10, 6, 2, 71, 72, 7, 22, 2, 2,
	72, 73, 5, 16, 9, 2, 73, 74, 7, 16, 2, 2, 74, 75, 8, 5, 1, 2, 75, 101,
	3, 2, 2, 2, 76, 77, 7, 45, 2, 2, 77, 78, 7, 46, 2, 2, 78, 79, 7, 72, 2,
	2, 79, 80, 7, 22, 2, 2, 80, 81, 5, 16, 9, 2, 81, 82, 7, 16, 2, 2, 82, 83,
	8, 5, 1, 2, 83, 101, 3, 2, 2, 2, 84, 85, 7, 45, 2, 2, 85, 86, 7, 72, 2,
	2, 86, 87, 7, 14, 2, 2, 87, 88, 5, 10, 6, 2, 88, 89, 7, 22, 2, 2, 89, 90,
	5, 16, 9, 2, 90, 91, 7, 16, 2, 2, 91, 92, 8, 5, 1, 2, 92, 101, 3, 2, 2,
	2, 93, 94, 7, 45, 2, 2, 94, 95, 7, 72, 2, 2, 95, 96, 7, 22, 2, 2, 96, 97,
	5, 16, 9, 2, 97, 98, 7, 16, 2, 2, 98, 99, 8, 5, 1, 2, 99, 101, 3, 2, 2,
	2, 100, 66, 3, 2, 2, 2, 100, 76, 3, 2, 2, 2, 100, 84, 3, 2, 2, 2, 100,
	93, 3, 2, 2, 2, 101, 9, 3, 2, 2, 2, 102, 103, 7, 35, 2, 2, 103, 115, 8,
	6, 1, 2, 104, 105, 7, 36, 2, 2, 105, 115, 8, 6, 1, 2, 106, 107, 7, 43,
	2, 2, 107, 115, 8, 6, 1, 2, 108, 109, 7, 40, 2, 2, 109, 115, 8, 6, 1, 2,
	110, 111, 7, 41, 2, 2, 111, 115, 8, 6, 1, 2, 112, 113, 7, 44, 2, 2, 113,
	115, 8, 6, 1, 2, 114, 102, 3, 2, 2, 2, 114, 104, 3, 2, 2, 2, 114, 106,
	3, 2, 2, 2, 114, 108, 3, 2, 2, 2, 114, 110, 3, 2, 2, 2, 114, 112, 3, 2,
	2, 2, 115, 11, 3, 2, 2, 2, 116, 117, 7, 72, 2, 2, 117, 118, 7, 22, 2, 2,
	118, 119, 5, 16, 9, 2, 119, 120, 7, 16, 2, 2, 120, 121, 8, 7, 1, 2, 121,
	13, 3, 2, 2, 2, 122, 123, 7, 70, 2, 2, 123, 139, 8, 8, 1, 2, 124, 125,
	7, 71, 2, 2, 125, 139, 8, 8, 1, 2, 126, 127, 7, 74, 2, 2, 127, 139, 8,
	8, 1, 2, 128, 129, 7, 75, 2, 2, 129, 139, 8, 8, 1, 2, 130, 131, 7, 42,
	2, 2, 131, 139, 8, 8, 1, 2, 132, 133, 7, 50, 2, 2, 133, 139, 8, 8, 1, 2,
	134, 135, 7, 51, 2, 2, 135, 139, 8, 8, 1, 2, 136, 137, 7, 72, 2, 2, 137,
	139, 8, 8, 1, 2, 138, 122, 3, 2, 2, 2, 138, 124, 3, 2, 2, 2, 138, 126,
	3, 2, 2, 2, 138, 128, 3, 2, 2, 2, 138, 130, 3, 2, 2, 2, 138, 132, 3, 2,
	2, 2, 138, 134, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139, 15, 3, 2, 2, 2,
	140, 141, 8, 9, 1, 2, 141, 142, 7, 24, 2, 2, 142, 143, 5, 16, 9, 20, 143,
	144, 8, 9, 1, 2, 144, 181, 3, 2, 2, 2, 145, 146, 7, 37, 2, 2, 146, 147,
	7, 20, 2, 2, 147, 148, 5, 16, 9, 2, 148, 149, 7, 15, 2, 2, 149, 150, 5,
	16, 9, 2, 150, 151, 7, 21, 2, 2, 151, 152, 8, 9, 1, 2, 152, 181, 3, 2,
	2, 2, 153, 154, 7, 39, 2, 2, 154, 155, 7, 20, 2, 2, 155, 156, 5, 16, 9,
	2, 156, 157, 7, 15, 2, 2, 157, 158, 5, 16, 9, 2, 158, 159, 7, 21, 2, 2,
	159, 160, 8, 9, 1, 2, 160, 181, 3, 2, 2, 2, 161, 162, 7, 30, 2, 2, 162,
	163, 5, 16, 9, 7, 163, 164, 8, 9, 1, 2, 164, 181, 3, 2, 2, 2, 165, 166,
	7, 20, 2, 2, 166, 167, 5, 16, 9, 2, 167, 168, 7, 21, 2, 2, 168, 181, 3,
	2, 2, 2, 169, 170, 5, 14, 8, 2, 170, 171, 7, 48, 2, 2, 171, 172, 7, 35,
	2, 2, 172, 181, 3, 2, 2, 2, 173, 174, 5, 14, 8, 2, 174, 175, 7, 48, 2,
	2, 175, 176, 7, 36, 2, 2, 176, 181, 3, 2, 2, 2, 177, 178, 5, 14, 8, 2,
	178, 179, 8, 9, 1, 2, 179, 181, 3, 2, 2, 2, 180, 140, 3, 2, 2, 2, 180,
	145, 3, 2, 2, 2, 180, 153, 3, 2, 2, 2, 180, 161, 3, 2, 2, 2, 180, 165,
	3, 2, 2, 2, 180, 169, 3, 2, 2, 2, 180, 173, 3, 2, 2, 2, 180, 177, 3, 2,
	2, 2, 181, 234, 3, 2, 2, 2, 182, 183, 12, 19, 2, 2, 183, 184, 9, 2, 2,
	2, 184, 185, 5, 16, 9, 20, 185, 186, 8, 9, 1, 2, 186, 233, 3, 2, 2, 2,
	187, 188, 12, 18, 2, 2, 188, 189, 9, 3, 2, 2, 189, 190, 5, 16, 9, 19, 190,
	191, 8, 9, 1, 2, 191, 233, 3, 2, 2, 2, 192, 193, 12, 15, 2, 2, 193, 194,
	7, 17, 2, 2, 194, 195, 5, 16, 9, 16, 195, 196, 8, 9, 1, 2, 196, 233, 3,
	2, 2, 2, 197, 198, 12, 14, 2, 2, 198, 199, 7, 18, 2, 2, 199, 200, 5, 16,
	9, 15, 200, 201, 8, 9, 1, 2, 201, 233, 3, 2, 2, 2, 202, 203, 12, 13, 2,
	2, 203, 204, 7, 8, 2, 2, 204, 205, 5, 16, 9, 14, 205, 206, 8, 9, 1, 2,
	206, 233, 3, 2, 2, 2, 207, 208, 12, 12, 2, 2, 208, 209, 7, 9, 2, 2, 209,
	210, 5, 16, 9, 13, 210, 211, 8, 9, 1, 2, 211, 233, 3, 2, 2, 2, 212, 213,
	12, 11, 2, 2, 213, 214, 7, 6, 2, 2, 214, 215, 5, 16, 9, 12, 215, 216, 8,
	9, 1, 2, 216, 233, 3, 2, 2, 2, 217, 218, 12, 10, 2, 2, 218, 219, 7, 7,
	2, 2, 219, 220, 5, 16, 9, 11, 220, 221, 8, 9, 1, 2, 221, 233, 3, 2, 2,
	2, 222, 223, 12, 9, 2, 2, 223, 224, 7, 4, 2, 2, 224, 225, 5, 16, 9, 10,
	225, 226, 8, 9, 1, 2, 226, 233, 3, 2, 2, 2, 227, 228, 12, 8, 2, 2, 228,
	229, 7, 5, 2, 2, 229, 230, 5, 16, 9, 9, 230, 231, 8, 9, 1, 2, 231, 233,
	3, 2, 2, 2, 232, 182, 3, 2, 2, 2, 232, 187, 3, 2, 2, 2, 232, 192, 3, 2,
	2, 2, 232, 197, 3, 2, 2, 2, 232, 202, 3, 2, 2, 2, 232, 207, 3, 2, 2, 2,
	232, 212, 3, 2, 2, 2, 232, 217, 3, 2, 2, 2, 232, 222, 3, 2, 2, 2, 232,
	227, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 235,
	3, 2, 2, 2, 235, 17, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 237, 238, 7, 49,
	2, 2, 238, 239, 7, 20, 2, 2, 239, 240, 5, 16, 9, 2, 240, 241, 7, 21, 2,
	2, 241, 242, 7, 16, 2, 2, 242, 243, 8, 10, 1, 2, 243, 251, 3, 2, 2, 2,
	244, 245, 7, 49, 2, 2, 245, 246, 7, 20, 2, 2, 246, 247, 5, 16, 9, 2, 247,
	248, 5, 20, 11, 2, 248, 249, 7, 16, 2, 2, 249, 251, 3, 2, 2, 2, 250, 237,
	3, 2, 2, 2, 250, 244, 3, 2, 2, 2, 251, 19, 3, 2, 2, 2, 252, 253, 8, 11,
	1, 2, 253, 254, 7, 15, 2, 2, 254, 255, 5, 16, 9, 2, 255, 256, 7, 21, 2,
	2, 256, 257, 7, 16, 2, 2, 257, 258, 8, 11, 1, 2, 258, 268, 3, 2, 2, 2,
	259, 260, 12, 4, 2, 2, 260, 261, 7, 15, 2, 2, 261, 262, 5, 16, 9, 2, 262,
	263, 7, 21, 2, 2, 263, 264, 7, 16, 2, 2, 264, 265, 8, 11, 1, 2, 265, 267,
	3, 2, 2, 2, 266, 259, 3, 2, 2, 2, 267, 270, 3, 2, 2, 2, 268, 266, 3, 2,
	2, 2, 268, 269, 3, 2, 2, 2, 269, 21, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2,
	271, 272, 5, 24, 13, 2, 272, 273, 8, 12, 1, 2, 273, 23, 3, 2, 2, 2, 274,
	275, 7, 67, 2, 2, 275, 276, 5, 16, 9, 2, 276, 277, 5, 30, 16, 2, 277, 278,
	8, 13, 1, 2, 278, 295, 3, 2, 2, 2, 279, 280, 7, 67, 2, 2, 280, 281, 5,
	16, 9, 2, 281, 282, 5, 30, 16, 2, 282, 283, 7, 69, 2, 2, 283, 284, 5, 30,
	16, 2, 284, 285, 8, 13, 1, 2, 285, 295, 3, 2, 2, 2, 286, 287, 7, 67, 2,
	2, 287, 288, 5, 16, 9, 2, 288, 289, 5, 30, 16, 2, 289, 290, 5, 28, 15,
	2, 290, 291, 7, 69, 2, 2, 291, 292, 5, 30, 16, 2, 292, 293, 8, 13, 1, 2,
	293, 295, 3, 2, 2, 2, 294, 274, 3, 2, 2, 2, 294, 279, 3, 2, 2, 2, 294,
	286, 3, 2, 2, 2, 295, 25, 3, 2, 2, 2, 296, 297, 7, 68, 2, 2, 297, 298,
	5, 16, 9, 2, 298, 299, 5, 30, 16, 2, 299, 300, 8, 14, 1, 2, 300, 27, 3,
	2, 2, 2, 301, 302, 5, 26, 14, 2, 302, 303, 8, 15, 1, 2, 303, 29, 3, 2,
	2, 2, 304, 305, 7, 10, 2, 2, 305, 306, 5, 4, 3, 2, 306, 307, 7, 11, 2,
	2, 307, 308, 8, 16, 1, 2, 308, 313, 3, 2, 2, 2, 309, 310, 7, 10, 2, 2,
	310, 311, 7, 11, 2, 2, 311, 313, 8, 16, 1, 2, 312, 304, 3, 2, 2, 2, 312,
	309, 3, 2, 2, 2, 313, 31, 3, 2, 2, 2, 14, 44, 64, 100, 114, 138, 180, 232,
	234, 250, 268, 294, 312,
}
var literalNames = []string{
	"", "'->'", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'{'", "'}'",
	"'['", "']'", "':'", "','", "';'", "'<'", "'>'", "'.'", "'('", "')'", "'='",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'", "'!'", "'?'", "", "",
	"", "'i64'", "'f64'", "'pow'", "'vec'", "'powf'", "'bool'", "'char'", "'&str'",
	"'String'", "'usize'", "'let'", "'mut'", "'struct'", "'as'", "'println!'",
	"'true'", "'false'", "'fn'", "'return'", "'abs'", "'sqrt'", "'to_string()'",
	"'clone()'", "'new'", "'len'", "'push'", "'remove'", "'contains'", "'insert'",
	"'capacity'", "'witch_capacity'", "'main'", "'if'", "'else if'", "'else'",
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
	"TKR_if", "TKR_elseif", "TKR_else", "TK_entero", "TK_decimal", "TK_id",
	"Letra", "TK_cadena", "TK_caracter",
}

var ruleNames = []string{
	"start", "instrucciones", "instruccion", "declaracion", "tipovariable",
	"identificadores", "valores", "expresion", "impresion", "impresioncomas",
	"condicionales", "funcionif", "funcionelseif", "listaelseif", "bloque",
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
	gramaticaParserTK_entero               = 68
	gramaticaParserTK_decimal              = 69
	gramaticaParserTK_id                   = 70
	gramaticaParserLetra                   = 71
	gramaticaParserTK_cadena               = 72
	gramaticaParserTK_caracter             = 73
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
	gramaticaParserRULE_condicionales   = 10
	gramaticaParserRULE_funcionif       = 11
	gramaticaParserRULE_funcionelseif   = 12
	gramaticaParserRULE_listaelseif     = 13
	gramaticaParserRULE_bloque          = 14
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
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

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	listainstrucciones *arrayList.List
	_instrucciones     IInstruccionesContext
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

func (s *StartContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *StartContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *StartContext) GetListainstrucciones() *arrayList.List { return s.listainstrucciones }

func (s *StartContext) SetListainstrucciones(v *arrayList.List) { s.listainstrucciones = v }

func (s *StartContext) TKR_fn() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_fn, 0)
}

func (s *StartContext) TKR_main() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_main, 0)
}

func (s *StartContext) TK_par_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_apertura, 0)
}

func (s *StartContext) TK_par_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_par_cierre, 0)
}

func (s *StartContext) TK_corchete_apertura() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_corchete_apertura, 0)
}

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) TK_corchete_cierre() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTK_corchete_cierre, 0)
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
		p.SetState(30)
		p.Match(gramaticaParserTKR_fn)
	}
	{
		p.SetState(31)
		p.Match(gramaticaParserTKR_main)
	}
	{
		p.SetState(32)
		p.Match(gramaticaParserTK_par_apertura)
	}
	{
		p.SetState(33)
		p.Match(gramaticaParserTK_par_cierre)
	}
	{
		p.SetState(34)
		p.Match(gramaticaParserTK_corchete_apertura)
	}
	{
		p.SetState(35)

		var _x = p.Instrucciones()

		localctx.(*StartContext)._instrucciones = _x
	}
	{
		p.SetState(36)
		p.Match(gramaticaParserTK_corchete_cierre)
	}
	localctx.(*StartContext).listainstrucciones = localctx.(*StartContext).Get_instrucciones().GetLista()

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
	p.EnterRule(localctx, 2, gramaticaParserRULE_instrucciones)

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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-18)&-(0x1f+1)) == 0 && ((1<<uint((_la-18)))&((1<<(gramaticaParserTK_par_apertura-18))|(1<<(gramaticaParserTK_menos-18))|(1<<(gramaticaParserTK_sig_admiracion-18))|(1<<(gramaticaParserTKR_pow-18))|(1<<(gramaticaParserTKR_powf-18))|(1<<(gramaticaParserTKR_Str-18))|(1<<(gramaticaParserTKR_let-18))|(1<<(gramaticaParserTKR_println-18))|(1<<(gramaticaParserTKR_true-18))|(1<<(gramaticaParserTKR_false-18)))) != 0) || (((_la-65)&-(0x1f+1)) == 0 && ((1<<uint((_la-65)))&((1<<(gramaticaParserTKR_if-65))|(1<<(gramaticaParserTK_entero-65))|(1<<(gramaticaParserTK_decimal-65))|(1<<(gramaticaParserTK_id-65))|(1<<(gramaticaParserTK_cadena-65))|(1<<(gramaticaParserTK_caracter-65)))) != 0) {
		{
			p.SetState(39)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(44)
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

func (s *InstruccionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *InstruccionContext) Set_impresion(v IImpresionContext) { s._impresion = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_identificadores(v IIdentificadoresContext) { s._identificadores = v }

func (s *InstruccionContext) Set_condicionales(v ICondicionalesContext) { s._condicionales = v }

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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)

			var _x = p.expresion(0)

			localctx.(*InstruccionContext)._expresion = _x
		}
		fmt.Println("mensaje en instrucciones: ", localctx.(*InstruccionContext).Get_expresion().GetExp())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)

			var _x = p.Impresion()

			localctx.(*InstruccionContext)._impresion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_impresion().GetInst()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(53)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_declaracion().GetInst()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)

			var _x = p.Identificadores()

			localctx.(*InstruccionContext)._identificadores = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_identificadores().GetInst()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)

			var _x = p.Condicionales()

			localctx.(*InstruccionContext)._condicionales = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_condicionales().GetInst()

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

	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(65)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(66)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(67)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(68)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(69)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(70)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(71)
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
			p.SetState(74)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(75)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(76)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(77)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(78)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(79)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracion((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), Interfaces.SINTIPO, localctx.(*DeclaracionContext).Get_expresion().GetExp(), false, false, true)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(83)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(84)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(85)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(86)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(87)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(88)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracion((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipovariable().GetTipovar(), localctx.(*DeclaracionContext).Get_expresion().GetExp(), false, false, false)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(92)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(93)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(94)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(95)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*DeclaracionContext).inst = Instruccion.NuevaDeclaracion((func() string {
			if localctx.(*DeclaracionContext).GetIdd() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetIdd().GetText()
			}
		}()), Interfaces.SINTIPO, localctx.(*DeclaracionContext).Get_expresion().GetExp(), false, false, false)

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

	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTKR_numericos_enteros:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Match(gramaticaParserTKR_numericos_enteros)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.INTEGER

	case gramaticaParserTKR_numericos_flotantes:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Match(gramaticaParserTKR_numericos_flotantes)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.FLOAT

	case gramaticaParserTKR_String:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.Match(gramaticaParserTKR_String)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.STRING

	case gramaticaParserTKR_bool:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.Match(gramaticaParserTKR_bool)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.BOOLEAN

	case gramaticaParserTKR_char:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(108)
			p.Match(gramaticaParserTKR_char)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.CHAR

	case gramaticaParserTKR_usize:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(110)
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
		p.SetState(114)

		var _m = p.Match(gramaticaParserTK_id)

		localctx.(*IdentificadoresContext).e1 = _m
	}
	{
		p.SetState(115)
		p.Match(gramaticaParserTK_igual)
	}
	{
		p.SetState(116)

		var _x = p.expresion(0)

		localctx.(*IdentificadoresContext).e2 = _x
	}
	{
		p.SetState(117)
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

	// Get_TK_id returns the _TK_id token.
	Get_TK_id() antlr.Token

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

	// Set_TK_id sets the _TK_id token.
	Set_TK_id(antlr.Token)

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
	_TK_id       antlr.Token
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

func (s *ValoresContext) Get_TK_id() antlr.Token { return s._TK_id }

func (s *ValoresContext) Set_TK_entero(v antlr.Token) { s._TK_entero = v }

func (s *ValoresContext) Set_TK_decimal(v antlr.Token) { s._TK_decimal = v }

func (s *ValoresContext) Set_TK_cadena(v antlr.Token) { s._TK_cadena = v }

func (s *ValoresContext) Set_TK_caracter(v antlr.Token) { s._TK_caracter = v }

func (s *ValoresContext) Set_TKR_Str(v antlr.Token) { s._TKR_Str = v }

func (s *ValoresContext) Set_TK_id(v antlr.Token) { s._TK_id = v }

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

func (s *ValoresContext) TKR_Str() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTKR_Str, 0)
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

	p.SetState(136)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTK_entero:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)

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
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)

	case gramaticaParserTK_decimal:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)

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
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(decimal, Interfaces.FLOAT)

	case gramaticaParserTK_cadena:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)

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
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(str, Interfaces.STRING)

	case gramaticaParserTK_caracter:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)

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
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(str, Interfaces.CHAR)

	case gramaticaParserTKR_Str:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)

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
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(str, Interfaces.STR)

	case gramaticaParserTKR_true:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(130)
			p.Match(gramaticaParserTKR_true)
		}
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(true, Interfaces.BOOLEAN)

	case gramaticaParserTKR_false:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(132)
			p.Match(gramaticaParserTKR_false)
		}
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(false, Interfaces.BOOLEAN)

	case gramaticaParserTK_id:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(134)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*ValoresContext)._TK_id = _m
		}
		localctx.(*ValoresContext).exp = Expresiones.NewLlamarvariable((func() string {
			if localctx.(*ValoresContext).Get_TK_id() == nil {
				return ""
			} else {
				return localctx.(*ValoresContext).Get_TK_id().GetText()
			}
		}()))

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

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// GetVall returns the vall rule contexts.
	GetVall() IValoresContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

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
	e2     IExpresionContext
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

func (s *ExpresionContext) GetE2() IExpresionContext { return s.e2 }

func (s *ExpresionContext) GetVall() IValoresContext { return s.vall }

func (s *ExpresionContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *ExpresionContext) SetE2(v IExpresionContext) { s.e2 = v }

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
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(139)
			p.Match(gramaticaParserTK_menos)
		}
		{
			p.SetState(140)

			var _x = p.expresion(18)

			localctx.(*ExpresionContext).e1 = _x
		}
		numero := -1
		numm := Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), numm, Interfaces.MULTIPLICACION)

	case 2:
		{
			p.SetState(143)
			p.Match(gramaticaParserTKR_pow)
		}
		{
			p.SetState(144)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(145)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(146)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(147)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e2 = _x
		}
		{
			p.SetState(148)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.POW)

	case 3:
		{
			p.SetState(151)
			p.Match(gramaticaParserTKR_powf)
		}
		{
			p.SetState(152)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(153)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(154)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(155)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e2 = _x
		}
		{
			p.SetState(156)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.POWF)

	case 4:
		{
			p.SetState(159)
			p.Match(gramaticaParserTK_sig_admiracion)
		}
		{
			p.SetState(160)

			var _x = p.expresion(5)

			localctx.(*ExpresionContext).e1 = _x
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE1().GetExp(), Interfaces.NOT)

	case 5:
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

	case 6:
		{
			p.SetState(167)
			p.Valores()
		}
		{
			p.SetState(168)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(169)
			p.Match(gramaticaParserTKR_numericos_enteros)
		}

	case 7:
		{
			p.SetState(171)
			p.Valores()
		}
		{
			p.SetState(172)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(173)
			p.Match(gramaticaParserTKR_numericos_flotantes)
		}

	case 8:
		{
			p.SetState(175)

			var _x = p.Valores()

			localctx.(*ExpresionContext).vall = _x
		}
		localctx.(*ExpresionContext).exp = localctx.(*ExpresionContext).GetVall().GetExp()
		fmt.Println(localctx.(*ExpresionContext).exp)

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(180)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(181)

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
					p.SetState(182)

					var _x = p.expresion(18)

					localctx.(*ExpresionContext).e2 = _x
				}

				if (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()) == "*" {
					localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.MULTIPLICACION)
				} else if (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()) == "%" {
					localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.MODULO)
				} else {
					localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.DIVISION)
				}

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(185)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(186)

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
					p.SetState(187)

					var _x = p.expresion(17)

					localctx.(*ExpresionContext).e2 = _x
				}

				if (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()) == "+" {
					localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.SUMA)
				} else {
					localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.RESTA)
				}

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(190)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(191)
					p.Match(gramaticaParserTK_menor)
				}
				{
					p.SetState(192)

					var _x = p.expresion(14)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.MENOR_QUE)

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(196)
					p.Match(gramaticaParserTK_mayor)
				}
				{
					p.SetState(197)

					var _x = p.expresion(13)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.MAYOR_QUE)

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(201)
					p.Match(gramaticaParserTK_mayor_igual)
				}
				{
					p.SetState(202)

					var _x = p.expresion(12)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.MAYOR_IGUAL)

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(205)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(206)
					p.Match(gramaticaParserTK_menor_igual)
				}
				{
					p.SetState(207)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.MENOR_IGUAL)

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(210)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(211)
					p.Match(gramaticaParserTK_igualacion)
				}
				{
					p.SetState(212)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.IGUALDAD)

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(216)
					p.Match(gramaticaParserTK_diferente)
				}
				{
					p.SetState(217)

					var _x = p.expresion(9)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.DESIGUAL)

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(221)
					p.Match(gramaticaParserTK_or)
				}
				{
					p.SetState(222)

					var _x = p.expresion(8)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.OR)

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(225)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(226)
					p.Match(gramaticaParserTK_and)
				}
				{
					p.SetState(227)

					var _x = p.expresion(7)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.AND)

			}

		}
		p.SetState(234)
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

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

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

func (s *ImpresionContext) SetE1(v IExpresionContext) { s.e1 = v }

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

	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(236)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(237)

			var _x = p.expresion(0)

			localctx.(*ImpresionContext).e1 = _x
		}
		{
			p.SetState(238)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(239)
			p.Match(gramaticaParserTK_pcoma)
		}
		fmt.Println("Impresion")
		localctx.(*ImpresionContext).inst = Instruccion.NuevoPrint(localctx.(*ImpresionContext).GetE1().GetExp())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(243)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(244)
			p.expresion(0)
		}
		{
			p.SetState(245)
			p.impresioncomas(0)
		}
		{
			p.SetState(246)
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
		p.SetState(251)
		p.Match(gramaticaParserTK_coma)
	}
	{
		p.SetState(252)
		p.expresion(0)
	}
	{
		p.SetState(253)
		p.Match(gramaticaParserTK_par_cierre)
	}
	{
		p.SetState(254)
		p.Match(gramaticaParserTK_pcoma)
	}
	fmt.Println("Impresion especial")

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(266)
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
			p.SetState(257)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(258)
				p.Match(gramaticaParserTK_coma)
			}
			{
				p.SetState(259)
				p.expresion(0)
			}
			{
				p.SetState(260)
				p.Match(gramaticaParserTK_par_cierre)
			}
			{
				p.SetState(261)
				p.Match(gramaticaParserTK_pcoma)
			}
			fmt.Println("Impresion especial")

		}
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

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
	p.EnterRule(localctx, 20, gramaticaParserRULE_condicionales)

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
		p.SetState(269)

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
	p.EnterRule(localctx, 22, gramaticaParserRULE_funcionif)

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

	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(273)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(274)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).ee = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetEe().GetLista(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(278)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(279)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).e5 = _x
		}
		{
			p.SetState(280)
			p.Match(gramaticaParserTKR_else)
		}
		{
			p.SetState(281)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b2 = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetE5().GetLista(), nil, localctx.(*FuncionifContext).GetB2().GetLista())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(284)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(285)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(286)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b1 = _x
		}
		{
			p.SetState(287)

			var _x = p.Listaelseif()

			localctx.(*FuncionifContext)._listaelseif = _x
		}
		{
			p.SetState(288)
			p.Match(gramaticaParserTKR_else)
		}
		{
			p.SetState(289)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b2 = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetB1().GetLista(), localctx.(*FuncionifContext).Get_listaelseif().GetLista(), localctx.(*FuncionifContext).GetB2().GetLista())

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
	p.EnterRule(localctx, 24, gramaticaParserRULE_funcionelseif)

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
		p.SetState(294)
		p.Match(gramaticaParserTKR_elseif)
	}
	{
		p.SetState(295)

		var _x = p.expresion(0)

		localctx.(*FuncionelseifContext).e1 = _x
	}
	{
		p.SetState(296)

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

func (s *ListaelseifContext) Funcionelseif() IFuncionelseifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionelseifContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 26, gramaticaParserRULE_listaelseif)
	localctx.(*ListaelseifContext).lista = arrayList.New()

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
		p.SetState(299)

		var _x = p.Funcionelseif()

		localctx.(*ListaelseifContext)._funcionelseif = _x
	}
	localctx.(*ListaelseifContext).list = append(localctx.(*ListaelseifContext).list, localctx.(*ListaelseifContext)._funcionelseif)

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
	p.EnterRule(localctx, 28, gramaticaParserRULE_bloque)

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

	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(302)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(303)

			var _x = p.Instrucciones()

			localctx.(*BloqueContext)._instrucciones = _x
		}
		{
			p.SetState(304)
			p.Match(gramaticaParserTK_corchete_cierre)
		}
		localctx.(*BloqueContext).lista = localctx.(*BloqueContext).Get_instrucciones().GetLista()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(307)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(308)
			p.Match(gramaticaParserTK_corchete_cierre)
		}
		localctx.(*BloqueContext).lista = arrayList.New()

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
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *gramaticaParser) Impresioncomas_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
