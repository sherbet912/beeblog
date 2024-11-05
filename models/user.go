package models

import (
	"time"
	"github.com/beego/beego/v2/client/orm"
)

// 定義 user model, 綁定 user 表結構, 用來保存 sql 查詢結果
type User struct {
	Id           int64
    Account      string    `orm:"unique;size(16);index"`     // 唯一約束
    Password     string    `orm:"size(32)"`
    Nickname     string    `orm:"size(16);index"`
    Email        string    `orm:"size(255);index"`
    Avator       string    `orm:"size(255)"`
    LastLogin    time.Time `orm:"type(datetime);auto_now_add;index"`
    LastIp       string    `orm:"size(32);index"`
    LoginCount   int64     `orm:"index"`
    CreationDate time.Time `orm:"type(datetime);auto_now_add;index"`
    ModifiedDate time.Time `orm:"type(datetime);auto_now;index"`
}

// user json model
type UserJson struct {
    Id           int64     `json:"id"`
    Account      string    `json:"account"`
    Password     string    `json:"password"`
    Nickname     string    `json:"nickname"`
    Email        string    `json:"email"`
    Avator       string    `json:"avator"`
    LastLogin    time.Time `json:"last_login"`
    LastIp       string    `json:"last_ip"`
    LoginCount   int64     `json:"login_count"`
    CreationDate time.Time `json:"creation_date"`
}

// 定義 user model, 對應到哪張表
func (m *User) TableName() string {
	return "user"
}

func (m *User) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(m)
	return id, err
}

func (m *User) Update(fields ...string) (int64, error) {
	id, err := orm.NewOrm().Update(m, fields...)
	return id, err
}

func (m *User) Delete() error {
    if _, err := orm.NewOrm().Delete(m); err != nil {
        return err
    }

    return nil
}

func (m *User) Query() orm.QuerySeter {
    return orm.NewOrm().QueryTable(m)
}

// 透過 user_id 查詢用戶訊息
func GetUserById(id int64) *User {
	if id == 0 {
		return nil
	}

	// 建立 orm, 後面是透過 orm 操作 DB
	o := orm.NewOrm()
	
	user := User{}
	// 設定查詢參數
	user.Id = id

	// 使用 Read function, 根據 user 設定的參數來查詢, 將結果存到 user 中
	// 等價 sql: SELECT * FROM `users` WHERE `user_id` = user_id
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