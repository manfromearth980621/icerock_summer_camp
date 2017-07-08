package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	this.TplName = ""
	
	/*c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	c.Data["html"]="<div>hello</div>"

	c.Data=("TrueCond") = true
	c.Data=("FalseCond") = false*/
}
