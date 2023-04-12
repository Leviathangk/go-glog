// Package glog 配置设置方面的
package glog

import "io"

// SetTimeFormat 设置时间格式化的格式
func (logger *Logger) SetTimeFormat(format string) {
	logger.Config.Formatter.TimeFormat = format
}

// AddOutPut 增加输出：必须实现 Writer 接口
func (logger *Logger) AddOutPut(writer ...io.Writer) {
	logger.Config.Out = append(logger.Config.Out, writer...)
}

// AddHook 添加钩子函数
func (logger *Logger) AddHook(hookFunc ...HookFunc) {
	logger.Hook = append(logger.Hook, hookFunc...)
}
