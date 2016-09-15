package router

import (
	overview "beego/test1/function/overview"

	"github.com/astaxie/beego"
)

func Addroute() {
	beego.Router("/overview", &overview.OverviewController{})
}
