package main

import (
	"github.com/Leviathangk/go-glog/glog"
	"os"
)

func main() {
	// 创建新 logger
	//logger := glog.New(&glog.Config{
	//	Print:   true,
	//	PrintLevel: glog.DebugLevel,
	//	OutPutLevel: glog.DebugLevel,
	//	Formatter:  glog.DefaultConfig,
	//})

	// 默认 logger
	logger := glog.DefaultLogger
	logger.Config.PrintLevel = glog.TraceLevel  // 默认是 debug 级别，这里修改
	logger.Config.OutPutLevel = glog.TraceLevel // 默认是 debug 级别，这里修改

	// AddOutPut
	file, _ := os.OpenFile("demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger.AddOutPut(file)

	// AddHook：实现分类处理
	errFile, _ := os.OpenFile("err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	logger.AddHook(func(level int, out string) {
		if level == glog.ErrorLevel {
			_, err := errFile.Write([]byte(out))
			if err != nil {
				return
			}
		}
	})

	// 正常使用：默认 logger 可以直接使用 glog.Debugln()
	logger.Traceln("Trace")
	logger.Debugln("Debug")
	logger.Infoln("Info")
	logger.Warningln("Warning")
	logger.Errorln("Error")
	//logger.Fatalln("Fatalln 结束程序执行")
	//logger.Panicln("Panic 结束程序执行")
}
