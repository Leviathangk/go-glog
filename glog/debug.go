// Package glog Debug 输出
package glog

// Debug 直接输出
func Debug(a ...any) {
	outPut(DefaultLogger, DebugLevel, a...)
}

func (logger *Logger) Debug(a ...any) {
	outPut(logger, DebugLevel, a...)
}

// Debugln 换行输出
func Debugln(a ...any) {
	outPutln(DefaultLogger, DebugLevel, a...)
}

func (logger *Logger) Debugln(a ...any) {
	outPutln(logger, DebugLevel, a...)
}

// Debugf 自定义输出
func Debugf(format string, a ...any) {
	outPutf(DefaultLogger, DebugLevel, format, a...)
}

func (logger *Logger) Debugf(format string, a ...any) {
	outPutf(logger, DebugLevel, format, a...)
}
