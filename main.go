package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	//"fyne.io/fyne/v2/widget"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Print("Hola mundo")
	}
	a := app.New()
	w := a.NewWindow("Hola mundo")
	rect := canvas.NewRectangle(color.Black)
	w.SetContent(rect)

	w.ShowAndRun()
}
