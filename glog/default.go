// Package glog 控制台输出学习链接：https://zhuanlan.zhihu.com/p/76751728
package glog

// DefaultLogger 默认记录器
var DefaultLogger = NewLogger(DefaultConfig)

// DefaultConfig 默认配置
var DefaultConfig = &Config{
	Print:       true,
	PrintLevel:  DebugLevel,
	OutPutLevel: DebugLevel,
	Formatter:   DefaultFormatter,
}

// DefaultFormatter 默认 Formatter
var DefaultFormatter = &Formatter{
	TimeFormat:   "2006-01-02 15:04:05",
	TimeColor:    DefaultTimeColor,
	TraceColor:   DefaultTraceColor,
	DebugColor:   DefaultDebugColor,
	InfoColor:    DefaultInfoColor,
	WarningColor: DefaultWarningColor,
	ErrorColor:   DefaultErrorColor,
	FatalColor:   DefaultFatalColor,
	PanicColor:   DefaultPanicColor,
}

// DefaultTimeColor 时间默认显示的颜色
var DefaultTimeColor = &ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "32", // 前景色
	BackgroundColor: "",   // 背景色
}

// DefaultTraceColor Trace 级别默认显示的颜色
var DefaultTraceColor = &ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "37", // 前景色
	BackgroundColor: "",   // 背景色
}

// DefaultDebugColor Debug 级别默认显示的颜色
var DefaultDebugColor = &ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "34", // 前景色
	BackgroundColor: "",   // 背景色
}

// DefaultInfoColor Info 级别默认显示的颜色
var DefaultInfoColor = &ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "36", // 前景色
	BackgroundColor: "",   // 背景色
}

// DefaultWarningColor Warning 级别默认显示的颜色
var DefaultWarningColor = &ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "33", // 前景色
	BackgroundColor: "",   // 背景色
}

// DefaultErrorColor Error 级别默认显示的颜色
var DefaultErrorColor = &ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "31", // 前景色
	BackgroundColor: "",   // 背景色
}

// DefaultFatalColor Fatal 级别默认显示的颜色
var DefaultFatalColor = &ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "30", // 前景色
	BackgroundColor: "",   // 背景色
}

// DefaultPanicColor Panic 级别默认显示的颜色
var DefaultPanicColor = &ColorFormatter{
	Model:           1,    // 显示模式
	ForegroundColor: "35", // 前景色
	BackgroundColor: "",   // 背景色
}
