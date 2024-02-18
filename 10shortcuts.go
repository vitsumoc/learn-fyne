package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type myEntry struct {
	widget.Entry
}

// 重载快捷键处理方法
func (m *myEntry) TypedShortcut(s fyne.Shortcut) {
	log.Println("Shortcut typed:", s)
	// 非自定义快捷键 交由父级处理
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		m.Entry.TypedShortcut(s)
		return
	}

	log.Println("Shortcut typed:", s)
}

func shortcuts() {
	// 创建 app
	a := app.New()
	// 创建窗口
	w := a.NewWindow("Hello World")

	// 添加内容 文本标签
	w.SetContent(&myEntry{})

	// 快捷键设置
	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	ctrlAltTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl | fyne.KeyModifierAlt}
	// 事件
	w.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})
	w.Canvas().AddShortcut(ctrlAltTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Alt+Tab")
	})

	// 显示窗口并运行程序
	w.ShowAndRun()
}
