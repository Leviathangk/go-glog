# glog

一个日志记录模块  
主要是简化使用

效果图：颜色可自定义

# 安装

![img.png](img.png)

```
go get -u github.com/Leviathangk/go-glog@latest
```

# 配置

注意：这里有两个级别，输出到控制台和输出到自定义是不同的级别，可以分开控制

```
// Logger 结构体
type Logger struct {
	ShowColor    bool        // 是否在输出的时候输出颜色（仅控制台）
	ShowCaller   bool        // 显示调用栈
	Print        bool        // 是否输出到控制台
	PrintLevel   int         // 日志限制输出级别（仅控制台）
	Handler      []io.Writer // 其它输出方式，受 OutPutLevel 限制
	HandlerLevel int         // 日志限制输出级别（其余输出）
	Hook         []HookFunc  // 钩子函数
	TimeFormat   string      // 时间格式化
}
```

# logger 方法

## AddHandler

给日志增加一个输出，受 OutPutLevel 限制  
注意：要实现 Writer 接口

```
file, _ := os.OpenFile("demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
logger.AddHandler(file)
```

## AddHook

利用 hook 实现分类处理  
需要实现的方法

```
type HookFunc func(level int, out string)
```

### 示例

假设现在的需求是只存 error 级别及以上的日志

```
// 新建文件
errFile, _ := os.OpenFile("err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

// 添加 Hook
logger.AddHook(func(level int, out string) {
    if level >= glog.ErrorLevel {
        _, err := errFile.Write([]byte(out))
        if err != nil {
            return
        }
    }
})
```