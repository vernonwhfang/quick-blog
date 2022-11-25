package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type MajorController struct {
	Controller
}

func (c MajorController) Register(r gin.IRouter) {
	// register api function here
	r.GET("", c.List)
}

func (c MajorController) Middle(g *gin.Context) {
	// init usercase here
}

// List @Summary 获取所有专业（包含本科和专科）
// @Description  获取所有专业（包含本科和专科）
// @Produce json
// @Success 200 {object} models.Major
// @Tags /major
// @Router /v1/major [get]
func (c MajorController) List(g *gin.Context) {
	fmt.Println("list")
	return
}
