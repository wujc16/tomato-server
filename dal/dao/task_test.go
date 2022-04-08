/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-09 00:42:47
 * @LastEditTime: 2022-04-09 01:02:00
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/dal/dao/task_test.go
 */
package dao_test

import (
	"testing"
	"time"

	"github.com/wujc16/tomato-server/dal"
	"github.com/wujc16/tomato-server/dal/dao"
)

func TestInsertTask(t *testing.T) {
	dal.Init()
	err := dao.InsertTask(&dao.Task{
		Uid:       1,
		TaskTitle: "测试任务 2",

		BeginAt:    time.Now(),
		DueAt:      time.Now().Add(24 * time.Hour),
		FinishedAt: time.Now(),
	})
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
}

func TestGetUserTasks(t *testing.T) {
	dal.Init()
	tasks, err := dao.SelectUidTasks(1, 20, 0)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
		return
	}

	t.Logf("length: %d", len(tasks))
}

func TestDelTask(t *testing.T) {
	dal.Init()
	err := dao.DeleteTask(&dao.Task{
		Id: 1,
	})

	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
}
