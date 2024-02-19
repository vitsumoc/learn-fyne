// 自定义数字输入组件
package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

// 自定义数字输入组件
type numericalEntry struct {
	widget.Entry
}

// 其构造函数
func newNumericalEntry() *numericalEntry {
	entry := &numericalEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

// 检查接受的输入, 只接受数字
func (e *numericalEntry) TypedRune(r rune) {
	if r >= '0' && r <= '9' {
		e.Entry.TypedRune(r)
	}
}

// 只允许粘贴数字
func (e *numericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

// 移动端仅允许数字键盘
func (e *numericalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}

func fNumericalEntry() {
	a := app.New()
	w := a.NewWindow("Numerical")

	entry := newNumericalEntry()

	w.SetContent(entry)
	w.ShowAndRun()
}
