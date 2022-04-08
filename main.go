/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-06 23:23:25
 * @LastEditTime: 2022-04-08 23:58:09
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description: main
 * @FilePath: /tomato-server/main.go
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wujc16/tomato-server/dal"
	"github.com/wujc16/tomato-server/routers"
)

func main() {
	// 首先初始化存储
	dal.Init()

	r := gin.Default()

	// 注册路由，使用 middleware 中间件
	routers.RegisterRoutersAndMiddlewares(r)

	r.Run(":8000")
}

// func main() {
// 	r := gin.Default()
// 	r.GET("/user/profile", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"name": "jinchao",
// 			"age":  28,
// 		})
// 	})

// 	r.Run(":8000")
// }
