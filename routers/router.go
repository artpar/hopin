package routers

import (
	"github.com/artpar/hopin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/person", &controllers.PersonController{})
	beego.Router("/person/feed", &controllers.PersonController{}, "get:Feed")
	beego.Router("/person/travel/options", &controllers.PersonController{}, "get:TravelOptions")
	beego.Router("/person/travel/join", &controllers.PersonController{}, "post:RequestJoin")
	beego.Router("/traveller", &controllers.TravellerController{})
	beego.Router("/placesapi", &controllers.PlacesController{})
	beego.Router("/placesapi/autocomplete/json", &controllers.PlacesController{})
}
