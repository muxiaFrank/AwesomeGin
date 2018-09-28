/*
@Time : 2018/9/12 下午4:03 
@Author : Muxia
@File : router
@Software: GoLand
*/
package router

import (
	"github.com/gin-gonic/gin"
	. "../apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/api/users", Users)

	router.GET("/api/apps", GetApplist)

	router.POST("/api/user", Store)

	router.PUT("/api/update/user/:id", Update)

	router.DELETE("/api/delete/user/:id", Destroy)

	router.GET("/api/json", Gojson)

	return router
}