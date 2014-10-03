package main

import (
	_ "github.com/artpar/hopin/routers"
	"github.com/astaxie/beego"
	"github.com/artpar/hopin/models"
	"github.com/astaxie/beego/orm"
	"os"
	"github.com/artpar/hopin/helper"
)

func main() {

	vars := helper.ProcessArguments(os.Args)
	beego.Info(vars)
	beego.RunMode = vars["profile"]
	beego.Info("Run mode is ", beego.RunMode)
	beego.ParseConfig()
	beego.Info("Logging at : ", helper.String("logfile"))
	models.Orm = orm.NewOrm()
	beego.SetLogger("file", helper.String("logfile"))
	beego.Run(helper.String("serverhost") + ":" + helper.String("serverport"))
}
