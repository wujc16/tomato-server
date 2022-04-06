/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-06 23:28:43
 * @LastEditTime: 2022-04-06 23:49:15
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/routers/routers.go
 */
package routers

import "github.com/gin-gonic/gin"

func RegisterRoutersAndMiddlewares(r *gin.Engine) {
	rootRouter := r.Group("/api")
	{
		// 登陆登出
		registerAuthRouter(rootRouter)
		// 用户信息
		registerUserRouter(rootRouter)
	}
}
