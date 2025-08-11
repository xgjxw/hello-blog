package router

import (
	"encoding/json"
	"hello-blog/config"
	"net/http"
)

// BlogInfoResponse 定义返回信息的结构体
type BlogInfoResponse struct {
	Version string `json:"version"`
	Message string `json:"message"`
}

// GetBlogInfo 处理 /blog/info 请求
func GetBlogInfo(w http.ResponseWriter, r *http.Request) {
	// 设置响应头为 JSON 格式
	w.Header().Set("Content-Type", "application/json")
	
	// 创建响应数据
	response := BlogInfoResponse{
		Version: config.BlogInfo.Version,
		Message: config.BlogInfo.Message,
	}
	
	// 将数据编码为 JSON 并返回
	json.NewEncoder(w).Encode(response)
}