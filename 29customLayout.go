// 自定义布局
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// 自定义布局只需实现 Layout 接口
// type Layout interface {
// 	Layout([]CanvasObject, Size)
// 	MinSize(objects []CanvasObject) Size
// }

type diagonal struct {
}

// 用来计算此布局的最小尺寸
func (d *diagonal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	// 本例中采用每个组件最小尺寸的和
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

// 通过 Layout 实现布局, 即摆放每个 object 的位置
func (d *diagonal) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, containerSize.Height-d.MinSize(objects).Height)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width, size.Height))
	}
}

func customLayout() {
	a := app.New()
	w := a.NewWindow("Diagonal")

	text1 := widget.NewLabel("topleft")
	text2 := widget.NewLabel("Middle Label")
	text3 := widget.NewLabel("bottomright")

	w.SetContent(container.New(&diagonal{}, text1, text2, text3))
	w.ShowAndRun()
}
