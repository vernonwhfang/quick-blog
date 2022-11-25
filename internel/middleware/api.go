package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// ApiStart 应该放置作为第一个拦截器，完成所有相关处理链的准备工作
func ApiStart() gin.HandlerFunc {
	fmt.Println("middleware ApiStart ")
	return nil
}
