package log

import (
	"testing"
)

func Test(t *testing.T) {
	Debug("这是一条Debug信息")
	Info("这是一条普通信息")
	Warn("这是一条警告信息")
	Error("这是一条错误信息")
}
