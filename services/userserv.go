/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-07 00:24:17
 * @LastEditTime: 2022-04-07 00:39:02
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/services/userserv.go
 */
package services

import (
	"github.com/wujc16/tomato-server/dal/dao"
	"github.com/wujc16/tomato-server/models"
)

func GetUserInfoById(id int64) (models.UserInfo, error) {
	daoUser, err := dao.SelectUserById(id)

	if err != nil {
		return models.UserInfo{}, err
	}

	return models.UserInfo{
		Username: daoUser.Username,
		Mobile:   daoUser.Mobile,

		// 下面是非必需
		Email:   daoUser.Email,
		Address: daoUser.Address,
		Gender:  daoUser.Gender,
		Avatar:  daoUser.Avatar,
	}, nil
}

// 校验用户密码
func ValidateUserPasswordWithMobile(mobile string, password string) bool {
	daoUser, err := dao.SelectUserByMobile(mobile)
	if err != nil {
		return false
	}
	return password == daoUser.Password
}

// 校验用户 id 和密码是否一致
func ValidateUserPasswordWithId(id int64, password string) bool {
	daoUser, err := dao.SelectUserById(id)
	if err != nil {
		return false
	}
	return password == daoUser.Password
}

func ValidateUserPasswordWithUserName(userName string, password string) bool {
	return false
}

// func UpdateUserInfo
//
