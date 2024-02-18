package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func layoutGridWarp() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	text1 := canvas.NewText("1", blue)
	text2 := canvas.NewText("2", blue)
	text3 := canvas.NewText("3", blue)
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
		text1, text2, text3)
	myWindow.SetContent(grid)

	// myWindow.Resize(fyne.NewSize(180, 75))
	myWindow.ShowAndRun()
}
