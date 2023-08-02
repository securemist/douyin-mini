package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/util"
	"net/http"
	"time"
)

/*
*
自定义中间件
截取请求参数中的token，解析token得到用户id放到请求的上下文中
其中feed接口是可以没有token的，会给一个id为0的默认用户
*/
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		var response resp.Response
		var claims *util.Claims
		var err error
		response = constant.AUTH_SUCCESS

		token := c.Query("token")
		if token == "" {
			if c.FullPath() != "/douyin/feed/" {
				response = constant.AUTH_FAILED
			} else {
				// feed是的token是可选参数，如果请求中没有token，就给定默认的游客用户
				c.Set("userId", int64(0))
				c.Next()
				return
			}
		} else {
			claims, err = util.ParseToken(token)
			if err != nil { // token解析失败
				response = constant.AUTH_FAILED
			} else if time.Now().Unix() > claims.ExpiresAt { // token过期
				response = constant.AUTH_EXPIRE
			}
		}

		// 权限验证失败
		if response != constant.AUTH_SUCCESS {
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// 将解析出的userId放进Context中
		c.Set("userId", claims.Id)
		c.Next()
	}
}
