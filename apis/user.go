/*
@Time : 2018/9/12 下午3:56 
@Author : Muxia
@File : user
@Software: GoLand
*/
package apis

import (
	"github.com/gin-gonic/gin"
	model "../models"
	"net/http"
	"strconv"
	"fmt"
)

//列表数据
func Users(c *gin.Context) {
	var user model.User
	//user.Username = c.Request.FormValue("username")
	//user.Password = c.Request.FormValue("password")
	result, err := user.Users()

	//UA := c.Request.UserAgent()
	//headers:=c.Request.Header

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		//"UA":UA,
		//"headers":headers,
		"code": 1,
		"data": result,
	})
}

//列表数据
func GetApplist(c *gin.Context) {
	var apps model.Applist
	//user.Username = c.Request.FormValue("username")
	//user.Password = c.Request.FormValue("password")
	result, err := apps.Applists()
	fmt.Print(result)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{

		"code": 1,
		"data": result,
	})
}

//添加数据
func Store(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	id, err := user.Insert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"http_code": http.StatusOK,
			"code":      -1,
			"message":   "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
		"data":    id,
	})
}

//修改数据
func Update(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "修改成功",
	})
}

//删除数据
func Destroy(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "删除成功",
	})
}

func Gojson(c *gin. Context){

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "删除成功",
		"ddd": 0,
		"ind": 1,
		"ty": 4,
	})

}