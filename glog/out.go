// Package glog 统一输出
package glog

import (
	"fmt"
	"os"
	"time"
)

// out 输出
func out(logger *Logger, level int, prefix1, prefix2, str string) {
	var err error

	for _, o := range logger.Config.Out {
		if o == os.Stdout {
			_, err = o.Write([]byte(prefix1))
			if err != nil {
				fmt.Printf("日志输出错误：%s", err)
				return
			}
			_, err = o.Write([]byte(str))
		} else {
			_, err = o.Write([]byte(prefix2))
			if err != nil {
				fmt.Printf("日志输出错误：%s", err)
				return
			}
			_, err = o.Write([]byte(str))
		}
		if err != nil {
			fmt.Printf("日志输出错误：%s", err)
			return
		}
	}

	// 执行 hook
	for _, hookFunc := range logger.Hook {
		hookFunc(level, prefix2+str)
	}
}

// outPut 格式化
func outPut(logger *Logger, level int, a ...any) {
	var outStr string
	var prefix1 string
	var prefix2 string
	timeNow := time.Now().Format(logger.Config.Formatter.TimeFormat)

	if DefaultLogger.Config.Print && DefaultLogger.Config.PrintLevel <= level {
		outStr = fmt.Sprint(a...)
		prefix1 = prefix(logger, level, timeNow, true)
		prefix2 = prefix(logger, level, timeNow, false)
	}

	out(logger, level, prefix1, prefix2, outStr)
}

// outPutln 格式化
func outPutln(logger *Logger, level int, a ...any) {
	var outStr string
	var prefix1 string
	var prefix2 string
	timeNow := time.Now().Format(logger.Config.Formatter.TimeFormat)

	if DefaultLogger.Config.Print && DefaultLogger.Config.PrintLevel <= level {
		outStr = fmt.Sprintln(a...)
		prefix1 = prefix(logger, level, timeNow, true)
		prefix2 = prefix(logger, level, timeNow, false)
	}

	out(logger, level, prefix1, prefix2, outStr)
}

// outPutf 格式化
func outPutf(logger *Logger, level int, format string, a ...any) {
	var outStr string
	var prefix1 string
	var prefix2 string
	timeNow := time.Now().Format(logger.Config.Formatter.TimeFormat)

	if DefaultLogger.Config.Print && DefaultLogger.Config.PrintLevel <= level {
		outStr = fmt.Sprintf(format, a...)
		prefix1 = prefix(logger, level, timeNow, true)
		prefix2 = prefix(logger, level, timeNow, false)
	}

	out(logger, level, prefix1, prefix2, outStr)
}
