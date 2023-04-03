// Package glog Trace 输出
package glog

// Trace 直接输出
func Trace(a ...any) {
	outPut(DefaultLogger, TraceLevel, a...)
}

func (logger *Logger) Trace(a ...any) {
	outPut(logger, TraceLevel, a...)
}

// Traceln 换行输出
func Traceln(a ...any) {
	outPutln(DefaultLogger, TraceLevel, a...)
}

func (logger *Logger) Traceln(a ...any) {
	outPutln(logger, TraceLevel, a...)
}

// Tracef 自定义输出
func Tracef(format string, a ...any) {
	outPutf(DefaultLogger, TraceLevel, format, a...)
}

func (logger *Logger) Tracef(format string, a ...any) {
	outPutf(logger, TraceLevel, format, a...)
}
