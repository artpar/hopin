package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/artpar/hopin/helper"
)

var Orm orm.Ormer
const TimeFormat = "Jan 2, 2006 at 3:04pm (MST)"

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql",
									helper.String("mysqluser")+":" +
										helper.String("mysqlpass")+"@" +
								helper.String("mysqlurl")+"/"+helper.String("mysqldb")+"?charset=utf8&parseTime=true")
	beego.Info("Initiated Databse")
}

