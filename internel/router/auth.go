package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"quick-blog/internel/modules/auth"

	"net/http"
)

// UserAuth 鉴权
func UserAuth(g *gin.Context) {
	// 获取 token
	tokenStr := g.GetHeader(auth.TokenName)
	if len(tokenStr) == 0 {
		Unauthorized(g, errors.New("认证信息已失效 请重新登录"))
		return
	}

	// 解析 token
	_, err := auth.Parse(tokenStr)
	if err != nil {
		Unauthorized(g, err)
		return
	}

	/*	// 校验 token
		loginUser, err := models.ParseLoginUser(token.UserId)
		if err != nil {
			Unauthorized(g, err)
			return
		}

		// 刷新统一登录平台 token 有效期，每次增加 2 小时
		if g.FullPath() != PathUrlLogout {
			data, err := unify_req.RefreshTokenTimeOut(loginUser.PdToken)
			if err != nil {
				Unauthorized(g, log.LogError("认证信息已失效 请重新登录", err, loginUser.PdToken))
				return
			}
			if data == nil {
				Unauthorized(g, log.LogError("认证信息已失效 请重新登录", err))
				return
			}
		}*/

	// 在上下文中保存用户信息
	// g.Set(models.LoginUserParamKey, loginUser)
}

func Unauthorized(g *gin.Context, err error) {
	g.JSON(http.StatusOK, gin.H{
		"code":    http.StatusUnauthorized,
		"message": err.Error(),
	})
	g.Abort()
}
