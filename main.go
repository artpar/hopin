package main

import (
	_ "github.com/artpar/hopin/routers"
	"github.com/astaxie/beego"
	"github.com/artpar/hopin/models"
	"github.com/astaxie/beego/orm"
)

func main() {
	beego.Info("Logging at : ", beego.AppConfig.String("logfile"))
	beego.SetLogger("file", beego.AppConfig.String("logfile"))
	models.Orm = orm.NewOrm()
	beego.Run(beego.AppConfig.String("serverhost") + ":" + beego.AppConfig.String("serverport"))
}

