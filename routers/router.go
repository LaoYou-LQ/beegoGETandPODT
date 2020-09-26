package routers

import (
	"beegoResp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.RegisterConeroller{})
    beego.Router("/", &controllers.MainController{})
}
