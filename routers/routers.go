/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-06 23:28:43
 * @LastEditTime: 2022-04-09 23:46:14
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/routers/routers.go
 */
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wujc16/tomato-server/middlewares"
)

func RegisterRoutersAndMiddlewares(r *gin.Engine) {
	rootRouter := r.Group("/api")
	{
		// 登陆登出
		registerAuthRouter(rootRouter)
		// 用户信息
		registerUserRouter(rootRouter, middlewares.UserAuthMiddleware)
		registerTaskRouter(rootRouter, middlewares.UserAuthMiddleware)
	}
}
