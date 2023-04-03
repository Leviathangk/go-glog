// Package glog 格式化定义
package glog

import "fmt"

// ColorFormatter 颜色格式
type ColorFormatter struct {
	Model           int    // 显示的模式
	ForegroundColor string // 前景色
	BackgroundColor string // 背景色
}

// Formatter 输出格式化
type Formatter struct {
	TimeFormat   string          // 时间格式
	TimeColor    *ColorFormatter // 时间显示格式
	TraceColor   *ColorFormatter // Trace 级别显示格式
	DebugColor   *ColorFormatter // Debug 级别显示格式
	InfoColor    *ColorFormatter // Info 级别显示格式
	WarningColor *ColorFormatter // Warning 级别显示格式
	ErrorColor   *ColorFormatter // Error 级别显示格式
}

// format 根据定义好的颜色，进行格式化
func (c *ColorFormatter) format(s string) string {
	var backgroundStr string
	var foregroundColor string

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
