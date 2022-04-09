/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-07 00:29:41
 * @LastEditTime: 2022-04-10 01:16:21
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/models/user.go
 */
package models

// 相较于 dao.User，删除了一些不必要的信息
type UserInfo struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`

	// 下面是非必需
	Email   string `json:"email"`
	Address string `json:"address"`
	Gender  int8   `json:"gender"`
	Avatar  string `json:"avatar"`
}
