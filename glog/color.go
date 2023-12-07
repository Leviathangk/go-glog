package glog

import "fmt"

// ColorFormatter 颜色格式
type ColorFormatter struct {
	Model           int    // 显示的模式
	ForegroundColor string // 前景色
	BackgroundColor string // 背景色
}

// format 根据定义好的颜色，进行格式化
func (c ColorFormatter) format(s string) string {
	var foregroundColor string
	var backgroundStr string

	if c.ForegroundColor != "" {
		foregroundColor = ";" + c.ForegroundColor
	} else {
		foregroundColor = c.ForegroundColor
	}

	if c.BackgroundColor != "" {
		backgroundStr = ";" + c.BackgroundColor
	} else {
		backgroundStr = c.BackgroundColor
	}
	return fmt.Sprintf("\033[%d%s%sm%s\033[0m", c.Model, foregroundColor, backgroundStr, s)
}

// ColorTime 时间默认显示的颜色
var ColorTime = ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "32", // 前景色
	BackgroundColor: "",   // 背景色
}

// ColorTrace Trace 级别默认显示的颜色
var ColorTrace = ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "37", // 前景色
	BackgroundColor: "",   // 背景色
}

// ColorInfo Info 级别默认显示的颜色
var ColorInfo = ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "36", // 前景色
	BackgroundColor: "",   // 背景色
}

// ColorDebug Debug 级别默认显示的颜色
var ColorDebug = ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "34", // 前景色
	BackgroundColor: "",   // 背景色
}

// ColorWarn Warn 级别默认显示的颜色
var ColorWarn = ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "33", // 前景色
	BackgroundColor: "",   // 背景色
}

// ColorError Error 级别默认显示的颜色
var ColorError = ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "31", // 前景色
	BackgroundColor: "",   // 背景色
}

// ColorFatal Fatal 级别默认显示的颜色
var ColorFatal = ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "30", // 前景色
	BackgroundColor: "",   // 背景色
}

// ColorPanic Panic 级别默认显示的颜色
var ColorPanic = ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "35", // 前景色
	BackgroundColor: "",   // 背景色
}
