package model

import (
	"gin/common"
	"time"
)

var user *User

type User struct {
	Id        int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	UserName  string    `gorm:"type:varchar(255);unique_index;not null" json:"user_name"`
	PassWord  string    `gorm:"type:varchar(255);not null" json:"pass_word"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserOto struct {
	Id        int
	UserName  string
	CreatedAt time.Time
}

func init() {
	db := common.GetDb()

	if !db.HasTable(&User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	}
	user = new(User)
}

func GetUserInstance() *User {
	return user
}

//查找
func (user *User) Find() *User {
	userResult := new(User)
	common.GetDb().Where(user).First(&userResult)
	return userResult
}

//add
func (user *User) Add() {
	common.GetDb().Create(user)
}

//update
func (user *User) Update(updateUser User) User {
	common.GetDb().Model(user).Updates(updateUser)
	return updateUser
}

//del
func (user *User) Del() {
	common.GetDb().Delete(user)
}

//搜索查找
func (user *User) GetUserInfoArr(username string) []*User {
	var userResult []*User
	db := common.GetDb()
	if username != "" {

		db = db.Where("user_name LIKE ?", "%"+username+"%")
	}
	db.Find(&userResult)
	return userResult
}
