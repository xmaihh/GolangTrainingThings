package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beeticket/controllers:AdminController"] = append(beego.GlobalControllerRouter["beeticket/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:AdminController"] = append(beego.GlobalControllerRouter["beeticket/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:AdminController"] = append(beego.GlobalControllerRouter["beeticket/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:adminId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:AdminController"] = append(beego.GlobalControllerRouter["beeticket/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:adminId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:AdminController"] = append(beego.GlobalControllerRouter["beeticket/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:adminId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:AdminController"] = append(beego.GlobalControllerRouter["beeticket/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:AdminController"] = append(beego.GlobalControllerRouter["beeticket/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:TicketController"] = append(beego.GlobalControllerRouter["beeticket/controllers:TicketController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:TicketController"] = append(beego.GlobalControllerRouter["beeticket/controllers:TicketController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:TicketController"] = append(beego.GlobalControllerRouter["beeticket/controllers:TicketController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:ticketId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:TicketController"] = append(beego.GlobalControllerRouter["beeticket/controllers:TicketController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:ticketId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:TicketController"] = append(beego.GlobalControllerRouter["beeticket/controllers:TicketController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:ticketId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:UserController"] = append(beego.GlobalControllerRouter["beeticket/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:UserController"] = append(beego.GlobalControllerRouter["beeticket/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:UserController"] = append(beego.GlobalControllerRouter["beeticket/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:UserController"] = append(beego.GlobalControllerRouter["beeticket/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:UserController"] = append(beego.GlobalControllerRouter["beeticket/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:UserController"] = append(beego.GlobalControllerRouter["beeticket/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeticket/controllers:UserController"] = append(beego.GlobalControllerRouter["beeticket/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
