package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestGreeting(t *testing.T) {
	out, in := makeUI()

	// 测试初始获取的 out 标签
	if out.Text != "Hello world!" {
		t.Error("Incorrect initial greeting")
	}

	// 模拟对 in 的输入
	test.Type(in, "Andy")

	// 测试修改后的 out 标签
	if out.Text != "Hello Andy!" {
		t.Error("Incorrect user greeting")
	}
}
