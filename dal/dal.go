/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-07 00:08:56
 * @LastEditTime: 2022-04-07 00:21:35
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/dal/dal.go
 */
package dal

import "github.com/wujc16/tomato-server/dal/dao"

func Init() {
	dao.Init()
}
