// 自定义主题
package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// 主题需实现 Theme 接口
// type Theme interface {
// 	Color(ThemeColorName, ThemeVariant) color.Color
// 	Font(TextStyle) Resource
// 	Icon(ThemeIconName) Resource
// 	Size(ThemeSizeName) float32
// }

type myTheme struct{}

// 通过断言确保我们实现了接口
var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	// 背景颜色根据主题变量使用白色或黑色
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return blue
		}
		return color.Black
	}

	// 其他使用默认
	return theme.DefaultTheme().Color(name, variant)
}

// 自定义图标
func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	icon, err := fyne.LoadResourceFromPath("./icon.png")
	if err != nil {
		return theme.DefaultTheme().Icon(name)
	}

	if name == theme.IconNameHome {
		return fyne.NewStaticResource("myHome", icon.Content())
	}

	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func customTheme() {
	// 创建 app
	a := app.New()
	// 应用主题
	a.Settings().SetTheme(&myTheme{})
	// 创建窗口
	w := a.NewWindow("Hello World")

	// 添加内容 文本标签
	w.SetContent(widget.NewLabel("Hello World!"))
	// 显示窗口并运行程序
	w.ShowAndRun()

}
