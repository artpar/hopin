package controllers

import (
	"github.com/astaxie/beego"
	"github.com/artpar/hopin/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	beego.Info("test insert")
	models.Test()
	this.Data["json"] = new(models.Person)
	this.ServeJson()

	//	this.TplNames = "index.tpl"
	//	beego.Info(beego.AppConfig.String("logfile"))
}
