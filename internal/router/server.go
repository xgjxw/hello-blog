package router

import (
	"context"
	"hello-blog/hlog"
	"log"
	"net/http"
)

type ApiHandler struct {
	server *http.Server
	addr   string
}

func NewApiHandler(addr string) *ApiHandler {
	return &ApiHandler{server: &http.Server{Addr: addr}, addr: addr}
}

func (api *ApiHandler) Start() {
	//加载路由
	LoadHandler()
	//打印路由
	PrintRoutes()
	//启动服务
	hlog.Debug("Server starting on " + api.addr)
	if err := api.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		hlog.Panic("Server failed to start: ", err)
	}
}

func (api *ApiHandler) Stop(ctx context.Context) {
	log.Println("Api server stop!!!")
	if err := api.server.Shutdown(ctx); err != nil {
		hlog.Panic("Server forced to shutdown: ", err)
	}
}

func (api *ApiHandler) Server() {
	api.Start()
}
