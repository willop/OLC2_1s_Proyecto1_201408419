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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 78, 358,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 53, 10, 3, 12, 3, 14, 3, 56, 11,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 81, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 121, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 135, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 159, 10, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 202,
	10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10,
	254, 10, 10, 12, 10, 14, 10, 257, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11,
	273, 10, 11, 3, 12, 7, 12, 276, 10, 12, 12, 12, 14, 12, 279, 11, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 5, 15, 316, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 17, 7, 17, 324, 10, 17, 12, 17, 14, 17, 327, 11, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 339, 10,
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 347, 10, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 2, 3, 18,
	22, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 2, 4, 3, 2, 25, 27, 3, 2, 23, 24, 2, 384, 2, 42, 3, 2, 2, 2, 4,
	54, 3, 2, 2, 2, 6, 80, 3, 2, 2, 2, 8, 82, 3, 2, 2, 2, 10, 120, 3, 2, 2,
	2, 12, 134, 3, 2, 2, 2, 14, 136, 3, 2, 2, 2, 16, 158, 3, 2, 2, 2, 18, 201,
	3, 2, 2, 2, 20, 272, 3, 2, 2, 2, 22, 277, 3, 2, 2, 2, 24, 282, 3, 2, 2,
	2, 26, 286, 3, 2, 2, 2, 28, 315, 3, 2, 2, 2, 30, 317, 3, 2, 2, 2, 32, 325,
	3, 2, 2, 2, 34, 338, 3, 2, 2, 2, 36, 346, 3, 2, 2, 2, 38, 348, 3, 2, 2,
	2, 40, 353, 3, 2, 2, 2, 42, 43, 7, 52, 2, 2, 43, 44, 7, 66, 2, 2, 44, 45,
	7, 20, 2, 2, 45, 46, 7, 21, 2, 2, 46, 47, 7, 10, 2, 2, 47, 48, 5, 4, 3,
	2, 48, 49, 7, 11, 2, 2, 49, 50, 8, 2, 1, 2, 50, 3, 3, 2, 2, 2, 51, 53,
	5, 6, 4, 2, 52, 51, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2,
	54, 55, 3, 2, 2, 2, 55, 57, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 57, 58, 8,
	3, 1, 2, 58, 5, 3, 2, 2, 2, 59, 60, 5, 18, 10, 2, 60, 61, 8, 4, 1, 2, 61,
	81, 3, 2, 2, 2, 62, 63, 5, 20, 11, 2, 63, 64, 8, 4, 1, 2, 64, 81, 3, 2,
	2, 2, 65, 66, 5, 10, 6, 2, 66, 67, 8, 4, 1, 2, 67, 81, 3, 2, 2, 2, 68,
	69, 5, 14, 8, 2, 69, 70, 8, 4, 1, 2, 70, 81, 3, 2, 2, 2, 71, 72, 5, 26,
	14, 2, 72, 73, 8, 4, 1, 2, 73, 81, 3, 2, 2, 2, 74, 75, 5, 36, 19, 2, 75,
	76, 8, 4, 1, 2, 76, 81, 3, 2, 2, 2, 77, 78, 5, 8, 5, 2, 78, 79, 8, 4, 1,
	2, 79, 81, 3, 2, 2, 2, 80, 59, 3, 2, 2, 2, 80, 62, 3, 2, 2, 2, 80, 65,
	3, 2, 2, 2, 80, 68, 3, 2, 2, 2, 80, 71, 3, 2, 2, 2, 80, 74, 3, 2, 2, 2,
	80, 77, 3, 2, 2, 2, 81, 7, 3, 2, 2, 2, 82, 83, 7, 71, 2, 2, 83, 84, 7,
	16, 2, 2, 84, 85, 8, 5, 1, 2, 85, 9, 3, 2, 2, 2, 86, 87, 7, 45, 2, 2, 87,
	88, 7, 46, 2, 2, 88, 89, 7, 75, 2, 2, 89, 90, 7, 14, 2, 2, 90, 91, 5, 12,
	7, 2, 91, 92, 7, 22, 2, 2, 92, 93, 5, 18, 10, 2, 93, 94, 7, 16, 2, 2, 94,
	95, 8, 6, 1, 2, 95, 121, 3, 2, 2, 2, 96, 97, 7, 45, 2, 2, 97, 98, 7, 46,
	2, 2, 98, 99, 7, 75, 2, 2, 99, 100, 7, 22, 2, 2, 100, 101, 5, 18, 10, 2,
	101, 102, 7, 16, 2, 2, 102, 103, 8, 6, 1, 2, 103, 121, 3, 2, 2, 2, 104,
	105, 7, 45, 2, 2, 105, 106, 7, 75, 2, 2, 106, 107, 7, 14, 2, 2, 107, 108,
	5, 12, 7, 2, 108, 109, 7, 22, 2, 2, 109, 110, 5, 18, 10, 2, 110, 111, 7,
	16, 2, 2, 111, 112, 8, 6, 1, 2, 112, 121, 3, 2, 2, 2, 113, 114, 7, 45,
	2, 2, 114, 115, 7, 75, 2, 2, 115, 116, 7, 22, 2, 2, 116, 117, 5, 18, 10,
	2, 117, 118, 7, 16, 2, 2, 118, 119, 8, 6, 1, 2, 119, 121, 3, 2, 2, 2, 120,
	86, 3, 2, 2, 2, 120, 96, 3, 2, 2, 2, 120, 104, 3, 2, 2, 2, 120, 113, 3,
	2, 2, 2, 121, 11, 3, 2, 2, 2, 122, 123, 7, 35, 2, 2, 123, 135, 8, 7, 1,
	2, 124, 125, 7, 36, 2, 2, 125, 135, 8, 7, 1, 2, 126, 127, 7, 43, 2, 2,
	127, 135, 8, 7, 1, 2, 128, 129, 7, 40, 2, 2, 129, 135, 8, 7, 1, 2, 130,
	131, 7, 41, 2, 2, 131, 135, 8, 7, 1, 2, 132, 133, 7, 44, 2, 2, 133, 135,
	8, 7, 1, 2, 134, 122, 3, 2, 2, 2, 134, 124, 3, 2, 2, 2, 134, 126, 3, 2,
	2, 2, 134, 128, 3, 2, 2, 2, 134, 130, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2,
	135, 13, 3, 2, 2, 2, 136, 137, 7, 75, 2, 2, 137, 138, 7, 22, 2, 2, 138,
	139, 5, 18, 10, 2, 139, 140, 7, 16, 2, 2, 140, 141, 8, 8, 1, 2, 141, 15,
	3, 2, 2, 2, 142, 143, 7, 73, 2, 2, 143, 159, 8, 9, 1, 2, 144, 145, 7, 74,
	2, 2, 145, 159, 8, 9, 1, 2, 146, 147, 7, 77, 2, 2, 147, 159, 8, 9, 1, 2,
	148, 149, 7, 78, 2, 2, 149, 159, 8, 9, 1, 2, 150, 151, 7, 42, 2, 2, 151,
	159, 8, 9, 1, 2, 152, 153, 7, 50, 2, 2, 153, 159, 8, 9, 1, 2, 154, 155,
	7, 51, 2, 2, 155, 159, 8, 9, 1, 2, 156, 157, 7, 75, 2, 2, 157, 159, 8,
	9, 1, 2, 158, 142, 3, 2, 2, 2, 158, 144, 3, 2, 2, 2, 158, 146, 3, 2, 2,
	2, 158, 148, 3, 2, 2, 2, 158, 150, 3, 2, 2, 2, 158, 152, 3, 2, 2, 2, 158,
	154, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 17, 3, 2, 2, 2, 160, 161, 8,
	10, 1, 2, 161, 162, 7, 24, 2, 2, 162, 163, 5, 18, 10, 20, 163, 164, 8,
	10, 1, 2, 164, 202, 3, 2, 2, 2, 165, 166, 7, 37, 2, 2, 166, 167, 7, 20,
	2, 2, 167, 168, 5, 18, 10, 2, 168, 169, 7, 15, 2, 2, 169, 170, 5, 18, 10,
	2, 170, 171, 7, 21, 2, 2, 171, 172, 8, 10, 1, 2, 172, 202, 3, 2, 2, 2,
	173, 174, 7, 39, 2, 2, 174, 175, 7, 20, 2, 2, 175, 176, 5, 18, 10, 2, 176,
	177, 7, 15, 2, 2, 177, 178, 5, 18, 10, 2, 178, 179, 7, 21, 2, 2, 179, 180,
	8, 10, 1, 2, 180, 202, 3, 2, 2, 2, 181, 182, 7, 30, 2, 2, 182, 183, 5,
	18, 10, 7, 183, 184, 8, 10, 1, 2, 184, 202, 3, 2, 2, 2, 185, 186, 7, 20,
	2, 2, 186, 187, 5, 18, 10, 2, 187, 188, 7, 21, 2, 2, 188, 189, 8, 10, 1,
	2, 189, 202, 3, 2, 2, 2, 190, 191, 5, 16, 9, 2, 191, 192, 7, 48, 2, 2,
	192, 193, 7, 35, 2, 2, 193, 202, 3, 2, 2, 2, 194, 195, 5, 16, 9, 2, 195,
	196, 7, 48, 2, 2, 196, 197, 7, 36, 2, 2, 197, 202, 3, 2, 2, 2, 198, 199,
	5, 16, 9, 2, 199, 200, 8, 10, 1, 2, 200, 202, 3, 2, 2, 2, 201, 160, 3,
	2, 2, 2, 201, 165, 3, 2, 2, 2, 201, 173, 3, 2, 2, 2, 201, 181, 3, 2, 2,
	2, 201, 185, 3, 2, 2, 2, 201, 190, 3, 2, 2, 2, 201, 194, 3, 2, 2, 2, 201,
	198, 3, 2, 2, 2, 202, 255, 3, 2, 2, 2, 203, 204, 12, 19, 2, 2, 204, 205,
	9, 2, 2, 2, 205, 206, 5, 18, 10, 20, 206, 207, 8, 10, 1, 2, 207, 254, 3,
	2, 2, 2, 208, 209, 12, 18, 2, 2, 209, 210, 9, 3, 2, 2, 210, 211, 5, 18,
	10, 19, 211, 212, 8, 10, 1, 2, 212, 254, 3, 2, 2, 2, 213, 214, 12, 15,
	2, 2, 214, 215, 7, 17, 2, 2, 215, 216, 5, 18, 10, 16, 216, 217, 8, 10,
	1, 2, 217, 254, 3, 2, 2, 2, 218, 219, 12, 14, 2, 2, 219, 220, 7, 18, 2,
	2, 220, 221, 5, 18, 10, 15, 221, 222, 8, 10, 1, 2, 222, 254, 3, 2, 2, 2,
	223, 224, 12, 13, 2, 2, 224, 225, 7, 8, 2, 2, 225, 226, 5, 18, 10, 14,
	226, 227, 8, 10, 1, 2, 227, 254, 3, 2, 2, 2, 228, 229, 12, 12, 2, 2, 229,
	230, 7, 9, 2, 2, 230, 231, 5, 18, 10, 13, 231, 232, 8, 10, 1, 2, 232, 254,
	3, 2, 2, 2, 233, 234, 12, 11, 2, 2, 234, 235, 7, 6, 2, 2, 235, 236, 5,
	18, 10, 12, 236, 237, 8, 10, 1, 2, 237, 254, 3, 2, 2, 2, 238, 239, 12,
	10, 2, 2, 239, 240, 7, 7, 2, 2, 240, 241, 5, 18, 10, 11, 241, 242, 8, 10,
	1, 2, 242, 254, 3, 2, 2, 2, 243, 244, 12, 9, 2, 2, 244, 245, 7, 4, 2, 2,
	245, 246, 5, 18, 10, 10, 246, 247, 8, 10, 1, 2, 247, 254, 3, 2, 2, 2, 248,
	249, 12, 8, 2, 2, 249, 250, 7, 5, 2, 2, 250, 251, 5, 18, 10, 9, 251, 252,
	8, 10, 1, 2, 252, 254, 3, 2, 2, 2, 253, 203, 3, 2, 2, 2, 253, 208, 3, 2,
	2, 2, 253, 213, 3, 2, 2, 2, 253, 218, 3, 2, 2, 2, 253, 223, 3, 2, 2, 2,
	253, 228, 3, 2, 2, 2, 253, 233, 3, 2, 2, 2, 253, 238, 3, 2, 2, 2, 253,
	243, 3, 2, 2, 2, 253, 248, 3, 2, 2, 2, 254, 257, 3, 2, 2, 2, 255, 253,
	3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 19, 3, 2, 2, 2, 257, 255, 3, 2,
	2, 2, 258, 259, 7, 49, 2, 2, 259, 260, 7, 20, 2, 2, 260, 261, 5, 18, 10,
	2, 261, 262, 7, 21, 2, 2, 262, 263, 7, 16, 2, 2, 263, 264, 8, 11, 1, 2,
	264, 273, 3, 2, 2, 2, 265, 266, 7, 49, 2, 2, 266, 267, 7, 20, 2, 2, 267,
	268, 5, 18, 10, 2, 268, 269, 5, 22, 12, 2, 269, 270, 7, 16, 2, 2, 270,
	271, 8, 11, 1, 2, 271, 273, 3, 2, 2, 2, 272, 258, 3, 2, 2, 2, 272, 265,
	3, 2, 2, 2, 273, 21, 3, 2, 2, 2, 274, 276, 5, 24, 13, 2, 275, 274, 3, 2,
	2, 2, 276, 279, 3, 2, 2, 2, 277, 275, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2,
	278, 280, 3, 2, 2, 2, 279, 277, 3, 2, 2, 2, 280, 281, 8, 12, 1, 2, 281,
	23, 3, 2, 2, 2, 282, 283, 7, 15, 2, 2, 283, 284, 5, 18, 10, 2, 284, 285,
	8, 13, 1, 2, 285, 25, 3, 2, 2, 2, 286, 287, 5, 28, 15, 2, 287, 288, 8,
	14, 1, 2, 288, 27, 3, 2, 2, 2, 289, 290, 7, 67, 2, 2, 290, 291, 5, 18,
	10, 2, 291, 292, 5, 34, 18, 2, 292, 293, 8, 15, 1, 2, 293, 316, 3, 2, 2,
	2, 294, 295, 7, 67, 2, 2, 295, 296, 5, 18, 10, 2, 296, 297, 5, 34, 18,
	2, 297, 298, 7, 69, 2, 2, 298, 299, 5, 34, 18, 2, 299, 300, 8, 15, 1, 2,
	300, 316, 3, 2, 2, 2, 301, 302, 7, 67, 2, 2, 302, 303, 5, 18, 10, 2, 303,
	304, 5, 34, 18, 2, 304, 305, 5, 32, 17, 2, 305, 306, 7, 69, 2, 2, 306,
	307, 5, 34, 18, 2, 307, 308, 8, 15, 1, 2, 308, 316, 3, 2, 2, 2, 309, 310,
	7, 67, 2, 2, 310, 311, 5, 18, 10, 2, 311, 312, 5, 34, 18, 2, 312, 313,
	5, 32, 17, 2, 313, 314, 8, 15, 1, 2, 314, 316, 3, 2, 2, 2, 315, 289, 3,
	2, 2, 2, 315, 294, 3, 2, 2, 2, 315, 301, 3, 2, 2, 2, 315, 309, 3, 2, 2,
	2, 316, 29, 3, 2, 2, 2, 317, 318, 7, 68, 2, 2, 318, 319, 5, 18, 10, 2,
	319, 320, 5, 34, 18, 2, 320, 321, 8, 16, 1, 2, 321, 31, 3, 2, 2, 2, 322,
	324, 5, 30, 16, 2, 323, 322, 3, 2, 2, 2, 324, 327, 3, 2, 2, 2, 325, 323,
	3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326, 328, 3, 2, 2, 2, 327, 325, 3, 2,
	2, 2, 328, 329, 8, 17, 1, 2, 329, 33, 3, 2, 2, 2, 330, 331, 7, 10, 2, 2,
	331, 332, 5, 4, 3, 2, 332, 333, 7, 11, 2, 2, 333, 334, 8, 18, 1, 2, 334,
	339, 3, 2, 2, 2, 335, 336, 7, 10, 2, 2, 336, 337, 7, 11, 2, 2, 337, 339,
	8, 18, 1, 2, 338, 330, 3, 2, 2, 2, 338, 335, 3, 2, 2, 2, 339, 35, 3, 2,
	2, 2, 340, 341, 5, 38, 20, 2, 341, 342, 8, 19, 1, 2, 342, 347, 3, 2, 2,
	2, 343, 344, 5, 40, 21, 2, 344, 345, 8, 19, 1, 2, 345, 347, 3, 2, 2, 2,
	346, 340, 3, 2, 2, 2, 346, 343, 3, 2, 2, 2, 347, 37, 3, 2, 2, 2, 348, 349,
	7, 70, 2, 2, 349, 350, 5, 18, 10, 2, 350, 351, 5, 34, 18, 2, 351, 352,
	8, 20, 1, 2, 352, 39, 3, 2, 2, 2, 353, 354, 7, 72, 2, 2, 354, 355, 5, 34,
	18, 2, 355, 356, 8, 21, 1, 2, 356, 41, 3, 2, 2, 2, 16, 54, 80, 120, 134,
	158, 201, 253, 255, 272, 277, 315, 325, 338, 346,
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
	"start", "instrucciones", "instruccion", "control", "declaracion", "tipovariable",
	"identificadores", "valores", "expresion", "impresion", "impresionexpresion",
	"expcoma", "condicionales", "funcionif", "funcionelseif", "listaelseif",
	"bloque", "bucles", "fwhile", "floop",
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
	gramaticaParserRULE_instrucciones      = 1
	gramaticaParserRULE_instruccion        = 2
	gramaticaParserRULE_control            = 3
	gramaticaParserRULE_declaracion        = 4
	gramaticaParserRULE_tipovariable       = 5
	gramaticaParserRULE_identificadores    = 6
	gramaticaParserRULE_valores            = 7
	gramaticaParserRULE_expresion          = 8
	gramaticaParserRULE_impresion          = 9
	gramaticaParserRULE_impresionexpresion = 10
	gramaticaParserRULE_expcoma            = 11
	gramaticaParserRULE_condicionales      = 12
	gramaticaParserRULE_funcionif          = 13
	gramaticaParserRULE_funcionelseif      = 14
	gramaticaParserRULE_listaelseif        = 15
	gramaticaParserRULE_bloque             = 16
	gramaticaParserRULE_bucles             = 17
	gramaticaParserRULE_fwhile             = 18
	gramaticaParserRULE_floop              = 19
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
		p.SetState(40)
		p.Match(gramaticaParserTKR_fn)
	}
	{
		p.SetState(41)
		p.Match(gramaticaParserTKR_main)
	}
	{
		p.SetState(42)
		p.Match(gramaticaParserTK_par_apertura)
	}
	{
		p.SetState(43)
		p.Match(gramaticaParserTK_par_cierre)
	}
	{
		p.SetState(44)
		p.Match(gramaticaParserTK_corchete_apertura)
	}
	{
		p.SetState(45)

		var _x = p.Instrucciones()

		localctx.(*StartContext)._instrucciones = _x
	}
	{
		p.SetState(46)
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
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-18)&-(0x1f+1)) == 0 && ((1<<uint((_la-18)))&((1<<(gramaticaParserTK_par_apertura-18))|(1<<(gramaticaParserTK_menos-18))|(1<<(gramaticaParserTK_sig_admiracion-18))|(1<<(gramaticaParserTKR_pow-18))|(1<<(gramaticaParserTKR_powf-18))|(1<<(gramaticaParserTKR_Str-18))|(1<<(gramaticaParserTKR_let-18))|(1<<(gramaticaParserTKR_println-18))|(1<<(gramaticaParserTKR_true-18))|(1<<(gramaticaParserTKR_false-18)))) != 0) || (((_la-65)&-(0x1f+1)) == 0 && ((1<<uint((_la-65)))&((1<<(gramaticaParserTKR_if-65))|(1<<(gramaticaParserTKR_while-65))|(1<<(gramaticaParserTKR_break-65))|(1<<(gramaticaParserTKR_loop-65))|(1<<(gramaticaParserTK_entero-65))|(1<<(gramaticaParserTK_decimal-65))|(1<<(gramaticaParserTK_id-65))|(1<<(gramaticaParserTK_cadena-65))|(1<<(gramaticaParserTK_caracter-65)))) != 0) {
		{
			p.SetState(49)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(54)
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

func (s *InstruccionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *InstruccionContext) Set_impresion(v IImpresionContext) { s._impresion = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_identificadores(v IIdentificadoresContext) { s._identificadores = v }

func (s *InstruccionContext) Set_condicionales(v ICondicionalesContext) { s._condicionales = v }

func (s *InstruccionContext) Set_bucles(v IBuclesContext) { s._bucles = v }

func (s *InstruccionContext) Set_control(v IControlContext) { s._control = v }

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

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)

			var _x = p.expresion(0)

			localctx.(*InstruccionContext)._expresion = _x
		}
		fmt.Println("mensaje en instrucciones: ", localctx.(*InstruccionContext).Get_expresion().GetExp())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)

			var _x = p.Impresion()

			localctx.(*InstruccionContext)._impresion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_impresion().GetInst()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_declaracion().GetInst()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(66)

			var _x = p.Identificadores()

			localctx.(*InstruccionContext)._identificadores = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_identificadores().GetInst()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(69)

			var _x = p.Condicionales()

			localctx.(*InstruccionContext)._condicionales = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_condicionales().GetInst()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(72)

			var _x = p.Bucles()

			localctx.(*InstruccionContext)._bucles = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_bucles().GetInst()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(75)

			var _x = p.Control()

			localctx.(*InstruccionContext)._control = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_control().GetInst()

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
	p.EnterRule(localctx, 6, gramaticaParserRULE_control)

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
		p.SetState(80)
		p.Match(gramaticaParserTKR_break)
	}
	{
		p.SetState(81)
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
	p.EnterRule(localctx, 8, gramaticaParserRULE_declaracion)

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

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(85)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(86)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(87)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(88)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(89)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(90)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(91)
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
			p.SetState(94)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(95)
			p.Match(gramaticaParserTKR_mut)
		}
		{
			p.SetState(96)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(97)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(98)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(99)
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
			p.SetState(102)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(103)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(104)
			p.Match(gramaticaParserTK_dosPuntos)
		}
		{
			p.SetState(105)

			var _x = p.Tipovariable()

			localctx.(*DeclaracionContext)._tipovariable = _x
		}
		{
			p.SetState(106)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(107)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(108)
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
			p.SetState(111)
			p.Match(gramaticaParserTKR_let)
		}
		{
			p.SetState(112)

			var _m = p.Match(gramaticaParserTK_id)

			localctx.(*DeclaracionContext).idd = _m
		}
		{
			p.SetState(113)
			p.Match(gramaticaParserTK_igual)
		}
		{
			p.SetState(114)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}
		{
			p.SetState(115)
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
	p.EnterRule(localctx, 10, gramaticaParserRULE_tipovariable)

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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTKR_numericos_enteros:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(gramaticaParserTKR_numericos_enteros)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.INTEGER

	case gramaticaParserTKR_numericos_flotantes:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(gramaticaParserTKR_numericos_flotantes)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.FLOAT

	case gramaticaParserTKR_String:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Match(gramaticaParserTKR_String)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.STRING

	case gramaticaParserTKR_bool:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)
			p.Match(gramaticaParserTKR_bool)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.BOOLEAN

	case gramaticaParserTKR_char:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)
			p.Match(gramaticaParserTKR_char)
		}
		localctx.(*TipovariableContext).tipovar = Interfaces.CHAR

	case gramaticaParserTKR_usize:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(130)
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
	p.EnterRule(localctx, 12, gramaticaParserRULE_identificadores)

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
		p.SetState(134)

		var _m = p.Match(gramaticaParserTK_id)

		localctx.(*IdentificadoresContext).e1 = _m
	}
	{
		p.SetState(135)
		p.Match(gramaticaParserTK_igual)
	}
	{
		p.SetState(136)

		var _x = p.expresion(0)

		localctx.(*IdentificadoresContext).e2 = _x
	}
	{
		p.SetState(137)
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
	p.EnterRule(localctx, 14, gramaticaParserRULE_valores)

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

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTK_entero:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)

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
			p.SetState(142)

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
			p.SetState(144)

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
			p.SetState(146)

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
			p.SetState(148)

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
			p.SetState(150)
			p.Match(gramaticaParserTKR_true)
		}
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(true, Interfaces.BOOLEAN)

	case gramaticaParserTKR_false:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(152)
			p.Match(gramaticaParserTKR_false)
		}
		localctx.(*ValoresContext).exp = Expresion.NuevoPrimitivo(false, Interfaces.BOOLEAN)

	case gramaticaParserTK_id:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(154)

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

	// GetVa returns the va rule contexts.
	GetVa() IExpresionContext

	// GetVall returns the vall rule contexts.
	GetVall() IValoresContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// SetVa sets the va rule contexts.
	SetVa(IExpresionContext)

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
	va     IExpresionContext
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

func (s *ExpresionContext) GetVa() IExpresionContext { return s.va }

func (s *ExpresionContext) GetVall() IValoresContext { return s.vall }

func (s *ExpresionContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *ExpresionContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *ExpresionContext) SetVa(v IExpresionContext) { s.va = v }

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
	_startState := 16
	p.EnterRecursionRule(localctx, 16, gramaticaParserRULE_expresion, _p)
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
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(159)
			p.Match(gramaticaParserTK_menos)
		}
		{
			p.SetState(160)

			var _x = p.expresion(18)

			localctx.(*ExpresionContext).e1 = _x
		}
		numero := -1
		numm := Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), numm, Interfaces.MULTIPLICACION)

	case 2:
		{
			p.SetState(163)
			p.Match(gramaticaParserTKR_pow)
		}
		{
			p.SetState(164)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(165)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(166)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(167)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e2 = _x
		}
		{
			p.SetState(168)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.POW)

	case 3:
		{
			p.SetState(171)
			p.Match(gramaticaParserTKR_powf)
		}
		{
			p.SetState(172)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(173)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e1 = _x
		}
		{
			p.SetState(174)
			p.Match(gramaticaParserTK_coma)
		}
		{
			p.SetState(175)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e2 = _x
		}
		{
			p.SetState(176)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaAritmetica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.POWF)

	case 4:
		{
			p.SetState(179)
			p.Match(gramaticaParserTK_sig_admiracion)
		}
		{
			p.SetState(180)

			var _x = p.expresion(5)

			localctx.(*ExpresionContext).e1 = _x
		}
		localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE1().GetExp(), Interfaces.NOT)

	case 5:
		{
			p.SetState(183)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(184)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).va = _x
		}
		{
			p.SetState(185)
			p.Match(gramaticaParserTK_par_cierre)
		}
		localctx.(*ExpresionContext).exp = localctx.(*ExpresionContext).GetVa().GetExp()

	case 6:
		{
			p.SetState(188)
			p.Valores()
		}
		{
			p.SetState(189)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(190)
			p.Match(gramaticaParserTKR_numericos_enteros)
		}

	case 7:
		{
			p.SetState(192)
			p.Valores()
		}
		{
			p.SetState(193)
			p.Match(gramaticaParserTKR_as)
		}
		{
			p.SetState(194)
			p.Match(gramaticaParserTKR_numericos_flotantes)
		}

	case 8:
		{
			p.SetState(196)

			var _x = p.Valores()

			localctx.(*ExpresionContext).vall = _x
		}
		localctx.(*ExpresionContext).exp = localctx.(*ExpresionContext).GetVall().GetExp()
		fmt.Println(localctx.(*ExpresionContext).exp)

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(251)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(201)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(202)

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
					p.SetState(203)

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
				p.SetState(206)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(207)

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
					p.SetState(208)

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
				p.SetState(211)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(212)
					p.Match(gramaticaParserTK_menor)
				}
				{
					p.SetState(213)

					var _x = p.expresion(14)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.MENOR_QUE)

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(216)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(217)
					p.Match(gramaticaParserTK_mayor)
				}
				{
					p.SetState(218)

					var _x = p.expresion(13)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.MAYOR_QUE)

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(222)
					p.Match(gramaticaParserTK_mayor_igual)
				}
				{
					p.SetState(223)

					var _x = p.expresion(12)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.MAYOR_IGUAL)

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(226)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(227)
					p.Match(gramaticaParserTK_menor_igual)
				}
				{
					p.SetState(228)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.MENOR_IGUAL)

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(232)
					p.Match(gramaticaParserTK_igualacion)
				}
				{
					p.SetState(233)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.IGUALDAD)

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(236)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(237)
					p.Match(gramaticaParserTK_diferente)
				}
				{
					p.SetState(238)

					var _x = p.expresion(9)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaRelacional(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.DESIGUAL)

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(241)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(242)
					p.Match(gramaticaParserTK_or)
				}
				{
					p.SetState(243)

					var _x = p.expresion(8)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.OR)

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(247)
					p.Match(gramaticaParserTK_and)
				}
				{
					p.SetState(248)

					var _x = p.expresion(7)

					localctx.(*ExpresionContext).e2 = _x
				}
				localctx.(*ExpresionContext).exp = Expresiones.NuevaLogica(localctx.(*ExpresionContext).GetE1().GetExp(), localctx.(*ExpresionContext).GetE2().GetExp(), Interfaces.AND)

			}

		}
		p.SetState(255)
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
	p.EnterRule(localctx, 18, gramaticaParserRULE_impresion)

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

	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(257)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(258)

			var _x = p.expresion(0)

			localctx.(*ImpresionContext).e1 = _x
		}
		{
			p.SetState(259)
			p.Match(gramaticaParserTK_par_cierre)
		}
		{
			p.SetState(260)
			p.Match(gramaticaParserTK_pcoma)
		}
		localctx.(*ImpresionContext).inst = Instruccion.NuevoPrint(localctx.(*ImpresionContext).GetE1().GetExp(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
			p.Match(gramaticaParserTKR_println)
		}
		{
			p.SetState(264)
			p.Match(gramaticaParserTK_par_apertura)
		}
		{
			p.SetState(265)

			var _x = p.expresion(0)

			localctx.(*ImpresionContext).e2 = _x
		}
		{
			p.SetState(266)

			var _x = p.Impresionexpresion()

			localctx.(*ImpresionContext).li = _x
		}
		{
			p.SetState(267)
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
	p.EnterRule(localctx, 20, gramaticaParserRULE_impresionexpresion)
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
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserTK_coma {
		{
			p.SetState(272)

			var _x = p.Expcoma()

			localctx.(*ImpresionexpresionContext)._expcoma = _x
		}
		localctx.(*ImpresionexpresionContext).list = append(localctx.(*ImpresionexpresionContext).list, localctx.(*ImpresionexpresionContext)._expcoma)

		p.SetState(277)
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
	p.EnterRule(localctx, 22, gramaticaParserRULE_expcoma)

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
		p.SetState(280)
		p.Match(gramaticaParserTK_coma)
	}
	{
		p.SetState(281)

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
	p.EnterRule(localctx, 24, gramaticaParserRULE_condicionales)

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
		p.SetState(284)

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
	p.EnterRule(localctx, 26, gramaticaParserRULE_funcionif)

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

	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(288)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(289)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).ee = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetEe().GetLista(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(293)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(294)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).e5 = _x
		}
		{
			p.SetState(295)
			p.Match(gramaticaParserTKR_else)
		}
		{
			p.SetState(296)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b2 = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetE5().GetLista(), nil, localctx.(*FuncionifContext).GetB2().GetLista())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(299)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(300)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(301)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b1 = _x
		}
		{
			p.SetState(302)

			var _x = p.Listaelseif()

			localctx.(*FuncionifContext)._listaelseif = _x
		}
		{
			p.SetState(303)
			p.Match(gramaticaParserTKR_else)
		}
		{
			p.SetState(304)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b2 = _x
		}
		localctx.(*FuncionifContext).inst = Instruccion.NewIf(localctx.(*FuncionifContext).GetE1().GetExp(), localctx.(*FuncionifContext).GetB1().GetLista(), localctx.(*FuncionifContext).Get_listaelseif().GetLista(), localctx.(*FuncionifContext).GetB2().GetLista())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(307)
			p.Match(gramaticaParserTKR_if)
		}
		{
			p.SetState(308)

			var _x = p.expresion(0)

			localctx.(*FuncionifContext).e1 = _x
		}
		{
			p.SetState(309)

			var _x = p.Bloque()

			localctx.(*FuncionifContext).b1 = _x
		}
		{
			p.SetState(310)

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
	p.EnterRule(localctx, 28, gramaticaParserRULE_funcionelseif)

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
		p.SetState(315)
		p.Match(gramaticaParserTKR_elseif)
	}
	{
		p.SetState(316)

		var _x = p.expresion(0)

		localctx.(*FuncionelseifContext).e1 = _x
	}
	{
		p.SetState(317)

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
	p.EnterRule(localctx, 30, gramaticaParserRULE_listaelseif)
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
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserTKR_elseif {
		{
			p.SetState(320)

			var _x = p.Funcionelseif()

			localctx.(*ListaelseifContext)._funcionelseif = _x
		}
		localctx.(*ListaelseifContext).list = append(localctx.(*ListaelseifContext).list, localctx.(*ListaelseifContext)._funcionelseif)

		p.SetState(325)
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
	p.EnterRule(localctx, 32, gramaticaParserRULE_bloque)

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

	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(328)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(329)

			var _x = p.Instrucciones()

			localctx.(*BloqueContext)._instrucciones = _x
		}
		{
			p.SetState(330)
			p.Match(gramaticaParserTK_corchete_cierre)
		}
		localctx.(*BloqueContext).lista = localctx.(*BloqueContext).Get_instrucciones().GetLista()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(333)
			p.Match(gramaticaParserTK_corchete_apertura)
		}
		{
			p.SetState(334)
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
	p.EnterRule(localctx, 34, gramaticaParserRULE_bucles)

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

	p.SetState(344)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserTKR_while:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)

			var _x = p.Fwhile()

			localctx.(*BuclesContext)._fwhile = _x
		}
		localctx.(*BuclesContext).inst = localctx.(*BuclesContext).Get_fwhile().GetInst()

	case gramaticaParserTKR_loop:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(341)

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
	p.EnterRule(localctx, 36, gramaticaParserRULE_fwhile)

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
		p.SetState(346)
		p.Match(gramaticaParserTKR_while)
	}
	{
		p.SetState(347)

		var _x = p.expresion(0)

		localctx.(*FwhileContext).e1 = _x
	}
	{
		p.SetState(348)

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
	p.EnterRule(localctx, 38, gramaticaParserRULE_floop)

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
		p.SetState(351)
		p.Match(gramaticaParserTKR_loop)
	}
	{
		p.SetState(352)

		var _x = p.Bloque()

		localctx.(*FloopContext).bb = _x
	}
	localctx.(*FloopContext).inst = Instruccion.NewLoop(localctx.(*FloopContext).GetBb().GetLista())

	return localctx
}

func (p *gramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

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
