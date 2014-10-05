package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Traveller struct {
	Id          int64
	UserId      int
	From        string
	FromLat     float64
	FromLon     float64
	FromPlaceId string
	To          string
	ToLat       float64
	ToLon       float64
	ToPlaceId   string
	StartTime   string
	EndTime     string
	People      int64
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Traveller))
	beego.Info("Added model Traveller")
}

func AddTraveller(r Traveller) Traveller {
	Orm.Insert(&r)
	return r
}

func GetTravellerById(id int64) Traveller {
	r := Traveller{Id:id}
	Orm.Read(&r)
	return r
}
