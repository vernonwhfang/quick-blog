package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Controller
}

func (c *LoginController) Register(r gin.IRouter) {
	r.POST("", c.Login)
}

func (c *LoginController) Middle(_ *gin.Context) {}

// Login @title 登录
func (c *LoginController) Login(g *gin.Context) {
	// 校验用户名、密码
	// 将登录标识设置到上下文中
	//c.RespJSON(g, "登录成功！", nil)
	fmt.Println("login")
	return
}
