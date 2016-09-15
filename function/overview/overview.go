package overview

import (
	"github.com/astaxie/beego"
)

type OverviewController struct {
	beego.Controller
}

func (this *OverviewController) Get() {
	this.Ctx.WriteString("Time:")
}
