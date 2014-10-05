package controllers

import (
	"github.com/artpar/hopin/models"
	"github.com/astaxie/beego"
)

type TravellerController struct {
	beego.Controller
}

func (this *TravellerController) Get() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Data["json"] = "invalid id"

	}else {
		this.Data["json"] = models.GetTravellerById(id)
	}
	this.ServeJson()

}

func (this *TravellerController) Post() {

	people, err := this.GetInt("people")
	if err != nil {
		beego.Info("Bad integer for people", this.GetString("people"))
		this.Data["json"] = "Invalid number of people"
		this.ServeJson()
	}

	user := models.GetUserByEmail(this.GetString("email"))
	if user.RegId != this.GetString("regid") {
		this.Data["json"] = "Invalid Registration Id"
		this.ServeJson()
		return;
	}

	traveller := models.Traveller{
		From: this.GetString("from"),
		To: this.GetString("to"),
		FromPlaceId: this.GetString("fromPlaceId"),
		ToPlaceId: this.GetString("toPlaceId"),
		UserId: user.Id,
		StartTime: this.GetString("startTime"),
		EndTime: this.GetString("endTime"),
		People: people,
	}



	fromPlace := models.GooglePlaceDetailApi(traveller.FromPlaceId)
	traveller.FromLat = fromPlace.Lat
	traveller.FromLon = fromPlace.Lng

	toPlace := models.GooglePlaceDetailApi(traveller.ToPlaceId)
	traveller.ToLat = toPlace.Lat
	traveller.ToLon = toPlace.Lng

	beego.Info("Traveller ", traveller)
	traveller = models.AddTraveller(traveller)
	this.Data["json"] = traveller
	this.ServeJson()
}
