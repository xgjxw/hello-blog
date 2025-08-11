package utils

import "net/http"

// getClientIP 获取客户端真实IP地址
func GetClientIP(r *http.Request) string {
	// 检查 X-Forwarded-For 头部（负载均衡或代理后面）
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded
	}

	// 检查 X-Real-IP 头部
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// 直接使用 RemoteAddr
	return r.RemoteAddr
}
