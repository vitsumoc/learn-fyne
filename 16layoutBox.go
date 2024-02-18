package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func layoutBox() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	// 水平布局 左 中 间隔 右
	// 其中间隔用来适配拉伸
	text1 := canvas.NewText("Hello", blue)
	// text2 := canvas.NewText("Thereeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeea", blue)
	var tableData = [][]string{{"Thereeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeea"}}
	text2 := widget.NewTable(
		func() (int, int) {
			return len(tableData), len(tableData[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tableData[i.Row][i.Col])
		},
	)
	text3 := canvas.NewText("(right)", blue)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	// 上下布局
	text4 := canvas.NewText("centered", blue)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	myWindow.ShowAndRun()
}
