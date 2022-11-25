package router

import (
	"github.com/gin-gonic/gin"
	c "quick-blog/internel/controller"
	"quick-blog/internel/middleware"
)

const (
	PrefixVersion = "/v1/"
)

type IController interface {
	Register(r gin.IRouter)
}

type IPreHandle interface {
	Middle(g *gin.Context)
}

// 不需要鉴权的 controller 注册位置
var noAuthControllers = map[string]IController{
	"/login": new(c.LoginController),
}

// 需要鉴权的 controller 注册位置
var controllers = map[string]IController{
	"/major": new(c.MajorController),
}

func Register(r *gin.Engine) {
	// 注册全局中间件
	r.Use(gin.Recovery(), middleware.ApiStart())

	v1 := r.Group(PrefixVersion)

	for prefix, controller := range noAuthControllers {
		// 通过为指定的 controller 接口方法，单独定义中间件，来实现 controller 的预处理
		// 如果实现了自定义的预处理接口，会调用预处理方法
		handlers := make([]gin.HandlerFunc, 0, 1)
		if preHandle, ok := controller.(IPreHandle); ok {
			handlers = append(handlers, preHandle.Middle)
		}
		sg := v1.Group(prefix, handlers...)
		controller.Register(sg)
	}

	for prefix, controller := range controllers {
		// 需要鉴权的路由，插入鉴权中间件
		handlers := make([]gin.HandlerFunc, 0, 2)
		handlers = append(handlers, UserAuth)
		if preHandle, ok := controller.(IPreHandle); ok {
			handlers = append(handlers, preHandle.Middle)
		}
		sg := v1.Group(prefix, handlers...)
		controller.Register(sg)
	}
}
