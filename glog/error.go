// Package glog Error 输出
package glog

// Error 直接输出
func Error(a ...any) {
	outPut(DefaultLogger, ErrorLevel, a...)
}

func (logger *Logger) Error(a ...any) {
	outPut(logger, ErrorLevel, a...)
}

// Errorln 换行输出
func Errorln(a ...any) {
	outPutln(DefaultLogger, ErrorLevel, a...)
}

func (logger *Logger) Errorln(a ...any) {
	outPutln(logger, ErrorLevel, a...)
}

// Errorf 自定义输出
func Errorf(format string, a ...any) {
	outPutf(DefaultLogger, ErrorLevel, format, a...)
}

func (logger *Logger) Errorf(format string, a ...any) {
	outPutf(logger, ErrorLevel, format, a...)
}
