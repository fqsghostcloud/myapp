package db

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id    int    `orm:"pk;column(user_id);auto;"`
	Name  string `orm:"column(user_name);"`
	Pwd   string `orm:"column(user_pwd);"`
	Intro string `orm:"null;column(user_intro);"`
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.34.36:3306)/myapp?charset=utf8", 30)
	orm.RunSyncdb("default", false, false)
}

func (u *User) TableEngine() string {
	return "INNODB"
}

func (u *User) EncodePwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	return string(hash)
}

func (u *User) CheckPwd(encodePwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePwd), []byte(pwd))
	if err == nil {
		return true
	}
	return false
}

func (u *User) AddUser(name, pwd, intro string) (id int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(User)
	user.Name = name
	user.Pwd = u.EncodePwd(pwd)
	user.Intro = intro
	return o.Insert(user)
}

func (u *User) CheckUser(name, pwd string) bool {
	o := orm.NewOrm()
	o.Using("default")
	res := make(orm.Params)
	_, err := o.Raw("select user_name, user_pwd from user where user_name = ?", name).RowsToMap(&res, "user_name", "user_pwd")
	if err != nil || res[name] == nil {
		fmt.Println(err)
		return false
	}
	return u.CheckPwd(res[name].(string), pwd)
}

func (u *User) GetIntro(name string) *User {
	o := orm.NewOrm()
	o.Using("default")
	res := make(orm.Params)
	_, err := o.Raw("select user_name, user_intro from user where user_name = ?", name).RowsToMap(&res, "user_name", "user_intro")
	if err != nil || res[name] == nil {
		fmt.Println(err)
		return nil
	}
	user := &User{Name: name, Intro: res[name].(string)}

	return user

}
