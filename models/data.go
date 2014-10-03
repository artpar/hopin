package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Orm orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql",
									beego.AppConfig.String("mysqluser")+":" +
										beego.AppConfig.String("mysqlpass")+"@" +
								beego.AppConfig.String("mysqlurl")+"/"+beego.AppConfig.String("mysqldb")+"?charset=utf8")
	beego.Info("Initiated Databse")
}

