package main

import (
	"github.com/Leviathangk/go-glog/glog"
	"os"
)

func main() {
	// 获取 Logger
	//logger := glog.NewLogger()	// 创建新的 logger
	logger := glog.DLogger    // 使用默认的 logger
	logger.ShowColor = false  // 禁用控制台显示颜色
	logger.ShowCaller = false // 不显示堆栈信息

	// 修改日志级别
	logger.PrintLevel = glog.LevelTrace   // 日志输出级别（控制台）
	logger.HandlerLevel = glog.LevelTrace // 日志输出级别（handler）

	// 添加 handler
	file, _ := os.OpenFile("demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger.AddHandler(file)

	// 添加 hook 函数：示例保存错误日志（hook 是为了更自由的控制，当然可以直接使用 handler 统一控制）
	errFile, _ := os.OpenFile("err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	logger.AddHook(func(level int, out string) {
		if level >= glog.LevelError {
			_, err := errFile.Write([]byte(out))
			if err != nil {
				return
			}
		}
	})

	// 正常的日志输出
	logger.Traceln("Trace")
	logger.Debugln("Debug")
	logger.Infoln("Info")
	logger.Warnln("Warn")
	logger.Errorln("Error")
	logger.Fatalln("Fatal")
}
