package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func painter() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	// 获得窗口画布
	myCanvas := myWindow.Canvas()

	// 颜色 形状 设置内容
	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	myCanvas.SetContent(rect)

	// 画布内容可更新
	go func(w fyne.Window, c fyne.Canvas) {
		// 更新颜色
		time.Sleep(time.Second * 3)
		green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		rect.FillColor = green
		rect.Refresh()

		// 换为文本
		time.Sleep(time.Second * 3)
		setContentToText(c)

		// 换为其他图形
		time.Sleep(time.Second * 3)
		setContentToCircle(c)

		// 换为 widget
		time.Sleep(time.Second * 3)
		w.SetContent(widget.NewEntry())
	}(myWindow, myCanvas)

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}

func setContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}
