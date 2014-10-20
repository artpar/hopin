package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Travel struct {
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
	StartTime   time.Time
	EndTime     time.Time
	People      int64
}

type Arrangement struct {
	Id            int64
	TravelHostId  int64
	TravelGuestId int64
	RequestedOn   time.Time
	Status        int
	Message       string
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Travel))
	orm.RegisterModel(new(Arrangement))
	beego.Info("Added model Traveller, Arrangement")
}

func AddTraveller(r Travel) Travel {
	Orm.Insert(&r)
	return r
}

func GetArrangementByHostGuest(host, guest int64) (Arrangement, error) {
	a := Arrangement{TravelHostId: host, TravelGuestId: guest}
	err := Orm.Read(&a, "travel_host_id", "travel_guest_id")
	if err != nil {
		beego.Info("error while fetching existing arrangement", err)
	}
	beego.Info("arr", a)
	return a, err
}

func AddArrangement(a Arrangement) (int64, error) {
	id, err := Orm.Insert(&a)
	if err != nil {
		beego.Info("failed to insert", err)
	}
	return id, err
}

func GetTravelById(id int64) Travel {
	r := Travel{Id:id}
	Orm.Read(&r)
	return r
}
