package main

import (
	_ "guess/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:1314@tcp(127.0.0.1:3306)/beego?charset=utf8")

	orm.RegisterDataBase("default","mysql","root:1314@tcp(127.0.0.1:3306)/beego?charset=utf8")
}

func main() {
	beego.Run()
}

