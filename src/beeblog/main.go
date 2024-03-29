package main

import (
	"beeblog/models"
	_ "beeblog/routers"
	"beeblog/util"
	"net/http"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"

	"strings"
	"time"
)

func main() {
	beego.AddFuncMap("processTags", processTags)
	beego.AddFuncMap("getCateName", getCateName)
	beego.AddFuncMap("formatTime", formatTime)
	//后台鉴权
	beego.InsertFilter("/admin/*", beego.BeforeRouter, func(ctx *context.Context) {
		cookie, err := ctx.Request.Cookie("Authorization")
		beego.Error(cookie)
		if err != nil || !util.CheckToken(cookie.Value) {

			http.Redirect(ctx.ResponseWriter, ctx.Request, "/login", 302)

		}
	})
	beego.Run()
}

func processTags(in string) (out string) {
	out = strings.Trim(strings.Replace(in, "$#", ",", -1), "#$")
	return
}

func getCateName(in int64) (out string) {
	cate, err := models.GetCategoryOne(in)
	if err != nil {
		out = ""
	} else {
		out = cate.Name
	}
	return
}

func formatTime(in time.Time) (out string) {
	out = in.Format("2019-01-01 00:00")
	return
}
