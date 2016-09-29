package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"db-beego/models"
	_ "github.com/go-sql-driver/mysql"
	_ "db-beego/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:password@tcp(127.0.0.1:3636)/go")
	orm.RegisterModel(new(models.Article))
}

func main() {
	beego.Run()
}
