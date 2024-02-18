// 系统设置方式，包括 string int float bool
// 每种包括三个方法 存、取、取带默认
// 需要一个唯一的 appID 用来存放系统设置
package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func preferences() {
	a := app.NewWithID("com.vitsumoc.learn-fyne.preferences")
	w := a.NewWindow("Timeout")

	var timeout time.Duration

	// 选择超时时间，存入首选项
	timeoutSelector := widget.NewSelect([]string{"10 seconds", "30 seconds", "1 minute"}, func(selected string) {
		switch selected {
		case "10 seconds":
			timeout = 10 * time.Second
		case "30 seconds":
			timeout = 30 * time.Second
		case "1 minute":
			timeout = time.Minute
		}

		a.Preferences().SetString("AppTimeout", selected)
	})

	// 从首选项取回，赋值，空则用默认
	timeoutSelector.SetSelected(a.Preferences().StringWithFallback("AppTimeout", "10 seconds"))

	// 超时后跑路
	go func() {
		time.Sleep(timeout)
		a.Quit()
	}()

	w.SetContent(timeoutSelector)
	w.ShowAndRun()
}
