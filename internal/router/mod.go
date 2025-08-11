package router

import (
	"fmt"
	"hello-blog/internal/middleware"
	"net/http"
)

// Route 定义路由结构
type Route struct {
	Method string
	Path   string
	Func   http.HandlerFunc
}

// Routes 路由映射表
var Routes = map[string]Route{
	"blog_info": {
		Method: "GET",
		Path:   "/blog/info",
		Func:   GetBlogInfo,
	},
}

// LoadHandler 加载路由处理器
func LoadHandler() {
	// 创建默认的ServeMux
	mux := http.NewServeMux()

	// 注册路由到自定义的ServeMux
	for _, route := range Routes {
		mux.HandleFunc(route.Path, route.Func)
	}

	// 使用中间件包装mux
	http.Handle("/", middleware.LoggingMiddleware(mux))
}

// PrintRoutes 打印所有路由信息
func PrintRoutes() {
	fmt.Println("Registered Routes:")
	for name, route := range Routes {
		fmt.Printf("%s\t%s\t\t%s\n", route.Method, route.Path, name)
	}
}
