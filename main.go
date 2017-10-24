package main

import (
	_ "onlineshop/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8")
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()

}
