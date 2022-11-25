package server

import "C"
import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//var C conf.Config

//func Init(c conf.Config) {
//	C = c
//}

// Run engine 应该是注册好 路由 和 中间件 的 gin 实例
//
// 优雅停机，超时时间设定为 10s
// 参考：https://github.com/gin-gonic/gin#graceful-shutdown-or-restart
func Run(e *gin.Engine) {
	// 获取路由信息
	// e.Routes()
	// 静态目录代理
	// e.Static("/static", "static")
	// 404 时应当作为接口响应的内容
	// e.NoRoute()
	fmt.Println("cccc")
	// 优雅停机
	srv := &http.Server{
		Addr:           ":8080", //  fmt.Sprintf(":%d", C.Port)
		Handler:        e,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		//log.Printf("[%s] 服务运行于 [%s] 模式，监听端口：%d", C.AppName, C.RunMode, C.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	//log.Println("关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		//log.Fatalf("关闭异常:%s", err)
	}
	//log.Println("关闭成功")
}
