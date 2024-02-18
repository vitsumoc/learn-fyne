package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var topLevelLayOutContent *fyne.Container

func myTest() {
	a := app.New()
	win := a.NewWindow("Server Mon")
	c1 := canvas.NewText("Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1Canvas Object 1", color.Black)
	c2 := canvas.NewText("Canvas Object 2", color.Black)
	n1 := canvas.NewText("Canvas Object 3", color.Black)
	n2 := canvas.NewText("Canvas Object 4", color.Black)

	topLevelLayOutContent = container.New(layout.NewVBoxLayout(),
		container.New(layout.NewStackLayout(),
			container.New(layout.NewGridLayoutWithColumns(2), c1, c2)),
		container.New(layout.NewStackLayout(),
			container.New(layout.NewGridLayoutWithColumns(2), n1, n2)),
	)

	win.SetContent(topLevelLayOutContent)
	win.Resize(fyne.NewSize(float32(400), float32(30)))
	win.ShowAndRun()

}
