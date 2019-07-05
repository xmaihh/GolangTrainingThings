// @APIVersion 1.0.0
// @Title Get a ticket API
// @Description A program to receive QR code tickets online.
// @Contact https://github.com/xmaihh/GolangTrainingThings
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beeticket/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/ticket",
			beego.NSInclude(
				&controllers.TicketController{},
			)),
		beego.NSNamespace("/admin",
			beego.NSInclude(
				&controllers.AdminController{},
			)),
	)
	beego.AddNamespace(ns)
}
