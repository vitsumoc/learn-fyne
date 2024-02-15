package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func windows() {
	a := app.New()
	w := a.NewWindow("Hello World")

	// 窗口1 主窗口
	w.SetContent(widget.NewLabel("我是主窗口，关我关全部"))
	w.SetMaster()
	w.Show()

	// 窗口2 大一点 点击打开窗口3
	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewButton("窗口3", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	}))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	a.Run()
}
