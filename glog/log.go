// Package glog 日志级别的输出
package glog

import "os"

// logger 直接调用的方法

func (logger *Logger) Trace(a ...any) {
	outPut(logger, TraceLevel, a...)
}

func (logger *Logger) Traceln(a ...any) {
	outPutln(logger, TraceLevel, a...)
}

func (logger *Logger) Tracef(format string, a ...any) {
	outPutf(logger, TraceLevel, format, a...)
}

func (logger *Logger) Debug(a ...any) {
	outPut(logger, DebugLevel, a...)
}

func (logger *Logger) Debugln(a ...any) {
	outPutln(logger, DebugLevel, a...)
}

func (logger *Logger) Debugf(format string, a ...any) {
	outPutf(logger, DebugLevel, format, a...)
}

func (logger *Logger) Info(a ...any) {
	outPut(logger, InfoLevel, a...)
}

func (logger *Logger) Infoln(a ...any) {
	outPutln(logger, InfoLevel, a...)
}

func (logger *Logger) Infof(format string, a ...any) {
	outPutf(logger, InfoLevel, format, a...)
}

func (logger *Logger) Warning(a ...any) {
	outPut(logger, WarningLevel, a...)
}

func (logger *Logger) Warningln(a ...any) {
	outPutln(logger, WarningLevel, a...)
}

func (logger *Logger) Warningf(format string, a ...any) {
	outPutf(logger, WarningLevel, format, a...)
}

func (logger *Logger) Error(a ...any) {
	outPut(logger, ErrorLevel, a...)
}

func (logger *Logger) Errorln(a ...any) {
	outPutln(logger, ErrorLevel, a...)
}

func (logger *Logger) Errorf(format string, a ...any) {
	outPutf(logger, ErrorLevel, format, a...)
}

func (logger *Logger) Fatal(a ...any) {
	defer func() {
		os.Exit(1)
	}()
	outPut(logger, FatalLevel, a...)
}

func (logger *Logger) Fatalln(a ...any) {
	defer func() {
		os.Exit(1)
	}()
	outPutln(logger, FatalLevel, a...)
}

func (logger *Logger) Fatalf(format string, a ...any) {
	defer func() {
		os.Exit(1)
	}()
	outPutf(logger, FatalLevel, format, a...)
}

func (logger *Logger) Panic(a ...any) {
	outPut(logger, PanicLevel, a...)
}

func (logger *Logger) Panicln(a ...any) {
	outPutln(logger, PanicLevel, a...)
}

func (logger *Logger) Panicf(format string, a ...any) {
	outPutf(logger, PanicLevel, format, a...)
}

// 为了默认 logger 方便调用

func Trace(a ...any) {
	outPut(DefaultLogger, TraceLevel, a...)
}

func Traceln(a ...any) {
	outPutln(DefaultLogger, TraceLevel, a...)
}

func Tracef(format string, a ...any) {
	outPutf(DefaultLogger, TraceLevel, format, a...)
}

func Debug(a ...any) {
	outPut(DefaultLogger, DebugLevel, a...)
}

func Debugln(a ...any) {
	outPutln(DefaultLogger, DebugLevel, a...)
}

func Debugf(format string, a ...any) {
	outPutf(DefaultLogger, DebugLevel, format, a...)
}

func Info(a ...any) {
	outPut(DefaultLogger, InfoLevel, a...)
}

func Infoln(a ...any) {
	outPutln(DefaultLogger, InfoLevel, a...)
}

func Infof(format string, a ...any) {
	outPutf(DefaultLogger, InfoLevel, format, a...)
}

func Warning(a ...any) {
	outPut(DefaultLogger, WarningLevel, a...)
}

func Warningln(a ...any) {
	outPutln(DefaultLogger, WarningLevel, a...)
}

func Warningf(format string, a ...any) {
	outPutf(DefaultLogger, WarningLevel, format, a...)
}

func Error(a ...any) {
	outPut(DefaultLogger, ErrorLevel, a...)
}

func Errorln(a ...any) {
	outPutln(DefaultLogger, ErrorLevel, a...)
}

func Errorf(format string, a ...any) {
	outPutf(DefaultLogger, ErrorLevel, format, a...)
}

func Fatal(a ...any) {
	defer func() {
		os.Exit(1)
	}()
	outPut(DefaultLogger, FatalLevel, a...)
}

func Fatalln(a ...any) {
	defer func() {
		os.Exit(1)
	}()
	outPutln(DefaultLogger, FatalLevel, a...)
}

func Fatalf(format string, a ...any) {
	defer func() {
		os.Exit(1)
	}()
	outPutf(DefaultLogger, FatalLevel, format, a...)
}

func Panic(a ...any) {
	outPut(DefaultLogger, PanicLevel, a...)
}

func Panicln(a ...any) {
	outPutln(DefaultLogger, PanicLevel, a...)
}

func Panicf(format string, a ...any) {
	outPutf(DefaultLogger, PanicLevel, format, a...)
}
