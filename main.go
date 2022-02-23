package main

import (
	//"log"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"main.go/parser"
)

func main() {
	a := app.New()
	w := a.NewWindow("Proyecto 1 - OLC2")

	entradacodigo := widget.NewMultiLineEntry()
	entradacodigo.Resize(fyne.NewSize(500, 780))
	entradacodigo.Move(fyne.NewPos(5, 5))
	salidacodigo := widget.NewMultiLineEntry()
	salidacodigo.Resize(fyne.NewSize(500, 780))
	salidacodigo.Move(fyne.NewPos(650, 5))
	boton := widget.NewButtonWithIcon("Ejecutar", theme.HomeIcon(), func() {
		salidacodigo.SetText(entradacodigo.Text)
		is := antlr.NewInputStream(entradacodigo.Text)
		lexer := parser.NewgramaticaLexer(is)
		for {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
			fmt.Printf("%s (%q)\n",
				lexer.SymbolicNames[t.GetTokenType()], t.GetText())

		}
	})
	boton.Move(fyne.NewPos(530, 390))
	boton.Resize(fyne.NewSize(100, 30))
	boton2 := widget.NewButtonWithIcon("Reportes", theme.ComputerIcon(), func() {
		salidacodigo.SetText(entradacodigo.Text)
	})
	boton2.Move(fyne.NewPos(530, 425))
	boton2.Resize(fyne.NewSize(100, 30))
	//textArea := widget.NewMultiLineEntry()
	entradacodigo.SetPlaceHolder("Ingrese su codigo aca")
	salidacodigo.SetPlaceHolder("Salida del programa")
	//fondo.Resize(fyne.NewSize(1200, 800))
	gradiente := canvas.NewHorizontalGradient(color.Black, color.NRGBA{R: 81, G: 81, B: 81, A: 255})
	gradiente.Resize(fyne.NewSize(1200, 800))
	content := container.NewWithoutLayout(gradiente, entradacodigo, salidacodigo, boton, boton2)

	// we can also append items
	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 800))
	w.ShowAndRun()
}
