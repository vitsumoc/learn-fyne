// 将静态资源放入go代码, 以此可将程序打包为单个可执行文件
package main

// 1. 使用命令行工具绑定资源, 得到 bundled.go 文件
// fyne bundle -o bundled.go image.png
// 2. 程序中直接引用 bundled.go 文件, 使用其中资源即可
