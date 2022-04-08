/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-09 01:03:56
 * @LastEditTime: 2022-04-09 01:07:11
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/services/taskserv.go
 */
package services

import (
	"github.com/wujc16/tomato-server/dal/dao"
	"github.com/wujc16/tomato-server/models"
)

func GetTaskInfoById(id int64) (models.Task, error) {
	task, err := dao.SelectTaskById(id)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func ListUserTasks(uid int64, count int, offset int) ([]models.Task, error) {
	tasks, err := dao.SelectUidTasks(uid, count, offset)
	if err != nil {
		return make([]dao.Task, 0), err
	}

	return tasks, nil
}
