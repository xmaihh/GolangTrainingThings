package routers

import (
	"beeblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/category/:id", &controllers.MainController{}, "*:GetCateArticles")
	beego.Router("/tag/:name", &controllers.MainController{}, "*:GetTagArticles")
	beego.Router("/article/:id", &controllers.MainController{}, "*:GetArticleDetail")

}
