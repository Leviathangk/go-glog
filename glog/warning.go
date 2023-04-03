// Package glog Warning 输出
package glog

// Warning 直接输出
func Warning(a ...any) {
	outPut(DefaultLogger, WarningLevel, a...)
}

func (logger *Logger) Warning(a ...any) {
	outPut(logger, WarningLevel, a...)
}

// Warningln 换行输出
func Warningln(a ...any) {
	outPutln(DefaultLogger, WarningLevel, a...)
}

func (logger *Logger) Warningln(a ...any) {
	outPutln(logger, WarningLevel, a...)
}

// Warningf 自定义输出
func Warningf(format string, a ...any) {
	outPutf(DefaultLogger, WarningLevel, format, a...)
}

func (logger *Logger) Warningf(format string, a ...any) {
	outPutf(logger, WarningLevel, format, a...)
}
