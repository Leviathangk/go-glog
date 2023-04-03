// Package glog 方法
package glog

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
)

// New 创建记录器
func New(config *Config) (logger *Logger) {
	logger = new(Logger)

	// 对配置进行判断
	if len(config.Out) == 0 {
		config.Out = append(config.Out, os.Stdout)
	}
	if config.Formatter == nil {
		config.Formatter = DefaultConfig.Formatter
	}
	if config.Formatter.TimeFormat == "" {
		config.Formatter.TimeFormat = DefaultConfig.Formatter.TimeFormat
	}
	if config.Formatter.TimeColor == nil {
		config.Formatter.TimeColor = DefaultTimeColor
	}
	if config.Formatter.TraceColor == nil {
		config.Formatter.TraceColor = DefaultTraceColor
	}
	if config.Formatter.DebugColor == nil {
		config.Formatter.DebugColor = DefaultDebugColor
	}
	if config.Formatter.InfoColor == nil {
		config.Formatter.InfoColor = DefaultInfoColor
	}
	if config.Formatter.WarningColor == nil {
		config.Formatter.WarningColor = DefaultWarningColor
	}
	if config.Formatter.ErrorColor == nil {
		config.Formatter.ErrorColor = DefaultErrorColor
	}
	logger.Config = config

	return
}

// prefix 输出格式化的前缀
func prefix(logger *Logger, level int, timeNow string, color bool) string {
	// 格式化时间戳
	if color {
		timeNow = logger.Config.Formatter.TimeColor.format(timeNow)
	}

	// 格式化日志级别
	var levelStr string

	switch level {
	case TraceLevel:
		if color {
			levelStr = logger.Config.Formatter.TraceColor.format("Trace") + "  "
		} else {
			levelStr = "Trace" + "  "
		}
	case DebugLevel:
		if color {
			levelStr = logger.Config.Formatter.DebugColor.format("Debug") + "  "
		} else {
			levelStr = "Debug" + "  "
		}
	case InfoLevel:
		if color {
			levelStr = logger.Config.Formatter.InfoColor.format("Info") + "   "
		} else {
			levelStr = "Info" + "   "
		}
	case WarningLevel:
		if color {
			levelStr = logger.Config.Formatter.WarningColor.format("Warning")
		} else {
			levelStr = "Warning"
		}
	case ErrorLevel:
		if color {
			levelStr = logger.Config.Formatter.ErrorColor.format("Error") + "  "
		} else {
			levelStr = "Error" + "  "
		}
	case FatalLevel:
		levelStr = "Fatal"
	default:
		if color {
			levelStr = logger.Config.Formatter.InfoColor.format("Unknow") + " "
		} else {
			levelStr = "Unknow" + " "
		}
	}

	// 构造出
	callerDetail := GetCallerDetail(4)

	// 构造输出前缀字符
	outStr := fmt.Sprintf("%s | %s | %s:%d - ", levelStr, timeNow, callerDetail.Name, callerDetail.Line)

	return outStr
}

// GetCallerDetail 获取调用者的信息
func GetCallerDetail(deep int) (callerDetail *CallerDetail) {
	callerDetail = new(CallerDetail)

	pc, file, line, ok := runtime.Caller(deep) // 返回调用堆栈、文件、调用行号，这里的 deep 是堆栈深度
	if ok {
		callerDetail.Name = runtime.FuncForPC(pc).Name()
		callerDetail.Line = line
		callerDetail.FilePath = file
		callerDetail.FileName = path.Base(file)
	}

	return
}

// AddOutPut 增加输出：必须实现 Writer 接口
func (logger *Logger) AddOutPut(writer ...io.Writer) {
	for _, w := range writer {
		exists := false
		for _, o := range logger.Config.Out {
			if o == w {
				exists = true
				break
			}
		}

		if !exists {
			logger.Config.Out = append(logger.Config.Out, w)
		}
	}
}

// AddHook 添加钩子函数
func (logger *Logger) AddHook(hookFunc HookFunc) {
	logger.Hook = append(logger.Hook, hookFunc)
}
