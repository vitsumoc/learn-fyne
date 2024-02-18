package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func fDataBinding() {
	myApp := app.New()
	w := myApp.NewWindow("fDataBinding")

	// 直接创建绑定数据
	boundString := binding.NewString()
	s, _ := boundString.Get()
	log.Printf("Bound = '%s'", s)

	// 绑定现有基本数据
	myInt := 5
	boundInt := binding.BindInt(&myInt)
	i, _ := boundInt.Get()
	log.Printf("Source = %d, bound = %d", myInt, i)

	// 创建绑定数据赋值
	str := binding.NewString()
	str.Set("Initial value")

	// 组件 + 绑定数据
	text := widget.NewLabelWithData(str)

	// 绑定数据后修改
	go func() {
		time.Sleep(time.Second * 3)
		str.Set("A new string")
	}()

	// 双向绑定
	s2 := binding.NewString()
	s2.Set("Hi!")

	// 绑定带数据转换
	f := binding.NewFloat()
	s3 := binding.FloatToString(f)
	s4 := binding.FloatToStringWithFormat(f, "%0.0f%%")
	f.Set(25.0)

	// 数组绑定
	listData := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	list := widget.NewListWithData(listData,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", listData.Length()+1)
		listData.Append(val)
	})

	w.SetContent(
		container.NewHBox(text,
			widget.NewLabelWithData(s2),
			widget.NewEntryWithData(s2),
			widget.NewSliderWithData(0, 100.0, f),
			widget.NewLabelWithData(s3),
			widget.NewLabelWithData(s4),
			add,
			list,
		),
	)
	w.ShowAndRun()
}
