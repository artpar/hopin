package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Ride struct {
	Id         int64
	UserId     int
	From       string
	To         string
	StartTime  int64
	EndTime    int64
	Capacity   int
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Ride))
	beego.Info("Added model Ride")
}

func AddRide(r Ride) Ride {
	Orm.Insert(&r)
	return r
}

func GetRideById(id int64) Ride {
	r := Ride{Id:id}
	return Orm.Read(&r)
}
