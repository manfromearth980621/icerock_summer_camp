package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"bingyan/models"
	"fmt"
	//"github.com/astaxie/beego/session"  
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	sess := this.StartSession()
	uname := sess.Get("uname")
	fmt.Println(uname)
	// 判断是否为退出操作
	if this.Input().Get("exit") == "true" {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	// 获取表单信息
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	// 验证用户名及密码
	if uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass") {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}

	sess := this.StartSession()
	/*var account models.Account
	account.accountname = this.Input().Get("uname")
	account.password = this.Input().Get("pwd")*/
	
	err := models.ReadAccount(uname,pwd)
	if err == nil {
		sess.Set("username", uname)
		fmt.Println("username:", sess.Get("username"))
		this.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		this.TplName = "error.tpl"
	}

	this.Redirect("/", 302)
	return
}

//管理员登录的check
func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value
	return uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass")
}
