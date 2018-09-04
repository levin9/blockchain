/**********************************************
** @Des: This file ...
** @Author: lizg
** @Date:   2018-08-30 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-24 11:48:17
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	BaseEntity
	Id         int
	LoginNo    string `orm:"size(20)"`
	NickName   string `orm:"size(20)"`
	Password   string `orm:"size(50)"`
	IconUrl    string `orm:"size(50)"`
	Sexy       int
	Mobile     string `orm:"size(20)"`
	Email      string `orm:"size(50)"`
	UserStatus int
	SortNum    int
}

func (a *User) TableName() string {

	return TableName("uc_member")
}

func UsersAdd(a *User) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func UserGetList(page, pageSize int, filters ...interface{}) ([]*User, int64) {
	offset := (page - 1) * pageSize
	list := make([]*User, 0)
	query := orm.NewOrm().QueryTable(TableName("uc_member"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func UserGetByLoginNo(loginno string) (*User, error) {
	u := new(User)
	//error userstatus=?
	err := orm.NewOrm().QueryTable(TableName("uc_member")).Filter("LoginNo", loginno).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserGetById(id int) (*User, error) {
	u := new(User)
	//error userstatus=?
	err := orm.NewOrm().QueryTable(TableName("uc_member")).Filter("Id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (a *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
