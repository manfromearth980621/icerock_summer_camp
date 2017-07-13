package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"bingyan/models"
)

type IndexController struct {
	beego.Controller
}

func (index *IndexController) Get() {
	sess := index.StartSession()
	username := sess.Get("accountname")
	fmt.Println(accountname)
	if username == nil || username == "" {
		index.TplName = "index.tpl"
	} else {
		index.TplName = "success.tpl"
	}

}

func (index *IndexController) Post() {
	sess := index.StartSession()
	var user models.User
	inputs := index.Input()
	//fmt.Println(inputs)
	user.Accountname = inputs.Get("accountname")
	user.Password = inputs.Get("password")
	err := models.ValidateAccount(user)
	if err == nil {
		sess.Set("accountname", user.Accountname)
		fmt.Println("accountname:", sess.Get("accountname"))
		index.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		index.TplName = "error.tpl"
	}
}