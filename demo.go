package main

import (
	"github.com/Leviathangk/go-glog/glog"
	"os"
)

func main() {
	// 创建新 logger
	//logger := glog.New(&glog.Config{
	//	UsePrint:   true,
	//	PrintLevel: glog.DebugLevel,
	//	Formatter:  glog.DefaultConfig,
	//})

	// 默认 logger
	logger := glog.DefaultLogger
	logger.Config.PrintLevel = glog.TraceLevel // 默认是 debug 级别，这里修改

	// 保存
	file, _ := os.OpenFile("demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger.AddOutPut(file)

	// hook：分类保存
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
	logger.Traceln("xx")
	logger.Debugln("xx")
	logger.Infoln("xx")
	logger.Warningln("xx")
	logger.Errorln("xx")
}
