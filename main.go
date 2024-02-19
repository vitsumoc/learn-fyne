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
	i := 34
	switch i {
	case 0:
		myTest()
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
	// 7 构建
	case 8:
		painter() // 手动绘制
	case 9:
		fLyaout() // 布局能力
	case 10:
		shortcuts() // 窗口快捷键 组件快捷键
	case 11:
		preferences() // 首选项存取
	case 12:
		sysTray() // 系统托盘
	case 13:
		dataBinding() // 数据绑定
	case 14:
		drawing() // 图形绘制
	case 15:
		animation() // 动画
	case 16:
		layoutBox() // box 布局
	case 17:
		layoutGrid() // 网格布局
	case 18:
		layoutGridWarp() // 自适应网格布局
	case 19:
		layoutBorder() // 盒子布局
	case 20:
		layoutForm() // 表单布局
	case 21:
		layoutCenter() // 居中布局
	case 22:
		layoutMax() // 占满布局
	case 23:
		appTab() // 标签页
	case 24:
		fWidgets() // 各种组件
	case 25:
		fForm() // 表单
	case 26:
		fToolBar() // 工具栏
	case 27:
		fContainer() // 容器（列表 表格 树
	case 28:
		fDataBinding() // 数据绑定详细
	case 29:
		customLayout() // 自定义布局
	case 30:
		customWidget() // 自定义组件
		// 31 打包静态资源
	case 32:
		customTheme() // 自定义主题
	case 33:
		extendingWidgets() // 扩展现有组件
	case 34:
		fNumericalEntry() // 自定义数字输入组件
	}
}
