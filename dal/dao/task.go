/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-09 00:04:27
 * @LastEditTime: 2022-04-10 01:18:01
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/dal/dao/task.go
 */
package dao

import (
	"time"
)

type Task struct {
	Id  int64 `json:"id"`
	Uid int64 `json:"uid"`

	TaskTitle       string `json:"task_title"`       // 任务名称
	TaskDescription string `json:"task_description"` // 任务描述

	Priority int8 `json:"priority"` // 重要程度 0，1，2，3，4，5
	Urgency  int8 `json:"urgency"`  // 紧急程度 0，1，2，3，4，5
	Status   int8 `json:"status"`   // 任务状态 0：初始化，1：已完成

	TaskScore   int8   `json:"task_score"` // 总结评分，0，1，2，3，4，5
	TaskSummary string `json:"task_summary"`

	BeginAt    time.Time `json:"begin_at"`    // 开始时间
	DueAt      time.Time `json:"due_at"`      // 截止日期
	FinishedAt time.Time `json:"finished_at"` // 实际完成时间

	// 其他时间信息
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (task *Task) TableName() string {
	return "busi_task"
}

// 创建任务
func InsertTask(t *Task) error {
	result := db.Save(t)
	return result.Error
}

// 通过 id 查询任务详情信息
func SelectTaskById(id int64) (task Task, err error) {
	result := db.First(&task, id)
	return task, result.Error
}

// 获取用户的任务列表
func SelectUidTasks(uid int64, count int, offset int) (tasks []Task, err error) {
	result := db.Offset(offset).Limit(count).Where("uid = ?", uid).Find(&tasks)
	return tasks, result.Error
}

// 删除任务
func DeleteTask(t *Task) error {
	result := db.Delete(t)
	return result.Error
}
