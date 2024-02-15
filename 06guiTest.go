package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// 这里带了一个通过 in 改变 out 的逻辑，并对此逻辑进行测试
func makeUI() (*widget.Label, *widget.Entry) {
	// 文本标签
	out := widget.NewLabel("Hello world!")
	// input输入
	in := widget.NewEntry()
	// 输入事件
	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}

func guiTest() {
	a := app.New()
	w := a.NewWindow("Hello Person")

	w.SetContent(container.NewVBox(makeUI()))
	w.ShowAndRun()
}
