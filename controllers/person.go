package controllers

import (
	"github.com/astaxie/beego"
	"github.com/artpar/hopin/models"
)

type PersonController struct {
	beego.Controller
}

type PersonFeed struct {
	Travels []models.Travel
}

func (this *PersonController) Get() {
	this.Data["json"] = models.GetUserByEmail(this.GetString("email"))
	this.ServeJson()
}

// Todo: implement the method to get a single travel of a user by id
// Todo: Implement the table to store user request for other travels and accept and reject status

func (this *PersonController) Feed() {
	email := this.GetString("email")

	user := models.GetUserByEmail(email)
	regid := this.GetString("regid")
	if user.RegId != regid {
		this.Data["json"] = "not a proper request"
		this.ServeJson()
	}
	var feed PersonFeed;

	feed.Travels = models.GetUserTravels(user)
	this.Data["json"] = feed
	this.ServeJson()
}

func (this *PersonController) Post() {
	email := this.GetString("email")
	p := models.GetUserByEmail(email)

	if p.Id == 0 {
		p = models.Person{Email:email}
		p = models.CreateUser(p)
	}

	this.Data["json"] = p
	this.ServeJson()
}

func (this *PersonController) Put() {
	email := this.GetString("email")
	regId := this.GetString("regid")

	person := models.GetUserByEmail(email)
	if len(person.RegId) > 0 && len(person.RegId) > 0 {
		this.Data["json"] = "already exists"
		this.ServeJson()
		return;
	}
	person.RegId = regId

	models.UpdateUser(person)
	this.Data["json"] = person
	this.ServeJson()
}
