package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Ride struct {
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
	orm.RegisterModel(new(Ride))
	beego.Info("Added model Ride")
}

func AddRide(r Ride) Ride {
	Orm.Insert(&r)
	return r
}

func GetRideById(id int64) Ride {
	r := Ride{Id:id}
	Orm.Read(&r)
	return r
}
