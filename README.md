# glog

一个日志记录模块  
主要是简化使用

# 安装

```
go get github.com/Leviathangk/go-glog@latest
```

# 使用

使用默认的 Logger  
注意：直接使用 glog.Debugln 也是默认的 logger

```
import ""github.com/Leviathangk/go-glog/glog""

func main() {
    // 方式 1
    glog.Debugln("xx")              // 默认的 logger
    
    // 方式 2
    logger := glog.DefaultLogger    // 默认的 logger
    logger.Traceln("xx")
    logger.Debugln("xx")
    logger.Infoln("xx")
    logger.Warningln("xx")
    logger.Errorln("xx")
    logger.Fatalln("xx")
    logger.Panicln("xx")
}
```

# 保存

使用 AddOutPut 增加一个输出（要实现 Writer 方法），将会自动输出日志到文件

```
file, _ := os.OpenFile("demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
logger.AddOutPut(file)
```

# Hook

利用 hook 实现分类保存  
需要实现的方法

```
type HookFunc func(level int, out string)
```

示例

```
// 专门存 error 级别的日志
errFile, _ := os.OpenFile("err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

// 添加 Hook
logger.AddHook(func(level int, out string) {
    if level == glog.ErrorLevel {
        _, err := errFile.Write([]byte(out))
        if err != nil {
            return
        }
    }
})
```