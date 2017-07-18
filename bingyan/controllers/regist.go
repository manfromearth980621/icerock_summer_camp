package controllers

import (
	"bingyan/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegistController struct {
	beego.Controller
}

func (this *RegistController) Get() {
	//this.Data["IsSignup"] = true
	this.TplName = "regist.html"
	//this.Data["IsLogin"] = checkAccount(this.Ctx)

}

func (this *RegistController) Add(){
	if  checkAccount(this.Ctx) {
		this.Redirect("/home", 302)
		return
	}

	this.TplName = "regist.html"
}

func (this *RegistController) POST() {

			//var account models.Account
			inputs := this.Input()
			//user := models.Account
			/*tid := this.Input().Get("tid")
			title := this.Input().Get("title")
			content := this.Input().Get("content")*/

			Accountname := inputs.Get("account")
			Password := inputs.Get("passsword")
			Phonenumber := inputs.Get("phonenumber")

			err := models.AddAccount(Accountname,Password,Phonenumber)
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
			
			this.Redirect("/home", 302)
}

func (this *RegistController) Delete() {

	err := models.DeleteAccount(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/regist", 302)
}

