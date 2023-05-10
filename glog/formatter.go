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
	TimeFormat string          // 时间格式
	TimeColor  *ColorFormatter // 时间显示格式
	TraceColor *ColorFormatter // Trace 级别显示格式
	DebugColor *ColorFormatter // Debug 级别显示格式
	InfoColor  *ColorFormatter // Info 级别显示格式
	WarnColor  *ColorFormatter // Warn 级别显示格式
	ErrorColor *ColorFormatter // Error 级别显示格式
	FatalColor *ColorFormatter // Fatal 级别显示格式
	PanicColor *ColorFormatter // Panic 级别显示格式
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

// prefix 输出格式化的前缀
func prefix(logger *Logger, level int, timeNow string, color bool) string {
	// 格式化时间戳
	if color {
		timeNow = logger.Config.Formatter.TimeColor.format(timeNow)
	}

	// 格式化日志级别
	var levelStr string
	var outStr string

	switch level {
	case TraceLevel:
		if color {
			levelStr = logger.Config.Formatter.TraceColor.format("Trace")
		} else {
			levelStr = "Trace"
		}
	case DebugLevel:
		if color {
			levelStr = logger.Config.Formatter.DebugColor.format("Debug")
		} else {
			levelStr = "Debug"
		}
	case InfoLevel:
		if color {
			levelStr = logger.Config.Formatter.InfoColor.format("Info") + " "
		} else {
			levelStr = "Info" + " "
		}
	case WarnLevel:
		if color {
			levelStr = logger.Config.Formatter.WarnColor.format("Warn") + " "
		} else {
			levelStr = "Warn" + " "
		}
	case ErrorLevel:
		if color {
			levelStr = logger.Config.Formatter.ErrorColor.format("Error")
		} else {
			levelStr = "Error"
		}
	case FatalLevel:
		if color {
			levelStr = logger.Config.Formatter.FatalColor.format("Fatal")
		} else {
			levelStr = "Fatal"
		}
	case PanicLevel:
		if color {
			levelStr = logger.Config.Formatter.PanicColor.format("Panic")
		} else {
			levelStr = "Panic"
		}
	default:
		if color {
			levelStr = logger.Config.Formatter.InfoColor.format("Unknow")
		} else {
			levelStr = "Unknow"
		}
	}

	if logger.Config.ShowCaller {
		// 构造出
		callerDetail := GetCallerDetail(5)

		// 构造输出前缀字符
		outStr = fmt.Sprintf("%s | %s | %s:%d - ", levelStr, timeNow, callerDetail.Name, callerDetail.Line)

	} else {
		// 构造输出前缀字符
		outStr = fmt.Sprintf("%s | %s - ", levelStr, timeNow)

	}

	return outStr
}
