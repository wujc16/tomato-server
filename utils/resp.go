/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-07 00:42:27
 * @LastEditTime: 2022-04-07 00:48:58
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/utils/resp.go
 */
package utils

import "github.com/gin-gonic/gin"

func HandleRespOk(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"status":  0,
		"message": "success",
		"data":    data,
	})
}

func HandleRespError(c *gin.Context, errorCode int16, errorMsg string) {
	c.JSON(200, gin.H{
		"status":  errorCode,
		"message": errorMsg,
	})
}
