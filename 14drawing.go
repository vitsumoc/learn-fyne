// 图形绘制
package main

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func drawing() {
	myApp := app.New()
	w := myApp.NewWindow("drawing")

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	// 矩形
	rect := canvas.NewRectangle(blue)
	rect.Resize(fyne.NewSize(100, 100))
	rect.Move(fyne.NewPos(0, 0))

	// 文本
	text := canvas.NewText("Text Object", blue)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Italic: true}
	text.Resize(fyne.NewSize(100, 100))
	text.Move(fyne.NewPos(100, 0))

	// 直线
	line := canvas.NewLine(blue)
	line.StrokeWidth = 4
	line.Position1 = fyne.NewPos(210, 50)
	line.Position2 = fyne.NewPos(290, 50)

	// 圆
	circle := canvas.NewCircle(color.White)
	circle.StrokeColor = blue
	circle.StrokeWidth = 5
	circle.Resize(fyne.NewSize(100, 100))
	circle.Move(fyne.NewPos(300, 0))

	// 图片
	image := canvas.NewImageFromResource(theme.FyneLogo())
	image.FillMode = canvas.ImageFillOriginal
	image.Resize(fyne.NewSize(100, 100))
	image.Move(fyne.NewPos(400, 0))

	// 光栅
	raster := canvas.NewRasterWithPixels(func(_, _, w, h int) color.Color {
		return color.RGBA{R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)), A: 0xff}
	})
	raster.Resize(fyne.NewSize(100, 100))
	raster.Move(fyne.NewPos(500, 0))

	// 渐变
	gradient := canvas.NewHorizontalGradient(color.White, blue)
	gradient.Resize(fyne.NewSize(100, 100))
	gradient.Move(fyne.NewPos(600, 0))

	// 布局
	content := container.NewWithoutLayout(rect, text, line, circle, image, raster, gradient)

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
