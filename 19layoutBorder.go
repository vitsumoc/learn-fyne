package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func layoutBorder() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	top := canvas.NewText("top bar", blue)
	left := canvas.NewText("left", blue)
	middle := canvas.NewText("content", blue)
	content := container.NewBorder(top, nil, left, nil, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
