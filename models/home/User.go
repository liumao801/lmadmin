package home

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id 			int
	Username 	string	`orm:size(32)`
	Passwd 		string 	`orm:size(32)`
	Name 		string 	`orm:size(32)`
	Status 		int8
	CreatedAt	int
	UpdatedAt 	int
}

type UserQueryParam struct {
	HomeBaseQueryParam
	UsernameLike 	string
}

func (m *User) TableName() string {
	return UserTBName()
}

// 根据用户名和密码查询用户信息
func UserOneByUsername(username string) (*User, error) {
	m := User{}
	err := orm.NewOrm().QueryTable(UserTBName()).Filter("username", username).One(&m)
	if err != nil {
		return nil, err
	}
	beego.Info("查询成功", m)
	return &m, nil
}

// 添加用户
func UserAdd(m *User) (*User, error) {
	// 初始化 orm 对象
	o := orm.NewOrm()

	m.Status = 1
	m.CreatedAt = int(time.Now().Unix())
	m.UpdatedAt = m.CreatedAt

	_, err := o.Insert(m)
	if err != nil {
		return nil, err
	}
	beego.Info("插入数据成功", m)

	return m, nil
}

// 更新用户信息
func UserUpdate(Id int, m *User) (*User, error) {
	o := orm.NewOrm()
	me := User{}
	me.Id = Id
	err := o.Read(&me)
	if err != nil {
		return nil, err
	}
	if m.Username != me.Username {
		return nil, errors.New("用户名一经注册不能修改")
	}

	m.Id = me.Id
	if _, err := o.Update(m); err != nil {
		return nil, errors.New("更新用户信息失败")
	}

	return m, nil
}