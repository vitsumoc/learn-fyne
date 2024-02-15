package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func appLoop() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello"))

	// 等同于 showAndRun()
	myWindow.Show()
	myApp.Run()

	// 在退出时执行
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
