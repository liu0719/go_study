package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

// 中间件文件

// gin的handler处理函数没有返回值，因此中间件只能通过Context来通信
const (
	ctxErrKey   = "app.err"   //错误
	ctxTraceKey = "app.trace" //调试
)

var debugMode = os.Getenv("DEBUG") == "1"

var traceSeq atomic.Uint64

func newTraceId() string {
	return fmt.Sprintf("%08x%04x",
time.Now().UnixNano()&0xffffffff,traceSeq.Add(1)&0xffff)
}
