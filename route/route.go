package route

import (
	"github.com/astaxie/beego"
)

func addroute() {
	beego.Router("/", &MainController{})
}
