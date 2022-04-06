/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-07 00:11:58
 * @LastEditTime: 2022-04-07 00:20:55
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/dal/dao/user.go
 */
package dao

import "time"

type User struct {
	Id       int64
	Username string
	Password string
	Mobile   string

	// 下面是非必需
	Email     string
	Address   string
	Gender    int8
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 表名和 tomato.sql 用户表一致
func (user *User) TabelName() string {
	return "sys_user"
}

// 创建用户
func InsertUser(user *User) error {
	result := db.Save(user)
	return result.Error
}

// 通过 id 查询用户信息
func SelectUserById(id int64) (user User, err error) {
	result := db.First(&user, id)
	return user, result.Error
}

// 通过用户手机号查询用户信息
func SelectUserByMobile(mobile string) (user User, err error) {
	result := db.Where("mobile = ?", mobile).First(&user)
	return user, result.Error
}

// 更新用户信息
// func UpdateUser(user User) error {
// 	// result := db.Model(&user).Update(user)
// 	// return result.Error
// }
