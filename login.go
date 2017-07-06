package login
	import (
	"github.com/astaxie/beego"
	//"fmt"
	)
type Logincontrol struct{
	beego.Controller
}

func(this *Logincontrol) Get(){
	this.TplNames = "login.html"
}

func (this *Logincontrol) Post(){
	uname := this.Input().Get("unamen")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin" )== "om"
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
			maxAge := 0
			if autoLogin{
				maxAge =1<<31 -1
			}
	
		this.Ctx.SetCookie("uname",uname,maxAge,"/")
		this.Ctx.SetCookie("pwd", pwd, maxAge ,"/")
		}

	this.Redirect("/",301)
	return

} 