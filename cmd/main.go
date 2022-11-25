package main

import (
	"github.com/gin-gonic/gin"
)

//var conf = flag.String("config", "./conf/app.toml", "配置文件路径")

func main() {
	//flag.Parse()
	// 加载配置文件
	// 日志引入
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 在 0.0.0.0:8080 上监听并服务

	// run http server
	//fmt.Println("aaaaa")
	//r := gin.New()
	//router.Register(r)
	//fmt.Println("vvvv")
	//server.Run(r)
}
