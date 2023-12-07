package glog

import (
	"io"
)

// 日志级别
const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// Logger 结构体
type Logger struct {
	ShowColor    bool        // 是否在输出的时候输出颜色（仅控制台）
	ShowCaller   bool        // 显示调用栈
	Print        bool        // 是否输出到控制台
	PrintLevel   int         // 日志限制输出级别（仅控制台）
	Handler      []io.Writer // 其它输出方式，受 OutPutLevel 限制
	HandlerLevel int         // 日志限制输出级别（其余输出）
	Hook         []HookFunc  // 钩子函数
	TimeFormat   string      // 时间格式化
}

// DLogger 默认的 logger
var DLogger = NewLogger()

// NewLogger 新建一个 logger
func NewLogger() *Logger {
	return &Logger{
		ShowColor:    true,
		ShowCaller:   true,
		Print:        true,
		PrintLevel:   LevelTrace,
		HandlerLevel: LevelTrace,
		TimeFormat:   "2006-01-02 15:04:05",
	}
}

// SetTimeFormat 设置时间格式化的格式
func (logger *Logger) SetTimeFormat(format string) {
	logger.TimeFormat = format
}

// AddHandler 增加输出：必须实现 Writer 接口
func (logger *Logger) AddHandler(writer ...io.Writer) {
	logger.Handler = append(logger.Handler, writer...)
}

// HookFunc 钩子函数
type HookFunc func(level int, out string)

// AddHook 添加钩子函数
func (logger *Logger) AddHook(hookFunc ...HookFunc) {
	logger.Hook = append(logger.Hook, hookFunc...)
}
