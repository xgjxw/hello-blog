package middleware

import (
	"fmt"
	"hello-blog/internal/utils"
	"net/http"
	"time"
)

// loggingResponseWriter 自定义ResponseWriter，用于捕获状态码
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// newLoggingResponseWriter 创建新的loggingResponseWriter实例
func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

// WriteHeader 重写WriteHeader方法以捕获状态码
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware 日志中间件
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 创建自定义ResponseWriter来捕获状态码
		lrw := newLoggingResponseWriter(w)

		// 调用下一个处理函数
		next.ServeHTTP(lrw, r)

		// 获取客户端IP地址
		clientIP := utils.GetClientIP(r)

		// 计算处理时间
		duration := time.Since(start)

		// 打印日志信息
		fmt.Printf("[%s] %s %s from %s - Status: %d - Duration: %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			clientIP,
			lrw.statusCode,
			duration,
		)
	})
}
