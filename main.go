/*
@Time : 2018/9/12 下午4:04 
@Author : Muxia
@File : main
@Software: GoLand
*/
package main


import (
	_ "./database"
	"./routers"
	orm "./database"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":23333")
}