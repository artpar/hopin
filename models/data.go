package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/artpar/hopin/helper"
)

var Orm orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql",
									helper.String("mysqluser")+":" +
										helper.String("mysqlpass")+"@" +
								helper.String("mysqlurl")+"/"+helper.String("mysqldb")+"?charset=utf8")
	beego.Info("Initiated Databse")
}

