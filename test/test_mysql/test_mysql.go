package main

import (
	"github.com/beego/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int    `orm:"column(id)"`
	Name     string `orm:"column(name)"`
	Password string `orm:"column(password)"`
}

func init() {
	// register model
	orm.RegisterModel(new(User))

	// register default database
	// modify username, password and database for yourself
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(172.17.0.2:3306)/atc?charset=utf8")
}

func main() {
	orm.Debug = true

	// automatically build table
	orm.RunSyncdb("default", false, true)

	// create orm object
	o := orm.NewOrm()

	// data
	user := new(User)
	user.Name = "alice"

	// insert data
	o.Insert(user)
}
