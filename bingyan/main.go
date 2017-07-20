package main

import (
	"bingyan/controllers"
	"bingyan/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {

	beego.BConfig.WebConfig.Session.SessionOn = true//session ok

	
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	// 注册 beego 路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/category", &controllers.CategoryController{})
	
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})

	beego.Router("/login", &controllers.LoginController{})
	
	beego.Router("/signup", &controllers.SignupController{})
	beego.AutoRouter(&controllers.SignupController{})

	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{},"post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{},"get:Delete")
	//beego.Router("/index", &controllers.IndexController{})
	
	beego.Run()			// 启动 beego
	
}
