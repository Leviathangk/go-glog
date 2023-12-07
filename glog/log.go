// Package glog 日志级别的输出
package glog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

// CallerDetail 调用者信息
type CallerDetail struct {
	Name     string // 调用者
	Line     int    // 调用行号
	FilePath string // 调用文件
	FileName string // 调用文件名
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

// prefix 输出格式化的前缀
func prefix(logger *Logger, level int, timeNow string, color bool) string {
	// 格式化日志级别
	var levelStr string
	var outStr string

	switch level {
	case LevelTrace:
		if color {
			levelStr = ColorTrace.format("Trace")
		} else {
			levelStr = "Trace"
		}
	case LevelDebug:
		if color {
			levelStr = ColorDebug.format("Debug")
		} else {
			levelStr = "Debug"
		}
	case LevelInfo:
		if color {
			levelStr = ColorInfo.format("Info") + " "
		} else {
			levelStr = "Info" + " "
		}
	case LevelWarn:
		if color {
			levelStr = ColorWarn.format("Warn") + " "
		} else {
			levelStr = "Warn" + " "
		}
	case LevelError:
		if color {
			levelStr = ColorError.format("Error")
		} else {
			levelStr = "Error"
		}
	case LevelFatal:
		if color {
			levelStr = ColorFatal.format("Fatal")
		} else {
			levelStr = "Fatal"
		}
	case LevelPanic:
		if color {
			levelStr = ColorPanic.format("Panic")
		} else {
			levelStr = "Panic"
		}
	default:
		if color {
			levelStr = ColorInfo.format("Unknow")
		} else {
			levelStr = "Unknow"
		}
	}

	// 格式化时间戳
	if color {
		timeNow = ColorTime.format(timeNow)
	}

	if logger.ShowCaller {
		// 构造出
		callerDetail := GetCallerDetail(5)

		// 构造输出前缀字符
		outStr = fmt.Sprintf("%s | %s | %s:%d - ", levelStr, timeNow, callerDetail.Name, callerDetail.Line)

	} else {
		// 构造输出前缀字符
		outStr = fmt.Sprintf("%s | %s - ", levelStr, timeNow)

	}

	return outStr
}

// out 输出
func out(logger *Logger, level int, str string) {
	var err error
	var prefixColor string

	// 形成前缀的格式化字符
	timeNow := time.Now().Format(logger.TimeFormat)
	prefixDefault := prefix(logger, level, timeNow, false)

	if logger.ShowColor && logger.Print {
		prefixColor = prefix(logger, level, timeNow, true)
	} else {
		prefixColor = prefixDefault
	}

	// AddHook 进来的
	for _, hookFunc := range logger.Hook {
		hookFunc(level, prefixDefault+str)
	}

	// 控制台输出
	if logger.Print && logger.PrintLevel <= level {
		_, err = os.Stdout.Write([]byte(prefixColor + str))

		if err != nil {
			fmt.Printf("日志输出错误：%s\n", err)
		}
	}

	// AddOutPut 进来的
	for _, o := range logger.Handler {
		if logger.HandlerLevel <= level {
			_, err = o.Write([]byte(prefixDefault + str))

			if err != nil {
				fmt.Printf("日志输出错误：%s\n", err)
			}
		}
	}

	// 检查是否是 panic error 保证退出
	defer func() {
		if level == LevelPanic {
			panic(prefixDefault + str)
		}
	}()
}

// outPut 格式化
func outPut(logger *Logger, level int, a ...any) {
	out(logger, level, fmt.Sprint(a...))
}

// outPutln 格式化
func outPutln(logger *Logger, level int, a ...any) {
	out(logger, level, fmt.Sprintln(a...))
}

// outPutf 格式化
func outPutf(logger *Logger, level int, format string, a ...any) {
	out(logger, level, fmt.Sprintf(format, a...))
}

// logger 直接调用的方法

func (logger *Logger) Trace(a ...any) {
	outPut(logger, LevelTrace, a...)
}

func (logger *Logger) Traceln(a ...any) {
	outPutln(logger, LevelTrace, a...)
}

func (logger *Logger) Tracef(format string, a ...any) {
	outPutf(logger, LevelTrace, format, a...)
}

func (logger *Logger) Debug(a ...any) {
	outPut(logger, LevelDebug, a...)
}

func (logger *Logger) Debugln(a ...any) {
	outPutln(logger, LevelDebug, a...)
}

func (logger *Logger) Debugf(format string, a ...any) {
	outPutf(logger, LevelDebug, format, a...)
}

func (logger *Logger) Info(a ...any) {
	outPut(logger, LevelInfo, a...)
}

func (logger *Logger) Infoln(a ...any) {
	outPutln(logger, LevelInfo, a...)
}

func (logger *Logger) Infof(format string, a ...any) {
	outPutf(logger, LevelInfo, format, a...)
}

func (logger *Logger) Warn(a ...any) {
	outPut(logger, LevelWarn, a...)
}

func (logger *Logger) Warnln(a ...any) {
	outPutln(logger, LevelWarn, a...)
}

func (logger *Logger) Warnf(format string, a ...any) {
	outPutf(logger, LevelWarn, format, a...)
}

func (logger *Logger) Error(a ...any) {
	outPut(logger, LevelError, a...)
}

func (logger *Logger) Errorln(a ...any) {
	outPutln(logger, LevelError, a...)
}

func (logger *Logger) Errorf(format string, a ...any) {
	outPutf(logger, LevelError, format, a...)
}

func (logger *Logger) Fatal(a ...any) {
	defer func() {
		os.Exit(1)
	}()
	outPut(logger, LevelFatal, a...)
}

func (logger *Logger) Fatalln(a ...any) {
	defer func() {
		os.Exit(1)
	}()
	outPutln(logger, LevelFatal, a...)
}

func (logger *Logger) Fatalf(format string, a ...any) {
	defer func() {
		os.Exit(1)
	}()
	outPutf(logger, LevelFatal, format, a...)
}

func (logger *Logger) Panic(a ...any) {
	outPut(logger, LevelPanic, a...)
}

func (logger *Logger) Panicln(a ...any) {
	outPutln(logger, LevelPanic, a...)
}

func (logger *Logger) Panicf(format string, a ...any) {
	outPutf(logger, LevelPanic, format, a...)
}
