// Package glog 统一输出
package glog

import (
	"fmt"
	"os"
	"time"
)

// out 输出
func out(logger *Logger, level int, str string) {
	var err error

	// 形成前缀的格式化字符
	timeNow := time.Now().Format(logger.Config.Formatter.TimeFormat)
	prefix1 := prefix(logger, level, timeNow, true)
	prefix2 := prefix(logger, level, timeNow, false)

	// 检查是否是 panic error 保证退出
	defer func() {
		if level == PanicLevel {
			panic(prefix2 + str)
		}
	}()

	// 控制台输出
	if logger.Config.Print && logger.Config.PrintLevel <= level {
		_, err = os.Stdout.Write([]byte(prefix1))
		if err == nil {
			_, err = os.Stdout.Write([]byte(str))
		}

		if err != nil {
			fmt.Printf("日志输出错误：%s\n", err)
		}
	}

	// AddOutPut 进来的
	for _, o := range logger.Config.Out {
		if logger.Config.OutPutLevel <= level {
			_, err = o.Write([]byte(prefix2))
			if err == nil {
				_, err = o.Write([]byte(str))
			}
		}

		if err != nil {
			fmt.Printf("日志输出错误：%s\n", err)
		}
	}

	// AddHook 进来的
	for _, hookFunc := range logger.Hook {
		hookFunc(level, prefix2+str)
	}
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
