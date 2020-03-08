package log

import (
	"fmt"
)

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见
var (
	conf = 0
	bg   = 0
)

func Black(msg string) string {
	return SetColor(msg, TextBlack)
}

func Red(msg string) string {
	return SetColor(msg, TextRed)
}

func Green(msg string) string {
	return SetColor(msg, TextGreen)
}

func Yellow(msg string) string {
	return SetColor(msg, TextYellow)
}

func Blue(msg string) string {
	return SetColor(msg, TextBlue)
}

func Magenta(msg string) string {
	return SetColor(msg, TextMagenta)
}

func Cyan(msg string) string {
	return SetColor(msg, TextCyan)
}

func White(msg string) string {
	return SetColor(msg, TextWhite)
}

func SetColor(msg string, color int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, color, msg, 0x1B)
}
