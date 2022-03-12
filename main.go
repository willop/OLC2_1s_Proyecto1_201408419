package main

import (
	//"log"
	"fmt"
	"image/color"
	"proyecto1/Estructura"
	"proyecto1/Interfaces"
	"proyecto1/Utilitario"
	"proyecto1/parser"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

//incializanod
type gramaticaListener struct {
	*parser.BasegramaticaListener
}

type TreeShapeListener struct {
	*parser.BasegramaticaListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	result := ctx.GetListainstrucciones() //esto me retorna nul por que no existe ninguna instruccion
	var recolector Utilitario.Recolector = Utilitario.NuevoRecolector()
	var globalEnv Estructura.Entorno = Estructura.NuevoEntorno(nil, "global", 0)
	fmt.Println("Nuevo entorno", globalEnv)
	fmt.Println("Get result", result)
	//falta implementacion de listas de listas
	for _, s := range result.ToArray() {
		s.(Interfaces.Instruccion).Ejecutar(globalEnv, &recolector)
	}
	fmt.Println("/n/n******************************RECOLECTOR****************************/n")
	//fmt.Println(recolector.Consolavirtual)
	if recolector.Consolavirtual.Len() > 0 {
		for _, s := range recolector.Consolavirtual.ToArray() {
			fmt.Println(s)
		}
	}
	fmt.Println("/n/n*************************** FIN RECOLECTOR **************************/n/n/n")

	fmt.Println("/n/n****************************** ERRORES ******************************/n")
	//fmt.Println(recolector.Consolavirtual)
	if recolector.ListaErrores.Len() > 0 {
		for _, s := range recolector.ListaErrores.ToArray() {
			fmt.Println(s)
		}
	}
	fmt.Println("/n/n*************************** FIN ERRORES **************************/n")

}

func main() {

	a := app.New()
	w := a.NewWindow("Proyecto 1 - OLC2")
	temp := Interfaces.ConstructorSimbolo("AA", "", "", 2)
	fmt.Println("la importacion ", temp.GetID())
	entradacodigo := widget.NewMultiLineEntry()
	entradacodigo.Resize(fyne.NewSize(500, 780))
	entradacodigo.Move(fyne.NewPos(5, 5))
	salidacodigo := widget.NewMultiLineEntry()
	salidacodigo.Resize(fyne.NewSize(500, 780))
	salidacodigo.Move(fyne.NewPos(650, 5))
	boton := widget.NewButtonWithIcon("Ejecutar", theme.HomeIcon(), func() {
		fmt.Println("*********************** Compilando *******************")
		salidacodigo.SetText(entradacodigo.Text)
		is := antlr.NewInputStream(entradacodigo.Text)
		lexer := parser.NewgramaticaLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		//incluyendo el parser
		p := parser.NewgramaticaParser(stream)
		//aca inicial el nuevo programa
		//crear un enorno al iniciar
		//recorrer la lista de instrucciones con el metodo de ejecutar
		p.BuildParseTrees = true
		tree := p.Start()
		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
		//antlr.ParseTreeWalkerDefault.Walk(gramaticaListener{}, p.Start())
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
