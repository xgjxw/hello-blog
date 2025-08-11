package hlog

import (
	"fmt"
	"log"
	"runtime"
)

// Debug 支持可变参数的日志函数，能打印当前代码行数
func Debug(v ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		msg := fmt.Sprint(v...)
		log.Printf("%s:%d: %s", file, line, msg)
	} else {
		log.Print(v...)
	}
}

// Debugf 支持格式化字符串的日志函数，能打印当前代码行数
func Debugf(format string, v ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.Printf("%s:%d: %s", file, line, fmt.Sprintf(format, v...))
	} else {
		log.Printf(format, v...)
	}
}

// Panic 支持可变参数的panic函数，能打印当前代码行数
func Panic(v ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		msg := fmt.Sprint(v...)
		log.Panicf("%s:%d: %s", file, line, msg)
	} else {
		log.Panic(v...)
	}
}

// Panicf 支持格式化字符串的panic函数，能打印当前代码行数
func Panicf(format string, v ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.Panicf("%s:%d: %s", file, line, fmt.Sprintf(format, v...))
	} else {
		log.Panicf(format, v...)
	}
}
