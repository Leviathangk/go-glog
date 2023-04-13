package glog

import (
	"io"
	"path"
	"runtime"
)

// 日志级别
const (
	TraceLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

// Config 配置
type Config struct {
	Print       bool       // 是否进行打印：默认配置是是
	PrintLevel  int        // 打印级别：默认配置是 DebugLevel
	OutPutLevel int        // AddOutPut 的级别：默认是 DebugLevel
	Formatter   *Formatter // 输出格式化
}

// HookFunc 钩子函数
type HookFunc func(level int, out string)

// Logger 记录器，代表一个
type Logger struct {
	Out    []io.Writer // 输出日志的地方：可以用来自动保存日志
	Hook   []HookFunc  // hook 方法
	Config *Config
}

// CallerDetail 调用者信息
type CallerDetail struct {
	Name     string // 调用者
	Line     int    // 调用行号
	FilePath string // 调用文件
	FileName string // 调用文件名
}

// NewLogger 创建记录器
func NewLogger(config *Config) (logger *Logger) {
	logger = new(Logger)

	// 对配置进行判断
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
	if config.Formatter.WarnColor == nil {
		config.Formatter.WarnColor = DefaultWarnColor
	}
	if config.Formatter.ErrorColor == nil {
		config.Formatter.ErrorColor = DefaultErrorColor
	}
	if config.Formatter.FatalColor == nil {
		config.Formatter.FatalColor = DefaultFatalColor
	}
	if config.Formatter.PanicColor == nil {
		config.Formatter.PanicColor = DefaultPanicColor
	}
	logger.Config = config

	return
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
