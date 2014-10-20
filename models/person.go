package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Person struct {
	Id    int
	Email string
	RegId string
}

type TravelOptions struct {
	HostId    int
	Name      string
	Email     string
	StartTime time.Time
	From      string
	To        string
	People    int
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Person))
	beego.Info("Added model Person")
}

func UpdateUser(p Person) {
	Orm.Update(&p)
}

func GetUserTravels(user Person) []Travel {
	var travellers []Travel
	Orm.QueryTable("travel").Filter("user_id", user.Id).All(&travellers)
	return travellers
}

func GetTravelOptions(travelId int64) []TravelOptions {
	var options []TravelOptions
	_, err := Orm.Raw(`select p.name, p.email, t1.id as host_id, t1.start_time, t1.from, t1.to
from travel t1
join travel t2
join person p on p.id = t1.user_id
where t1.from = t2.from
and t1.to = t2.to
and t1.id != ?
and t1.people + t2.people + (select sum(people) from travel where id in (select travel_guest_id from arrangement where travel_host_id = t1.id)) >= 0
and t2.id = ?`, travelId, travelId).QueryRows(&options)
	if err != nil {
		beego.Info("Failed to get data ", err)
	}
	return options
}


func GetUserByEmail(email string) Person {
	p := Person{Email: email}
	Orm.Read(&p, "Email")
	return p
}

func Test() {
	beego.Info("start insert")
	p := Person{Email:"test@gmail.com"}
	beego.Info("insert p", p, Orm)
	id, err := Orm.Insert(&p)
	beego.Info("err", err)
	if err != nil {
		panic(err)
	}
	beego.Info("new Id: ", id)
}

func CreateUser(p Person) Person {
	Orm.Insert(&p)
	return p
}

