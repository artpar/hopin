package controllers

import (
	"github.com/astaxie/beego"
	"github.com/artpar/hopin/models"
)

type PersonController struct {
	beego.Controller
}

func (this *PersonController) Get() {
	this.Data["json"] = models.GetUserByEmail(this.GetString("email"))
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
	regId := this.GetString("reg_id")

	person := models.GetUserByEmail(email)
	person.RegId = regId

	models.UpdateUser(person)
	this.Data["json"] = person
	this.ServeJson()

}
