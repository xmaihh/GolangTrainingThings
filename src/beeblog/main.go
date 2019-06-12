package main

import (
    "beeblog/models"
    
    "github.com/astaxie/beego"

    "strings"
    "time"
)

func main() {
	beego.AddFuncMap("processTags", processTags)
	beego.AddFuncMap("getCateName", getCateName)
	beego.AddFuncMap("formatTime", formatTime)

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
