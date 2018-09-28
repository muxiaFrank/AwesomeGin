/*
@Time : 2018/9/12 下午4:03 
@Author : Muxia
@File : models
@Software: GoLand
*/
package models

import (
	orm "../database"
	)

type User struct {
	ID       int64  `json:"id"`       // 列名为 `id`
	Username string `json:"username"` // 列名为 `username`
	Password string `json:"password"` // 列名为 `password`
}

type Os_version struct {
	ID      int64  `json:"id"` // 列名为 `id`
	HEAD_OS string `json:"head_os"`
}

type Applist struct {
	//ID      int64  `json:"id"`
	APPID   int  `json:"appid"`
	APPNAME string `json:"appname"`
}

type App_content struct {
	CATEGORY_ID int64  `json:"category_id"`
	HEAD_OS     string `json:"head_os"`
	APPNAME     string `json:"appname"`
}

//var Users []User
//var Applists []Applist

//添加
func (user User) Insert() (id int64, err error) {

	//添加数据
	result := orm.Eloquent.Create(&user)

	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//列表
func (user *User) Users() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

//列表
func (apps *Applist) Applists() (applists []Applist, err error) {
	//orm.Eloquent.AutoMigrate(&applists)
	//orm.Eloquent.Find(&applists)
	if err = orm.Eloquent.Find(&applists).Error; err != nil {
		return
	}
	return
}

//修改
func (user *User) Update(id int64) (updateUser User, err error) {

	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

//删除数据
func (user *User) Destroy(id int64) (Result User, err error) {

	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}
