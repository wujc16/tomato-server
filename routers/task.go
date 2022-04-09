/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-09 23:44:16
 * @LastEditTime: 2022-04-10 00:06:13
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/routers/task.go
 */
package routers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wujc16/tomato-server/services"
	"github.com/wujc16/tomato-server/utils"
)

func getTaskList(c *gin.Context) {
	tasks, err := services.ListUserTasks(1, 20, 0)
	if err != nil {
		utils.HandleRespError(c, 4001, err.Error())
		return
	}

	utils.HandleRespOk(c, tasks)
}

func getTaskDetail(c *gin.Context) {
	taskId, qer := strconv.ParseInt(c.Query("task_id"), 10, 64)

	if qer != nil {
		utils.HandleRespError(c, 1000, "invalid query task_id!")
		return
	}

	task, err := services.GetTaskInfoById(taskId)
	if err != nil {
		utils.HandleRespError(c, 500, err.Error())
		return
	}

	utils.HandleRespOk(c, task)
}

func registerTaskRouter(rootRouter *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	taskSubRouter := rootRouter.Group("/task")
	{
		taskSubRouter.Use(middlewares...)
		taskSubRouter.GET("/list", getTaskList)
		taskSubRouter.GET("/detail", getTaskDetail)
	}
}
