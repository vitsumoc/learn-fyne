package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func updateContent() {
	// app 和 窗口
	a := app.New()
	w := a.NewWindow("Clock")

	// 时间标签
	clock := widget.NewLabel("")
	updateTime(clock)

	// 置入窗口
	w.SetContent(clock)
	// 定时更新标签内容
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.ShowAndRun()
}
