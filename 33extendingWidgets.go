// 扩展现有组件
// 本例中扩展图标, 使其支持点击
package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// 可点击图标
type tappableIcon struct {
	widget.Icon
}

func newTappableIcon(res fyne.Resource) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}

func (t *tappableIcon) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")
}

// 右键
func (t *tappableIcon) TappedSecondary(_ *fyne.PointEvent) {
	log.Println("I have been tapped secondary")
}

func extendingWidgets() {
	a := app.New()
	w := a.NewWindow("Tappable")
	w.SetContent(newTappableIcon(theme.FyneLogo()))
	w.ShowAndRun()
}
