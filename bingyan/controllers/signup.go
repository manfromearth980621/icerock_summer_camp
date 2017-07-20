package controllers

import (
	"bingyan/models"
	//"encoding/json"
	//"strconv"
	"github.com/astaxie/beego"
	//"fmt"
)

type SignupController struct {
	beego.Controller
}



func (this *SignupController) Get(){
	this.Data["IsSignup"] = true
	this.TplName = "signup.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	accountname := this.Input().Get("accountname")
			password := this.Input().Get("password")
			phonenumber := this.Input().Get("phonenumber")

			err :=models.AddAccount(accountname,password,phonenumber)
			
			if err != nil {
				beego.Error(err)
			}
}

/*func (this *SignupController) POST() {
			
			accountname := this.Input().Get("accountname")
			password := this.Input().Get("password")
			phonenumber := this.Input().Get("phonenumber")

			err :=models.AddAccount(accountname,password,phonenumber)
			
			if err != nil {
				beego.Error(err)
			}
			this.Redirect("/", 302)
}*/
			/*acc := account{}

			if err :=this.ParseForm(&acc);err !=nil{
				fmt.Println(err)
			}*/


func (this *SignupController) Add(){
	/*if  checkAccount(this.Ctx) {
		this.Redirect("/regist/add", 302)
		return
	}*/

	this.TplName = "signup_add.html"
}
			/*var account models.Account
    		json.Unmarshal(this.Ctx.Input.RequestBody, &account)
    		err := models.AddAccount(account)
    		//this.Data["json"] = "{\"ObjectId\":\"" + strconv.ParseInt(objectid,10,64) + "\"}"
   		  	this.ServeJSON()

			if err != nil {
			beego.Error(err)
			}*/
			
			//var account models.Account
			
			//user := models.Account
			/*tid := this.Input().Get("tid")
			title := this.Input().Get("title")
			content := this.Input().Get("content")*/

			/*Accountname := this.Input().Get("accountname")
			Password := this.Input().Get("passsword")
			Phonenumber := this.Input().Get("phonenumber")

			err := models.AddAccount(Accountname,Password,Phonenumber)
			if err == nil {
				this.TplName = "success.tpl"
			} else {
				fmt.Println(err)
				this.TplName = "error.tpl"
			}*/
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
			
			



/*func (this *RegistController) Delete() {

	err := models.DeleteAccount(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/regist", 302)
}
*/

