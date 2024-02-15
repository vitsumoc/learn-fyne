package main

import (
	"os"

	"github.com/flopp/go-findfont"
	"github.com/goki/freetype/truetype"
)

// 加载中文字体
func init() {
	fontPath, err := findfont.Find("simhei.ttf")
	if err != nil {
		panic(err)
	}

	// load the font with the freetype library
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	_, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	os.Setenv("FYNE_FONT", fontPath)
}

func main() {
	i := 6
	switch i {
	case 1:
		firstApp() // 基础的程序、窗口
	case 2:
		demo() // fyne 官方示例
	case 3:
		appLoop() // 窗口和主循环
	case 4:
		updateContent() // 内容更新
	case 5:
		windows() // 窗口管理
	case 6:
		guiTest() // 图形化程序测试
	}
}
