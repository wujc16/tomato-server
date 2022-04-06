/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-07 00:29:41
 * @LastEditTime: 2022-04-07 00:31:50
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/models/user.go
 */
package models

// 相较于 dao.User，删除了一些不必要的信息
type UserInfo struct {
	Id       int64
	Username string
	Mobile   string

	// 下面是非必需
	Email   string
	Address string
	Gender  int8
	Avatar  string
}
