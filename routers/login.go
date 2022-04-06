/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-06 23:46:21
 * @LastEditTime: 2022-04-06 23:50:04
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/routers/login.go
 */
package routers

import "github.com/gin-gonic/gin"

func postLogin(c *gin.Context) {
}

func postLogout(c *gin.Context) {
}

func registerAuthRouter(rootRouter *gin.RouterGroup) {
	rootRouter.POST("/login", postLogin)
	rootRouter.POST("/logout", postLogout)
}
