package controllers

import (
	"bingyan/models"

	"github.com/astaxie/beego"
)

type SignupController struct {
	beego.Controller
}

func (this *SignupController) Get() {
	//this.Data["IsSignup"] = true
	this.TplName = "signup.html"
	//this.Data["IsLogin"] = checkAccount(this.Ctx)

}

/*func (this *SignupController) Add(){
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplName = "signup_add.html"
}*/

func (this *SignupController) POST() {

			var account models.Account
			inputs := this.Input()

			user.Accountname = inputs.Get("account")
			user.Password = inputs.Get("passsword")
			user.Phonenumber = inputs.Get("phonenumber")

			err := models.SaveUser(user)
			if err == nil {
				this.TplName = "success.tpl"
			} else {
				fmt.Println(err)
				this.TplName = "error.tpl"
			}
			/*k := kenan{}
			if err := this.ParseForm(&k); err != nil {
        				beego.Error(err)				//handle error
    		}
			if !checkAccount(this.Ctx) {
						this.Redirect("/login", 302)
						return
			}*/
	

			
			/*account := this.Input().Get("account")
			password := this.Input().Get("password")
			phonenumber := this.Input().Get("phonenumber")*/

			/*if ((len(account)==0)&&(len(password)==0)&&(len(phonenumber)==0)){
				break
			}*/

			/*err := models.AddAccount(account,password,phonenumber)
			if err != nil {
				beego.Error(err)
			}*/
			
			this.Redirect("/signup", 302)
}


