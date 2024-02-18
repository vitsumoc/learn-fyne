// 绑定数据可以由 binding.New 创造或是 binding.bindInt 捆绑现有数据
// 组件可以使用 WithData 自带绑定数据, 或使用 Bind Unbind 手动操作
package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func dataBinding() {
	a := app.New()
	w := a.NewWindow("Hello")

	// 被绑定的字符串
	str := binding.NewString()
	// 字符串修改同步影响组件
	go func() {
		dots := "....."
		for i := 5; i >= 0; i-- {
			str.Set("Count down" + dots[:i])
			time.Sleep(time.Second)
		}
		str.Set("Blast off!")
	}()

	w.SetContent(widget.NewLabelWithData(str))
	w.ShowAndRun()
}
