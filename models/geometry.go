package models

import (
	"github.com/artpar/hopin/helper"
	"github.com/mitchellh/mapstructure"
	"github.com/astaxie/beego"
)

type Response struct {
	Result Result
}

type Result struct {
	Geometry Geometry
}

type Geometry struct {
	Location Location
}

type Location struct {
	Lat float64
	Lng float64
}

func GooglePlaceDetailApi(place_id string) Location {
	url := "https://maps.googleapis.com/maps/api/place/details/json?placeid=" + place_id + "&key=" + helper.GOOGLE_PLACES_API_KEY
	beego.Info("detail api url", url)
	resp := helper.JsonUrlToMap(url)
	var g Response
	mapstructure.Decode(resp, &g)
	beego.Info("geo loc", g)
	return g.Result.Geometry.Location
}
