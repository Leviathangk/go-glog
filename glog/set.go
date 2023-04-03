package glog

// SetTimeFormat 设置时间格式化的格式
func (logger *Logger) SetTimeFormat(format string) {
	logger.Config.Formatter.TimeFormat = format
}
