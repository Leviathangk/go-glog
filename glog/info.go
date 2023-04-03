// Package glog Info 输出
package glog

// Info 直接输出
func Info(a ...any) {
	outPut(DefaultLogger, InfoLevel, a...)
}

func (logger *Logger) Info(a ...any) {
	outPut(logger, InfoLevel, a...)
}

// Infoln 换行输出
func Infoln(a ...any) {
	outPutln(DefaultLogger, InfoLevel, a...)
}

func (logger *Logger) Infoln(a ...any) {
	outPutln(logger, InfoLevel, a...)
}

// Infof 自定义输出
func Infof(format string, a ...any) {
	outPutf(DefaultLogger, InfoLevel, format, a...)
}

func (logger *Logger) Infof(format string, a ...any) {
	outPutf(logger, InfoLevel, format, a...)
}
