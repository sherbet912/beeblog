package models

import (
	"fmt"
	"time"
	"github.com/beego/beego/v2/client/orm"
)

// 定義 user model, 綁定 users 表結構, 用來保存 sql 查詢結果
type User struct {
	UserID       int       `orm:"column(user_id);pk;auto"`     // auto 自增
    Username     string    `orm:"column(username);unique"`     // 唯一約束
    Password     string    `orm:"column(password)"`
    CreationDate time.Time `orm:"column(creation_date);type(datetime)"`
    ModifiedDate time.Time `orm:"column(modified_date);type(datetime)"`
}

// 定義 user model, 對應到哪張表
func (u *User) TableName() string {
	return "users"
}

// 初始化函數, 可用來向 orm 註冊 model
func init() {
	// 向 orm 註冊 user model
    orm.RegisterModel(&User{})
}

// 透過 user_id 查詢用戶訊息
func GetUserById(user_id int) *User {
	if user_id == 0 {
		return nil
	}

	// 建立 orm, 後面是透過 orm 操作 DB
	o := orm.NewOrm()
	
	user := User{}
	// 設定查詢參數
	user.UserID = user_id
	fmt.Println(user)

	// 使用 Read function, 根據 user 設定的參數來查詢, 將結果存到 user 中
	// 等價 sql: SELECT `user_id`, `username`, `password`, `creation_date`, `modified_date` FROM `users` WHERE `user_id` = user_id
	err := o.Read(&user)

	// 檢查結果
	if err == orm.ErrNoRows {
		// not found
		return nil
	} else if err == orm.ErrMissPK {
		// not found
        return nil
	}

	return &user
}