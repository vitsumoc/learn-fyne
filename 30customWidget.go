// 自定义组件
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// 自定义组件需要实现 Widget 接口, 用来实现数据状态
// 此接口通常会被封装为 BaseWidget, 自定义组件直接使用
// type Widget interface {
// 	CanvasObject

// 	CreateRenderer() WidgetRenderer
// }

// WidgetRenderer 接口则是描述了组件如何被绘制
// type WidgetRenderer interface {
// 	Layout(Size)
// 	MinSize() Size

// 	Refresh()
// 	Objects() []CanvasObject
// 	Destroy()
// }

type MyListItemWidget struct {
	widget.BaseWidget
	Title   *widget.Label
	Comment *widget.Label
}

func NewMyListItemWidget(title, comment string) *MyListItemWidget {
	item := &MyListItemWidget{
		Title:   widget.NewLabel(title),
		Comment: widget.NewLabel(comment),
	}
	item.Title.Truncation = fyne.TextTruncateEllipsis
	item.ExtendBaseWidget(item)

	return item
}

// 使用 simpleRender 的简单例子
func (item *MyListItemWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(nil, nil, nil, item.Comment, item.Title)
	return widget.NewSimpleRenderer(c)
}

func customWidget() {
	a := app.New()
	w := a.NewWindow("Diagonal")

	w1 := NewMyListItemWidget("我是标题", "我是内容")

	w.SetContent(w1)
	w.ShowAndRun()
}
