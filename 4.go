package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/* ANSI 颜色 */
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Brown  = "\033[38;5;94m"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	height := 12

	// 星星
	printCentered(Yellow+"★"+Reset, height)

	// 树叶
	for i := 0; i < height; i++ {
		width := 2*i + 1
		spaces := height - i - 1
		fmt.Print(strings.Repeat(" ", spaces))

		for j := 0; j < width; j++ {
			if rand.Intn(8) == 0 {
				fmt.Print(randomLight())
			} else {
				fmt.Print(Green + "▲" + Reset)
			}
		}
		fmt.Println()
	}

	// 树干
	for i := 0; i < 3; i++ {
		printCentered(Brown+"▮▮▮"+Reset, height)
	}

	// 底座
	printCentered(Brown+"▔▔▔▔▔"+Reset, height)

	fmt.Println()
	printGreeting()
}

/* 随机彩灯 */
func randomLight() string {
	lights := []string{
		Red + "●" + Reset,
		Yellow + "●" + Reset,
		Blue + "●" + Reset,
	}
	return lights[rand.Intn(len(lights))]
}

/* 居中打印 */
func printCentered(s string, h int) {
	total := 2*h - 1
	spaces := (total - visualLen(s)) / 2
	if spaces < 0 {
		spaces = 0
	}
	fmt.Print(strings.Repeat(" ", spaces))
	fmt.Println(s)
}

/* 计算去掉 ANSI 的长度 */
func visualLen(s string) int {
	n := 0
	esc := false
	for _, r := range s {
		if r == '\033' {
			esc = true
			continue
		}
		if esc {
			if r == 'm' {
				esc = false
			}
			continue
		}
		n++
	}
	return n
}

/* 彩色字符画祝福 */
func printGreeting() {
	text := []string{
		"  圣 诞 快 乐  ",
		" MERRY CHRISTMAS ",
	}

	colors := []string{Red, Green, Yellow, Blue}

	for _, line := range text {
		for i, r := range line {
			color := colors[i%len(colors)]
			fmt.Print(color + string(r) + Reset)
		}
		fmt.Println()
	}
}
