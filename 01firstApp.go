package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func firstApp() {
	// 创建 app
	a := app.New()
	// 创建窗口
	w := a.NewWindow("Hello World")

	// 添加内容 文本标签
	w.SetContent(widget.NewLabel("Hello World!"))
	// 显示窗口并运行程序
	w.ShowAndRun()
}
