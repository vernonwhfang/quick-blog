package controller

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"quick-blog/internel/dao"
)

type Controller struct {
}

type IController interface {
	FormInt64Slice(g *gin.Context, key string, def ...[]int64) ([]int64, error)
	QueryInt64Slice(g *gin.Context, key string, def ...[]int64) ([]int64, error)
	FormInt64(g *gin.Context, key string) int64
	QueryInt(g *gin.Context, key string) int
	QueryInt64(g *gin.Context, key string) int64
	PathInt64(g *gin.Context, key string) int64
	PathInt(g *gin.Context, key string) int
	QueryBool(g *gin.Context, key string) bool
	GetFiles(g *gin.Context, key string) ([]*multipart.FileHeader, error)
	// Valid 校验 param 参数是否合法，不合法将直接返回错误信息
	// param 必须是结构体类型
	// 生产模式下，不将具体的错误信息作为接口响应返回
	Valid(g *gin.Context, param interface{}) bool
	// ReqParseAndValid 将请求体中的 json 参数反序列化到 param 中，并进行校验
	// param 必须是结构体指针类型
	// 生产模式下，不将具体的错误信息作为接口响应返回
	ReqParseAndValid(g *gin.Context, param interface{}) bool
	// RespVJSON V（Valid），响应校验不通过的 json 数据信息
	// 生产模式下，不将具体的错误信息作为接口响应返回
	RespVJSON(g *gin.Context, kvs ...string)
	// ResponseJSON 参数详尽的响应 json 数据的方法
	ResponseJSON(g *gin.Context, httpCode int, code dao.RespCodeEnum, data interface{}, err error)
	// RespJSON 最常用的响应 json 数据的方法
	RespJSON(g *gin.Context, data interface{}, err error)
	ServeFile(g *gin.Context, file multipart.File, fileName string, fileSize int64)
}
