package main

import (
	"fmt"
)

type Passenger struct{ Name string }

func main() {
	passengers := []Passenger{{"张三"}, {"李四"}, {"王五"}}

	fmt.Println("❌ 错误示例：闭包捕获循环变量 p（几乎必翻车）")
	var funcs []func()

	for _, p := range passengers {
		funcs = append(funcs, func() {
			fmt.Println("wrong:", p.Name)
		})
	}

	// 注意：这里循环已经结束了，p 已经是最后一个乘客（王五）
	for _, f := range funcs {
		f()
	}

	fmt.Println("✅ 正确示例：把 p 复制一份（每个闭包一份副本）")
	funcs = nil

	for _, p := range passengers {
		ps := p // 关键：复制一份
		funcs = append(funcs, func() {
			fmt.Println("right:", ps.Name)
		})
	}

	for _, f := range funcs {
		f()
	}
}
