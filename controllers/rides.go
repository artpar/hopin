package controllers

import (
	"github.com/artpar/hopin/models"
	"github.com/astaxie/beego"
)

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
	ride := Ride{
		From: this.GetString("from"),
		To: this.GetString("from"),
		UserId: models.GetUserByEmail(this.GetString("email")).Id,
		StartTime: this.GetString("from"),
		EndTime: this.GetString("from"),
		Capacity: this.GetString("from"),
	}
	ride = models.AddRide(ride)
	this.Data["json"] = ride
	this.ServeJson()
}
