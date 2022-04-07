/*
 * @Author: jinchao.wu@bytedance.com
 * @Date: 2022-04-07 00:09:20
 * @LastEditTime: 2022-04-08 01:08:33
 * @LastEditors: jinchao.wu@bytedance.com
 * @Description:
 * @FilePath: /tomato-server/dal/dao/dao.go
 */
package dao

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error

	db, err = gorm.Open(mysql.New(mysql.Config{
		// 这里的 mysql host 需要使用 mysql 的 container 名称，后续需要改进为使用 ENV 提供
		DSN: "root:tomato_mysql_1231@tcp(tomato_mysql:3306)/tomato_db?charset=utf8mb4&parseTime=True&loc=Local",
	}))

	if err != nil {
		log.Fatal(2, "open mysql error"+err.Error())
	}
}
