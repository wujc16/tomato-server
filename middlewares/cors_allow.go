/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-10 01:12:01
 * @LastEditTime: 2022-04-10 01:12:01
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/middlewares/cors_allow.go
 */
package middlewares

import "github.com/gin-gonic/gin"

func CorsAllowedMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
