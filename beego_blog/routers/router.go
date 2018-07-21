package routers

import (
	"github.com/Echosong/beego_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.BlogController{}, "*:Home")
	beego.Router("/bloglist", &controllers.BlogController{}, "*:Bloglist")
	beego.Router("/blogtypelist/:hi:string", &controllers.BlogController{}, "*:Blogtypelist")
	beego.Router("/datetypelist", &controllers.BlogController{}, "*:Datetypelist")
	beego.Router("/blogdetail/?:id", &controllers.BlogController{}, "*:Blogdetail")


	beego.AutoRouter(&controllers.AdminController{})
}
