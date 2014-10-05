package controllers

import (
	"github.com/astaxie/beego"
	"github.com/artpar/hopin/helper"
)

type PlacesController struct {
	beego.Controller
}

func (this *PlacesController) Get() {
	requestParamMap := helper.MapArrayToMapString((map[string][]string)(this.Input()))
	data := helper.GooglePlacesAutocompleteApi(requestParamMap)
	beego.Info("response", data)
	this.Data["json"] = data;
	this.ServeJson()
}
