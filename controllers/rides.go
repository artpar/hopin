package controllers

import (
	"github.com/artpar/hopin/models"
	"github.com/astaxie/beego")

type RideController struct {
	beego.Controller
}

func (this *RideController) Get() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Data["json"] = "invalid id"

	}else {
		this.Data["json"] = models.GetRideById(id)
	}
	this.ServeJson()

}

func (this *RideController) Post() {

	people, err := this.GetInt("people")
	if err != nil {
		beego.Info("Bad integer for people", this.GetString("people"))
		this.Data["json"] = "Invalid number of people"
		this.ServeJson()
		return;
	}

	user := models.GetUserByEmail(this.GetString("email"))
	if user.RegId != this.GetString("regid") {
		this.Data["json"] = "Invalid Registration Id"
		this.ServeJson()
		return;
	}

	ride := models.Ride{
		From: this.GetString("from"),
		To: this.GetString("to"),
		FromPlaceId: this.GetString("fromPlaceId"),
		ToPlaceId: this.GetString("toPlaceId"),
		UserId: user.Id,
		StartTime: this.GetString("startTime"),
		EndTime: this.GetString("endTime"),
		People: people,
	}
	fromPlace := models.GooglePlaceDetailApi(ride.FromPlaceId)
	ride.FromLat = fromPlace.Lat
	ride.FromLon = fromPlace.Lng

	toPlace := models.GooglePlaceDetailApi(ride.ToPlaceId)
	ride.ToLat = toPlace.Lat
	ride.ToLon = toPlace.Lng


	beego.Info("Ride ", ride)
	ride = models.AddRide(ride)
	this.Data["json"] = ride
	this.ServeJson()
}
