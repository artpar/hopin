package controllers

import (
	"github.com/astaxie/beego"
	"github.com/artpar/hopin/models"
	"time"
)

type PersonController struct {
	beego.Controller
	user  models.Person
	email string
}

type PersonFeed struct {
	Travels []models.Travel
}

func (this *PersonController) TravelOptions() {
	travelId, err := this.GetInt("travelId")
	if err != nil {
		this.Data["json"] = "unable to get travel id"
		this.ServeJson()
		return
	}
	options := models.GetTravelOptions(travelId)
	if options != nil {

		this.Data["json"] = options
	} else {
		this.Data["json"] = []models.TravelOptions{}
	}
	this.ServeJson()
}

func (this *PersonController) ShowError(msg string) {
	info := make(map[string]string)
	info["Message"] = msg
	this.Data["json"] = info
	this.ServeJson()
}

func (this *PersonController) RequestJoin() {
	if this.ValidateRequest() {
		host, err := this.GetInt("hostId")
		if err != nil {
			this.ShowError("host id not present")
			return
		}
		guestId, err := this.GetInt("travelId")
		if err != nil {
			this.ShowError("guest id not present")
			return
		}
		guestTravel := models.GetTravelById(guestId)
		if guestTravel.UserId != this.user.Id {
			this.ShowError("you cannot request this join")
			return
		}

		hostTravel := models.GetTravelById(host)
		if hostTravel.Id != host {
			this.ShowError("This is not a valid host travel")
			return
		}

		arrangement, err := models.GetArrangementByHostGuest(host, guestId)
		if err != nil {
			beego.Info("adding new arrangement", host, guestId, err)
			arrangement.Status = 0
			arrangement.Message = "Join requested"
			arrangement.RequestedOn = time.Now()
			models.AddArrangement(arrangement)
		} else {
			arrangement.Message = "Join already requested"
		}
		this.Data["json"] = arrangement
		this.ServeJson()
	}
}

func (this *PersonController) Get() {
	this.Data["json"] = models.GetUserByEmail(this.GetString("email"))
	this.ServeJson()
}

func (this *PersonController) ValidateRequest() bool {

	this.email = this.GetString("email")
	this.user = models.GetUserByEmail(this.email)
	regid := this.GetString("regid")
	if this.user.RegId != regid {
		this.Data["json"] = "not a proper request"
		this.ServeJson()
		return false
	}

	return true
}

// Todo: implement the method to get a single travel of a user by id
// Todo: Implement the table to store user request for other travels and accept and reject status

func (this *PersonController) Feed() {

	if !this.ValidateRequest() {
		return;
	}

	var feed PersonFeed;

	feed.Travels = models.GetUserTravels(this.user)
	this.Data["json"] = feed
	// beego.Info("feed", feed)
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
