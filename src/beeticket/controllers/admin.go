package controllers

import (
	"beeticket/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Admin
type AdminController struct {
	beego.Controller
}

// @Title Create
// @Description create admin
// @Param		body		body 	models.Admin  true		"The Admin content"
// @Success 200 {string} models.Admin.Id
// @Failure 403 body is empty
// @router  / [post]
func (a *AdminController) Post() {
	var admin models.Admin
	json.Unmarshal(a.Ctx.Input.RequestBody, &admin)
	adminId := models.AddAdmin(admin)
	a.Data["json"] = map[string]string{"adminId": adminId}
	a.ServeJSON()
}

// @Title Get
// @Description find admin by adminId
// @Param 	adminId		path 	string 	true		"the adminId you want to get"
// @Success 200 {object} models.Admin
// @Failure 403 :adminId is empty
// @router /:adminId [get]
func (a *AdminController) Get() {
	adminId := a.Ctx.Input.Param(":adminId")
	if adminId != "" {
		admin, err := models.GetAdmin(adminId)
		if err != nil {
			a.Data["json"] = err.Error()
		} else {
			a.Data["json"] = admin
		}
	}
	a.ServeJSON()
}

// @Title GetAll
// @Description get all admins
// @Success 200 {object} models.Admin
// @Failure 403 :adminId is empty
// @router / [get]
func (a *AdminController) GetAll() {
	admins := models.GetAllAdmins()
	a.Data["json"] = admins
	a.ServeJSON()
}

// @Title Update
// @Description update the admin
// @Param adminId 	path 	string 	true	"The adminId you want to update"
// @Param body		body	models.Admin	true	"The body about Admin"
// @Success 200 {object} models.Admin
// @Failure 403 :adminId is empty
// @router /:adminId [put]
func (a *AdminController) Put() {
	adminId := a.Ctx.Input.Param(":adminId")
	if adminId != "" {
		var admin models.Admin
		json.Unmarshal(a.Ctx.Input.RequestBody, &admin)

		aa, err := models.UpdateAdmin(adminId, &admin)
		if err != nil {
			a.Data["json"] = err.Error()
		} else {
			a.Data["json"] = aa
		}
	}
	a.ServeJSON()
}

// @Title Delete
// @Description delete the admin
// @Param	adminId		path 		string 		true 	"The adminId you wnat to delete"
// @Success 200 {string} delete success!
// @Failure 403 adminId is empty
// @router /:adminId [delete]
func (a *AdminController) Delete() {
	adminId := a.Ctx.Input.Param(":adminId")
	models.DeleteAdmin(adminId)
	a.Data["json"] = "delete success!"
	a.ServeJSON()
}


// @Title Login
// @Description Logs admin into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 admin not exist
// @router /login [get]
func (admin *AdminController) Login() {
	username := admin.GetString("username")
	password := admin.GetString("password")
	if models.Login(username, password) {
		admin.Data["json"] = "login success"
	} else {
		admin.Data["json"] = "adminer not exist"
	}
	admin.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in admin session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *AdminController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
