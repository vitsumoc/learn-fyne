package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func fWidgets() {
	myApp := app.New()
	myWindow := myApp.NewWindow("fWidgets")

	// 文本标签
	label := widget.NewLabel("text")
	// 按钮
	buttion := widget.NewButton("click me", func() {
		log.Println("tapped")
	})
	// 输入框
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	// 勾选
	check := widget.NewCheck("Optional", func(value bool) {
		log.Println("Check set to", value)
	})
	// 单选
	radio := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Radio set to", value)
	})
	// 下拉选
	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Select set to", value)
	})
	// 进度条
	progress := widget.NewProgressBar()
	// 进度条加载
	infinite := widget.NewProgressBarInfinite()

	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 250)
			progress.SetValue(i)
		}
	}()

	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 100)),
		label, buttion, input, check, radio, combo, progress, infinite)

	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
