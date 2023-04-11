package glog

import "io"

// 日志级别
const (
	TraceLevel = iota
	DebugLevel
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

// Config 配置
type Config struct {
	Print       bool        // 是否进行打印：默认配置是是
	PrintLevel  int         // 打印级别：默认配置是 DebugLevel
	Out         []io.Writer // 输出日志的地方：可以用来自动保存日志
	OutPutLevel int         // AddOutPut 的级别：默认是 DebugLevel
	Formatter   *Formatter  // 输出格式化
}

// HookFunc 钩子函数
type HookFunc func(level int, out string)

// Logger 记录器，代表一个
type Logger struct {
	Hook   []HookFunc // hook 方法
	Config *Config
}

// CallerDetail 调用者信息
type CallerDetail struct {
	Name     string // 调用者
	Line     int    // 调用行号
	FilePath string // 调用文件
	FileName string // 调用文件名
}

// DefaultLogger 默认记录器
var DefaultLogger = NewLogger(DefaultConfig)
