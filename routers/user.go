/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-06 23:29:27
 * @LastEditTime: 2022-04-09 23:41:30
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/routers/user.go
 */
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wujc16/tomato-server/services"
	"github.com/wujc16/tomato-server/utils"
)

func getUserInfo(c *gin.Context) {
	user, err := services.GetUserInfoById(1)
	if err != nil {
		utils.HandleRespError(c, 4001, "user not found")
		return
	}
	utils.HandleRespOk(c, user)
}

func registerUserRouter(rootRouter *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	userRouter := rootRouter.Group("/user")
	{
		userRouter.Use(middlewares...)
		userRouter.GET("/profile", getUserInfo)
	}
}
