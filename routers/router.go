package routers

import (
	"github.com/artpar/hopin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/person", &controllers.PersonController{})
	beego.Router("/ride", &controllers.RideController{})
	beego.Router("/traveller", &controllers.TravellerController{})
	beego.Router("/placesapi", &controllers.PlacesController{})
	beego.Router("/placesapi/autocomplete/json", &controllers.PlacesController{})
}
