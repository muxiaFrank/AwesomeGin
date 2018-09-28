/*
@Time : 2018/9/12 下午3:57 
@Author : Muxia
@File : mysql
@Software: GoLand
*/
package database

import (
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"fmt"
)

var Eloquent *gorm.DB

func init() {
	var err error

	//db配置信息
	user := "root"
	psw := "aweme"
	host := "localhost"
	port := 3306
	dbname := "DEMOMYSQL"
	charset := "utf8"
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", user, psw, host, port, dbname, charset)
	Eloquent, _ = gorm.Open("mysql", config)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}