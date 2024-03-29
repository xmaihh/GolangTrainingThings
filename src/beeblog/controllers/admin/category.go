package admin

import (
	"beeblog/controllers"
	"beeblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"

	"strconv"
	"time"
)

type CategoryController struct {
	controllers.BaseController
}

//列表
func (this *CategoryController) Get() {
	name := this.Input().Get("name")
	page := this.Input().Get("p")
	var pageNo int64
	if len(page) == 0 {
		pageNo = 1
	} else {
		var err error
		pageNo, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	count, err := models.GetCategoryCount(name)
	if err != nil {
		beego.Error(err)
	}
	cates, err := models.GetCategoryList(name, pageNo, models.PageSize)
	if err != nil {
		beego.Error(err)
	}
	p := pagination.NewPaginator(this.Ctx.Request, models.PageSize, count)
	this.Data["Pager"] = p
	this.Data["List"] = cates
	this.Data["Name"] = name
	this.Data["Title"] = "分类"
	this.TplName = "admin/category/index.html"
}

func (this *CategoryController) GetAll() {
	cates, err := models.GetCategoryList("", 1, 999999)

	if err != nil {
		beego.Error(err)
		this.ResponseJson(500, err.Error(), true)
	} else {
		this.ResponseJson(200, cates, true)
	}
}

//添加和修改
func (this *CategoryController) Post() {
	strid := this.Input().Get("id")
	if len(strid) > 0 { //编辑
		id, err := strconv.ParseInt(strid, 10, 64)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(400, "invalid params", true)
		}

		cate, err := models.GetCategoryOne(id)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(400, "can not find this category", true)
		}
		cate.Name = this.Input().Get("name")
		cate.UpdatedAt = time.Now()
		err = models.Update(cate)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(500, err.Error(), true)
		}
	} else { //新增
		cate := &models.Category{
			Name: this.Input().Get("name"),
		}
		err := models.Add(cate)
		if err != nil {
			beego.Error(err)
		}
	}
	this.ResponseJson(200, "success", true)
}

func (this *CategoryController) Delete() {
	beego.Info(this.GetInt64("id"))
	strid := this.Input().Get("id")
	if len(strid) > 0 {
		id, err := strconv.ParseInt(strid, 10, 64)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(400, "invalid params", true)
		}
		cate := &models.Category{
			Id: id,
		}
		err = models.Delete(cate)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(400, "can not find this category", true)
		}
		this.ResponseJson(200, "success", true)
	} else {
		this.ResponseJson(400, "need id"+strid, true)
	}
}
